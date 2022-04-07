package blockingprocessor

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

func processTraces(_ context.Context, traces pdata.Traces) (pdata.Traces, error) {
	time.Sleep(time.Minute) // will block for one minute
	return traces, nil
}
