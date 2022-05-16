package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func ConnectCli(apiUrlCLI string) {

	start := time.Now()

	client, err := api.NewClient(api.Config{
		Address: apiUrlCLI,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	query := `sum(rate(ops_response_latency_ms_bucket{le="1000", direction="inbound", namespace="sre",  status!~"5.."}[6h])) by (alias) /  ignoring(le)  sum(rate(ops_response_latency_ms_count{direction="inbound", namespace="sre"}[6h])) by (alias) * 100`

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, query, r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)

	newresult, _ := v1api.Targets(ctx)
	prettyJson, _ := json.MarshalIndent(newresult, "", "    ")
	fmt.Println(string(prettyJson))

	fmt.Println("\n", "Query Duration:", time.Since(start))

}
