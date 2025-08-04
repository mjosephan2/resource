package heap

/*
	MinHeap is a binary tree where the parent node is always less than or equal to the child nodes.
	1. Insert:
		- Insert the value at the end of the array
		- Heapify up to maintain the heap property
	2. Extract:
		- Extract the minimum value from the heap
		- Replace the root with the last element
		- Heapify down to maintain the heap property
	3. Heapify up:
		- Compare the inserted value with its parent
		- If the inserted value is less than its parent, swap them
		- Repeat the process until the heap property is maintained
	4. Heapify down:
		- Compare the root with its children
		- If the root is greater than any of its children, swap them
		- Repeat the process until the heap property is maintained
	heap property
	- left child := index * 2 + 1
	- right child := index * 2 + 2
	- parent := (index - 1) / 2

	heap complexity
	- insert: O(log n)
	- extract: O(log n)
	- get min: O(1)

	heap application
	- priority queue
	- heap sort
	- dijkstra's algorithm
	- prim's algorithm
	- huffman coding
*/
type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		arr: make([]int, 0),
	}
}

func (h *MinHeap) Insert(val int) {
	/*
		Insert the value at the end of the array
		Heapify up to maintain the heap property
	*/
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) Extract() int {
	/*
		Extract the minimum value from the heap
		Replace the root with the last element
		Heapify down to maintain the heap property
	*/
	if len(h.arr) == 0 {
		return -1
	}
	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	smallest := index
	if leftChildIndex < len(h.arr) && h.arr[leftChildIndex] < h.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightChildIndex < len(h.arr) && h.arr[rightChildIndex] < h.arr[smallest] {
		smallest = rightChildIndex
	}
	if smallest != index {
		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.arr[index] < h.arr[parentIndex] {
			h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
		}
	}
}
