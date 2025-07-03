package main

import (
	"fmt"
	"math/rand"
	"runtime"
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

	histogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "demo-histogram-metric",
			Help:    "Histogram metric to test memory allocation",
			Buckets: prometheus.LinearBuckets(0, 10, 10),
		},
		[]string{"labelA", "labelB"},
	)

	counterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "demo-counter-metric",
			Help: "Counter metric to test memory allocation",
		},
		[]string{"labelA", "labelB"},
	)

	gaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "demo-gauge-metric",
			Help: "Gauge metric to test memory allocation",
		},
		[]string{"labelA", "labelB"},
	)
)

var memGauge = prometheus.NewGaugeFunc(
	prometheus.GaugeOpts{Name: "go_alloc_bytes", Help: "Heap alloc"},
	func() float64 {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		return float64(m.Alloc)
	},
)

func init() {
	prometheus.MustRegister(histogramVec, counterVec, gaugeVec, memGauge, summaryVec)
}

func startClient() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Stop after N entries
	max := 10000
	count := 0
	for range ticker.C {
		if count >= max {
			break
		}
		count++
		// Simulate high cardinality
		labelA := fmt.Sprintf("valA_%d", time.Now().UnixNano())
		labelB := fmt.Sprintf("valB_%d", rand.Intn(1000000))

		// Record values for each metric type
		summaryVec.WithLabelValues(labelA, labelB).Observe(float64(rand.Intn(100)))
		histogramVec.WithLabelValues(labelA, labelB).Observe(float64(rand.Intn(100)))
		counterVec.WithLabelValues(labelA, labelB).Inc()
		gaugeVec.WithLabelValues(labelA, labelB).Set(float64(rand.Intn(100)))
	}
}
