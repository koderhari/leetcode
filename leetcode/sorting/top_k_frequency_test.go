package sorting

import (
	"container/heap"
	"github.com/stretchr/testify/require"
	"testing"
)

type MaxHeap []info

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Freq > h[j].Freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(info))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type info struct {
	Freq   int
	Number int
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for number, rf := range freq {
		heap.Push(h, info{
			Freq:   rf,
			Number: number,
		})
	}

	res := make([]int, 0, k)

	for h.Len() > 0 && k > 0 {
		top := heap.Pop(h).(info)
		res = append(res, top.Number)
		k--
	}

	return res
}

func Test_topKFrequent(t *testing.T) {
	require.Equal(t, []int{0}, topKFrequent([]int{3, 0, 1, 0}, 1))
}
