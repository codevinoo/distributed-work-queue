package queue

import (
	"fmt"
	"sync"
	"time"
)

type MemoryBroker struct {
	mu       sync.Mutex
	pending  []Task
	dead     []Task
	seq      int
}

func NewMemoryBroker() *MemoryBroker {
	return &MemoryBroker{}
}

func (b *MemoryBroker) Enqueue(taskType string, payload map[string]string) Task {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.seq++
	t := Task{
		ID:        fmt.Sprintf("task-%d", b.seq),
		Type:      taskType,
		Payload:   payload,
		MaxRetry:  3,
		CreatedAt: time.Now(),
	}
	b.pending = append(b.pending, t)
	return t
}

func (b *MemoryBroker) Dequeue() (*Task, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.pending) == 0 {
		return nil, false
	}
	t := b.pending[0]
	b.pending = b.pending[1:]
	return &t, true
}

func (b *MemoryBroker) Retry(t *Task) {
	b.mu.Lock()
	defer b.mu.Unlock()
	t.Attempts++
	if t.Attempts >= t.MaxRetry {
		b.dead = append(b.dead, *t)
		return
	}
	b.pending = append(b.pending, *t)
}

func (b *MemoryBroker) DeadLetterCount() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.dead)
}
