package queue

import "testing"

func TestEnqueueDequeue(t *testing.T) {
	b := NewMemoryBroker()
	b.Enqueue("test", map[string]string{"key": "val"})
	task, ok := b.Dequeue()
	if !ok || task.Type != "test" {
		t.Fatalf("unexpected task: %+v ok=%v", task, ok)
	}
}

func TestDeadLetterAfterMaxRetry(t *testing.T) {
	b := NewMemoryBroker()
	task := b.Enqueue("fail", nil)
	task.MaxRetry = 2
	b.Retry(&task)
	b.Retry(&task)
	if b.DeadLetterCount() != 1 {
		t.Fatalf("expected 1 dead letter, got %d", b.DeadLetterCount())
	}
}
