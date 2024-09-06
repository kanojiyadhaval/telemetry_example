package main

import (
	"fmt"
	"log"

	"github.com/kanojiyadhaval/telemetry"
	"github.com/kanojiyadhaval/telemetry/config"
	"github.com/kanojiyadhaval/telemetry/drivers"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var logger telemetry.Logger

	// Choose logger based on config
	switch conf.DriverType {
	case "cli":
		logger = drivers.NewCLI()
	case "json":
		logger = drivers.NewJSONFile(conf.FilePath)
	case "text":
		logger = drivers.NewTextFile(conf.FilePath)
	default:
		log.Fatalf("Unsupported driver type: %s", conf.DriverType)
	}

	telemetryInstance := telemetry.NewTelemetry(logger)

	// Create transaction log
	transactionAttributes := map[string]string{"userId": "123"}
	transaction := telemetry.NewTransaction(telemetryInstance, "txn-456", transactionAttributes)

	// Log the transaction start
	transaction.Start()

	// Log a message within the transaction
	transaction.Log(telemetry.Info, "Processed request", map[string]string{"status": "200 OK"})

	// Log the transaction end
	transaction.End()

	fmt.Println("Logging complete.")
}
