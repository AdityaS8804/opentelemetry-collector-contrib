// Copyright The OpenTelemetry Authors
// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package processors

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/cmd/testutils"
)

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
