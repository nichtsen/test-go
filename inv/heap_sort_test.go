package inv

import "fmt"

// a zero-base heap
type heap struct{}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) left(i int) int {
	return (i+1)*2 - 1
}

func (h *heap) right(i int) int {
	return (i + 1) * 2
}

var xheap = heap{}

func HeapSort(a []int) {
	buildHeap(a)
	for i := len(a) - 1; i > 0; i-- {
		tmp := a[i]
		a[i] = a[0]
		a[0] = tmp
		if i > 1 {
			maxHeapify(a, 0, i)
		}
	}
}

// maintain a heap
func maxHeapify(a []int, i, N int) {
	largest := i
	l := xheap.left(i)
	r := xheap.right(i)
	if l < N && a[l] > a[i] {
		largest = l
	}
	if r < N && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		tmp := a[largest]
		a[largest] = a[i]
		a[i] = tmp
		maxHeapify(a, largest, N)
	}
}

func buildHeap(a []int) {
	N := len(a) - 1
	init := xheap.parent(N)
	for i := init; i >= 0; i-- {
		maxHeapify(a, i, N)
	}
}

func ExampleHeapSort() {
	a := []int{0, 6, 2, 5, 4}
	HeapSort(a)
	fmt.Println(a)
	// Output:
	// [0 2 4 5 6]
}
