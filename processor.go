package blockingprocessor

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

func processTraces(_ context.Context, traces pdata.Traces) (pdata.Traces, error) {
	fmt.Println("waiting one minute")
	time.Sleep(time.Minute) // will block for one minute
	return traces, nil
}
