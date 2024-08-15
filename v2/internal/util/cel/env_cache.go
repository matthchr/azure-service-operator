/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"reflect"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/jellydator/ttlcache/v3"
)

type envCacheItem struct {
	env *cel.Env
}

type EnvCache struct {
	cache  *ttlcache.Cache[string, envCacheItem]
	newEnv func(resource reflect.Type) (*cel.Env, error)
}

func NewEnvCache(newEnv func(resource reflect.Type) (*cel.Env, error)) *EnvCache {
	return &EnvCache{
		newEnv: newEnv,
		cache: ttlcache.New[string, envCacheItem](
			ttlcache.WithTTL[string, envCacheItem](24 * time.Hour), // TODO: Configurable?
		),
	}
}

func (c *EnvCache) Start() {
	go c.cache.Start()
}

func (c *EnvCache) Stop() {
	c.cache.Stop()
}

func (c *EnvCache) Get(resource reflect.Type) (*cel.Env, error) {
	key := getTypeImportPath(resource)

	item := c.cache.Get(key)
	if item != nil {
		// We found what we wanted, return it
		return item.Value().env, nil
	}

	env, err := c.newEnv(resource)
	if err != nil {
		return nil, err
	}

	c.cache.Set(key, envCacheItem{env: env}, ttlcache.DefaultTTL)
	return env, nil
}

func coerceList(types []reflect.Type) []any {
	anyTypes := make([]any, 0, len(types))
	for _, t := range types {
		anyTypes = append(anyTypes, t)
	}

	return anyTypes
}
