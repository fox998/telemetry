package common

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type SensorData struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestemp"`
}

func GenerateBaseSensorData(name string) SensorData {
	return SensorData{
		ID:   uuid.New(),
		Name: name,
	}
}

func (s *SensorData) GenerateValue() {
	s.Value = rand.Int()
	s.Timestamp = time.Now()
}
