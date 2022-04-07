package blockingprocessor

import (
	"context"

	"go.opentelemetry.io/collector/model/pdata"
)

func processTraces(_ context.Context, traces pdata.Traces) (pdata.Traces, error) {
	return traces, nil
}
