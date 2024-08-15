/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type CEL struct {
	envCacheHits       *prometheus.CounterVec
	envCacheMisses     *prometheus.CounterVec
	programCacheHits   *prometheus.CounterVec
	programCacheMisses *prometheus.CounterVec
	compilationTime    *prometheus.HistogramVec
}

var _ Metrics = &CEL{}

func NewCEL() *CEL {
	envCacheHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cel_env_cache_hits_total",
		Help: "Total number of CEL env cache hits",
	}, []string{"resource"})
	envCacheMisses := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cel_env_cache_misses_total",
		Help: "Total number of CEL env cache misses",
	}, []string{"resource"})
	programCacheHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cel_program_cache_hits_total",
		Help: "Total number of CEL env cache hits",
	}, []string{"resource"})
	programCacheMisses := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cel_program_cache_misses_total",
		Help: "Total number of CEL env cache misses",
	}, []string{"resource"})
	compilationTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "cel_compilation_time_seconds",
		Help: "Time spent on compiling/parsing CEL expressions",
	}, []string{"resource"})

	return &CEL{
		envCacheHits:       envCacheHits,
		envCacheMisses:     envCacheMisses,
		programCacheHits:   programCacheHits,
		programCacheMisses: programCacheMisses,
		compilationTime:    compilationTime,
	}
}

// RegisterMetrics registers the collectors with prometheus server.
func (c *CEL) RegisterMetrics() {
	metrics.Registry.MustRegister(c.envCacheHits, c.envCacheMisses, c.programCacheHits, c.programCacheMisses, c.compilationTime)
}

func (c *CEL) RecordEnvCacheHit(resourceName string) {
	c.envCacheHits.WithLabelValues(resourceName).Inc()
}

func (c *CEL) RecordEnvCacheMiss(resourceName string) {
	c.envCacheMisses.WithLabelValues(resourceName).Inc()
}

func (c *CEL) RecordProgramCacheHits(resourceName string) {
	c.programCacheHits.WithLabelValues(resourceName).Inc()
}

func (c *CEL) RecordProgramCacheMisses(resourceName string) {
	c.programCacheMisses.WithLabelValues(resourceName).Inc()
}

func (c *CEL) RecordCompilationTime(resourceName string, requestTime time.Duration) {
	c.compilationTime.WithLabelValues(resourceName).Observe(requestTime.Seconds())
}
