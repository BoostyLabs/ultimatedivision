// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Counter exposes functionality to count metric.
type Counter interface {
	// Inc used to increment Counter.
	Inc()
}

// Metric is an implementation of metrics using prometheus.
type Metric struct {
	Handler  http.Handler
	NewUsers Counter
	Logins   Counter
	Logouts  Counter
	Purchase Counter
}

// NewMetric is a constructor for a Metric.
func NewMetric() *Metric {
	newUsers := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_registrations",
		Help: "The total number of successful registrations.",
	})

	logins := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_logins",
		Help: "The total number of successful logins.",
	})

	logouts := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_logouts",
		Help: "The total number of successful logouts.",
	})

	purchase := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_purchase",
		Help: "The total number of successful purchase.",
	})

	// Create a custom registry.
	registry := prometheus.NewRegistry()
	// Register using our custom registry gauge.
	registry.MustRegister(newUsers)
	registry.MustRegister(logins)
	registry.MustRegister(logouts)
	registry.MustRegister(purchase)
	// Register system metrics.
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	return &Metric{
		// Expose metrics.
		Handler:  promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}),
		NewUsers: newUsers,
		Logins:   logins,
		Logouts:  logouts,
		Purchase: purchase,
	}
}
