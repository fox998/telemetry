package sink

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fox998/telemetry/internal/common"
)

type Buffer struct {
	data        []common.SensorData
	bufferSize  int
	flushDelta  time.Duration
	logFileName string
}

func NewBuffer(size int, flushDelta time.Duration, logFileName string) *Buffer {
	return &Buffer{
		data:        make([]common.SensorData, 0, size),
		bufferSize:  size,
		flushDelta:  flushDelta,
		logFileName: logFileName,
	}
}

func (b *Buffer) ListenTo(sensorData <-chan common.SensorData) {
	ticker := time.NewTicker(b.flushDelta)
	defer ticker.Stop()

	for {
		select {
		case d, open := <-sensorData:
			if !open {
				b.WriteToFile()
				return
			}

			b.data = append(b.data, d)
			if len(b.data) >= b.bufferSize {
				b.WriteToFile()
			}

		case <-ticker.C:
			b.WriteToFile()
		}
	}
}

func (b *Buffer) WriteToFile() {
	file, err := os.OpenFile(b.logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}
	defer file.Close()

	for _, entry := range b.data {
		logEntry := fmt.Sprintf("%s - Sensor: %s, Id: %v Value: %d\n", entry.Timestamp.Format(time.RFC3339), entry.Name, entry.ID, entry.Value)
		if _, err := file.WriteString(logEntry); err != nil {
			log.Println("Failed to write to log file:", err)
		}
	}

	b.data = b.data[:0]
}
