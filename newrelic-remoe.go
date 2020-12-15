package main

import (
	"context"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
	"github.com/reeve0930/go-remoe"
	"time"

	"os"
)

func main() {
	token := os.Getenv("NATURE_API_TOKEN")
	remoClient := remoe.NewClient(token)
	hervester, _ := telemetry.NewHarvester(telemetry.ConfigAPIKey(os.Getenv("NEW_RELIC_INSERT_API_KEY")))

	data, _ := remoClient.GetRawData()
	for _, d := range data {
		hervester.RecordMetric(telemetry.Gauge{
			Timestamp: time.Now(),
			Value:     float64(d.MeasuredInstantaneous),
			Name:      "remoe.MeasuredInstantaneous",
			Attributes: map[string]interface{}{
				"modelId": d.ModelID,
			},
		})
	}
	hervester.HarvestNow(context.Background())
}
