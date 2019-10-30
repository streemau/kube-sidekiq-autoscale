// Copyright (c) 2016, M Bogus.
// This source file is part of the KUBE-AMQP-AUTOSCALE open source project
// Licensed under Apache License v2.0
// See LICENSE file for license information

// +build integration

package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetEnqueuedNonExistentHost(t *testing.T) {
	_, err := getEnqueued("http://nonexistanthost/sidekiq/stats", "*")
	if err == nil {
		t.Fatalf("Error expected")
	}
	if got, want := err.Error(), "Get http://nonexistanthost/sidekiq/stats: dial tcp: lookup nonexistanthost: no such host"; !strings.HasPrefix(got, want) {
		t.Errorf("Expected err='%s', got: '%s'", want, got)
	}
}

func TestEnqueued(t *testing.T) {
	_, err := getEnqueued(sidekiqStatsURI(), "*")
	if err != nil {
		t.Fatal(err)
	}
}

func sidekiqStatsURI() string { return os.Getenv("SIDEKIQ_STATS_URI") }
