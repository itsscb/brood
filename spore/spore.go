package spore

import (
	"time"

	"github.com/google/uuid"
)

type Spore struct {
	ID        string    `json:"id"`
	Hive      string    `json:"hive"`
	Timestamp time.Time `json:"timestamp"`
	Data      []byte    `json:"data"`
}

func New(hive string, data []byte) Spore {
	id := uuid.New()
	return Spore{
		ID:        id.String(),
		Hive:      hive,
		Timestamp: time.Now(),
		Data:      data,
	}
}
