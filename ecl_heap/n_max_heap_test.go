package ecl_heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	maxheap := &nmaxHeap{}
	heap.Init(maxheap)

	heap.Push(maxheap, &LogElement{7, ""})
	heap.Push(maxheap, &LogElement{4, ""})
	heap.Push(maxheap, &LogElement{3, ""})
	heap.Push(maxheap, &LogElement{6, ""})
	heap.Push(maxheap, &LogElement{1, ""})
	heap.Push(maxheap, &LogElement{2, ""})
	heap.Push(maxheap, &LogElement{5, ""})

	expected := 1
	for maxheap.Len() > 0 {
		next := heap.Pop(maxheap).(*LogElement)
		if expected != next.Score {
			t.Errorf("%d is not %d", expected, next.Score)
		}
		expected += 1
	}
}

func TestTopLogHeap(t *testing.T) {
	top_n := 5
	logHeap := NewLogHeap(top_n)

	logHeap.TryAdd(&LogElement{7, ""})
	logHeap.TryAdd(&LogElement{4, ""})
	logHeap.TryAdd(&LogElement{3, ""})
	logHeap.TryAdd(&LogElement{6, ""})
	logHeap.TryAdd(&LogElement{1, ""})
	logHeap.TryAdd(&LogElement{2, ""})
	logHeap.TryAdd(&LogElement{5, ""})

	logs := logHeap.Write()
	if len(logs) != top_n {
		t.Errorf("Length %d is less than %d", len(logs), top_n)
	}
	expected := 7
	for _, log := range logs {
		if expected != log.Score {
			t.Errorf("%d is not %d", expected, log.Score)
		}
		expected -= 1
	}
}
