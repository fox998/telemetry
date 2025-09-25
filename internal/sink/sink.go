package sink

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fox998/telemetry/internal/common"
)

func requestConverter(r *http.Request) (common.SensorData, error) {
	if r.Method != http.MethodPost {
		return common.SensorData{}, fmt.Errorf("invalid request method")
	}

	defer r.Body.Close()

	var data common.SensorData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return common.SensorData{}, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return data, nil
}

func shutdownServer(server *http.Server) {
	fmt.Println()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server gracefully stopped.")
}

func RunSink(stop chan os.Signal, sinkArgs SinkArgs) error {

	buffer := NewBuffer(
		*sinkArgs.BufferSize,
		time.Duration(*sinkArgs.WriteInterval)*time.Millisecond,
		*sinkArgs.LogFile)

	messageChan := make(chan common.SensorData, 100)
	defer close(messageChan)

	go buffer.ListenTo(messageChan)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := requestConverter(r)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		fmt.Printf("telemetry from  %v recived\n", data.Name)

		messageChan <- data
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    *sinkArgs.Addr,
		Handler: http.DefaultServeMux,
	}

	go func() {
		log.Println("Server listening on", *sinkArgs.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", *sinkArgs.Addr, err)
		}
	}()

	<-stop

	shutdownServer(server)
	return nil
}
