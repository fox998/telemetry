package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fox998/telemetry/internal/common"
)

func sendTelemetryData(addr string, data common.SensorData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	log.Println(string(jsonData))
	fullAddr := fmt.Sprintf("http://%v", addr)
	resp, err := http.Post(fullAddr, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error sending telemetry data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func RunNode(stop chan os.Signal, nodeArgs NodeArgs) error {
	sensor := common.GenerateBaseSensorData(*nodeArgs.Name)

	ticker := time.NewTicker(time.Second / time.Duration(*nodeArgs.Rate))
	defer ticker.Stop()

	fmt.Printf("Starting node with sensor %s sending data to %s at %d messages per second\n", *nodeArgs.Name, *nodeArgs.Addr, *nodeArgs.Rate)
	for {
		select {
		case <-stop:
			fmt.Println()
			return nil
		case <-ticker.C:
			sensor.GenerateValue()
			err := sendTelemetryData(*nodeArgs.Addr, sensor)
			if err != nil {
				fmt.Println("Failed to send data: " + err.Error())
				continue
			}

			fmt.Println("Telemetry send")
		}
	}
}
