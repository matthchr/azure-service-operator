/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/jellydator/ttlcache/v3"
	"github.com/pkg/errors"
)

type ProgramCacher interface {
	Start()
	Stop()
	Get(resource reflect.Type, expression string) (*ProgramCacheItem, error)
}

type ProgramCacheItem struct {
	AST     *cel.Ast
	Program cel.Program
}

type ProgramCache struct {
	// key is <env-key>-<expression>. The key must contain a string uniquely identifying the env
	// as the same expression may have different meanings depending on the env.
	// Note that this is expressly safe to cache as
	// per https://github.com/google/cel-go?tab=readme-ov-file#parse-and-check.
	cache    *ttlcache.Cache[string, *ProgramCacheItem]
	envCache *EnvCache
	compile  func(env *cel.Env, expression string) (*ProgramCacheItem, error)
}

var _ ProgramCacher = &ProgramCache{}

func NewProgramCache(
	envCache *EnvCache,
	compile func(env *cel.Env, expression string) (*ProgramCacheItem, error),
) *ProgramCache {
	// TODO: Logging on cache hits/misses?
	// TODO: Logging/metrics on time spent compiling?
	// TODO: Comments
	// 	cache.OnInsertion(func(ctx context.Context, item *ttlcache.Item[string, string]) {
	//		fmt.Println(item.Value(), item.ExpiresAt())
	//	})

	return &ProgramCache{
		cache: ttlcache.New[string, *ProgramCacheItem](
			ttlcache.WithTTL[string, *ProgramCacheItem](24 * time.Hour),
		),
		envCache: envCache,
		compile:  compile,
	}
}

func (c *ProgramCache) Start() {
	go c.envCache.Start()
	go c.cache.Start()
}

func (c *ProgramCache) Stop() {
	c.cache.Stop()
}

func (c *ProgramCache) Get(resource reflect.Type, expression string) (*ProgramCacheItem, error) {
	envKey := getTypeImportPath(resource)

	env, err := c.envCache.Get(resource)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get CEL env")
	}
	key := fmt.Sprintf("%s-%s", envKey, expression)

	// TODO: Do we also want to cache errors?
	item := c.cache.Get(key)
	if item != nil {
		// We found what we wanted, return it
		return item.Value(), nil
	}

	result, err := c.compile(env, expression)
	if err != nil {
		return nil, err
	}

	c.cache.Set(key, result, ttlcache.DefaultTTL)
	return result, nil
}
