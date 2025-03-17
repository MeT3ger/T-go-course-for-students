package tagcloud

import (
	"container/heap"
)

type MaxHeap []TagStat

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].OccurrenceCount < h[j].OccurrenceCount }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(TagStat))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (c TagCloud) findKMax(k int) []TagStat {
	if k <= 0 {
		return make([]TagStat, 0)
	}
	if k > len(c.storage) {
		k = len(c.storage)
	}

	h := &MaxHeap{}
	heap.Init(h)

	i := 0
	for key, v := range c.storage {
		if i >= k {
			a := (*h)[0]
			if v > a.OccurrenceCount {
				heap.Pop(h)
				heap.Push(h, TagStat{
					Tag:             key,
					OccurrenceCount: v,
				})
			}
		} else {
			heap.Push(h, TagStat{
				Tag:             key,
				OccurrenceCount: v,
			})
		}
		i++
	}

	result := make([]TagStat, k)
	for i := 0; i < k; i++ {
		maxElem := heap.Pop(h).(TagStat)
		result[k-i-1] = TagStat{
			Tag:             maxElem.Tag,
			OccurrenceCount: maxElem.OccurrenceCount,
		}
	}
	return result
}
