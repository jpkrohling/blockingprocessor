This requires ocb: https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder

```
ocb --config resources/manifest.yaml
./_build/otelcol-blocking --config resources/otelcol.yaml
```

This will block in its default configuration. Then, add the batching processor to the `traces/1` pipeline and compare the difference in behavior.
