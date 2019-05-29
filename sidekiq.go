// Copyright (c) 2016, M Bogus.
// This source file is part of the KUBE-AMQP-AUTOSCALE open source project
// Licensed under Apache License v2.0
// See LICENSE file for license information

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type SidekiqJson struct {
	Enqueued int `json:"enqueued"`
}

type BaseJson struct {
	Sidekiq SidekiqJson `json:"sidekiq"`
}

var (
	statsSuccesses = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "stats_successes_total",
		Help:      "Number of successful stats retrievals.",
	})
	statsFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "stats_failures_total",
		Help:      "Number of failed stats etrievals.",
	})
	currentEnqueuedCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "current_enqueued_count",
			Help:      "Last enqueued count.",
		},
	)
	metricSaveFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "metric_save_failures_total",
		Help:      "Number of times saving metrics failed.",
	})
)

type saveStat func(int) error

func monitorSidekiqStats(uri string, interval int, f saveStat, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		case <-time.After(time.Duration(interval) * time.Second):
			errored := false
			enqueued, err := getEnqueued(uri)

			if err != nil {
				statsFailures.Inc()
				log.Printf("Failed to get stats: %v", err)
				errored = true
			} else {
				statsSuccesses.Inc()
				currentEnqueuedCount.Set(float64(enqueued))
			}

			// Only save metrics if both counts succeeded.
			if errored == false {
				err := f(enqueued)
				if err != nil {
					metricSaveFailures.Inc()
					log.Printf("Error saving metrics: %v", err)
				}
			}
		}
	}
}

func getEnqueued(uri string) (int, error) {
	response, err := http.Get(uri)

	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	data := &BaseJson{}
	err = json.Unmarshal([]byte(body), data)

	if err != nil {
		return 0, err
	}

	return data.Sidekiq.Enqueued, nil
}
