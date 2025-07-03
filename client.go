package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	summaryVec = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "demo_summary_metric",
			Help:       "Summary metric to test memory allocation",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"labelA", "labelB"},
	)
)

func init() {
	prometheus.MustRegister(summaryVec)
}

func startClient() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		// Generate random label values to simulate label cardinality growth
		labelA := fmt.Sprintf("valA_%d", rand.Intn(1000))
		labelB := fmt.Sprintf("valB_%d", rand.Intn(1000))

		// Record a dummy observation
		summaryVec.WithLabelValues(labelA, labelB).Observe(float64(rand.Intn(100)))
	}
}
