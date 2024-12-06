package ecl_heap

import "container/heap"

// LogElement represents a single log entry with a score and the full record value.
type LogElement struct {
	Score  int
	Record string
}

type nmaxHeap []*LogElement

func (nmh nmaxHeap) Head() *LogElement {
	return nmh[0]
}
func (nmh nmaxHeap) Len() int { return len(nmh) }

// Return the lower value because we are a min-heap
func (nmh nmaxHeap) Less(lhs, rhs int) bool {
	return nmh[lhs].Score < nmh[rhs].Score
}

func (nmh nmaxHeap) Swap(lhs, rhs int) {
	nmh[lhs], nmh[rhs] = nmh[rhs], nmh[lhs]
}

func (nmh *nmaxHeap) Push(item any) {
	element := item.(*LogElement)
	*nmh = append(*nmh, element)
}

func (nmh *nmaxHeap) Pop() any {
	old := *nmh
	num := len(old) - 1
	element := old[num]
	*nmh = old[0:num]
	return element
}

type TopLogHeap struct {
	heap *nmaxHeap
	top  int
}

func NewLogHeap(top int) TopLogHeap {
	maxheap := &nmaxHeap{}
	heap.Init(maxheap)

	return TopLogHeap{
		heap: maxheap,
		top:  top,
	}
}

// TryAdd will first fill the heap and then add an element if it is one of the top N, otherwise it will not be part of the heap.
func (tlr *TopLogHeap) TryAdd(element *LogElement) {
	if tlr.heap.Len() < tlr.top {
		heap.Push(tlr.heap, element)
		return
	}

	if element.Score > tlr.heap.Head().Score {
		heap.Pop(tlr.heap)
		heap.Push(tlr.heap, element)
	}
}

// Write will output an array of the top elements is descending order
func (tlr TopLogHeap) Write() []*LogElement {
	result := make([]*LogElement, tlr.heap.Len())
	for i := tlr.heap.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(tlr.heap).(*LogElement)
	}

	return result
}
