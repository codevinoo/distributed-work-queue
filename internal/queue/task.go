package queue

import (
	"encoding/json"
	"time"
)

type Task struct {
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	Payload   map[string]string `json:"payload"`
	Attempts  int               `json:"attempts"`
	MaxRetry  int               `json:"max_retry"`
	CreatedAt time.Time         `json:"created_at"`
}

func (t *Task) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func UnmarshalTask(data []byte) (*Task, error) {
	var t Task
	err := json.Unmarshal(data, &t)
	return &t, err
}
