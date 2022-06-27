// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure implements some data structure. eg. list, linklist, stack, queue, tree, graph.
package datastructure

import (
	"github.com/duke-git/lancet/v2/lancetconstraints"
)

// MaxHeap implements a binary max heap
// type T should implements Compare function in lancetconstraints.Comparator interface.
type MaxHeap[T any] struct {
	data       []T
	comparator lancetconstraints.Comparator
}

// NewMaxHeap returns a MaxHeap instance with the given comparator.
func NewMaxHeap[T any](comparator lancetconstraints.Comparator) *MaxHeap[T] {
	return &MaxHeap[T]{
		data:       make([]T, 0),
		comparator: comparator,
	}
}

// Push value into the heap
func (h *MaxHeap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// heapifyUp heapify the data from bottom to top
func (h *MaxHeap[T]) heapifyUp(i int) {
	for h.comparator.Compare(h.data[parentIndex(i)], h.data[i]) < 0 {
		h.swap(parentIndex(i), i)
		i = parentIndex(i)
	}
}

// Pop return the largest value, and remove it from the heap
// if heap is empty, return zero value and fasle
func (h *MaxHeap[T]) Pop() (T, bool) {
	var val T
	if h.Size() == 0 {
		return val, false
	}

	val = h.data[0]
	l := len(h.data) - 1

	h.data[0] = h.data[l]
	h.data = h.data[:l]
	h.heapifyDown(0)

	return val, true
}

// heapifyDown heapify the data from top to bottom
func (h *MaxHeap[T]) heapifyDown(i int) {
	lastIndex := len(h.data) - 1
	l, r := leftChildIndex(i), rightChildIndex(i)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.comparator.Compare(h.data[l], h.data[r]) > 0 {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.comparator.Compare(h.data[i], h.data[childToCompare]) < 0 {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = leftChildIndex(i), rightChildIndex(i)
		} else {
			break
		}
	}
}

// Peek returns the largest element from the heap without removing it.
// if heap is empty, it returns zero value and false.
func (h *MaxHeap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var val T
		return val, false
	}

	return h.data[0], true
}

// Size return the number of elements in the heap
func (h *MaxHeap[T]) Size() int {
	return len(h.data)
}

// Data return data of the heap
func (h *MaxHeap[T]) Data() []T {
	return h.data
}

// parentIndex get parent index of the given index
func parentIndex(i int) int {
	return (i - 1) / 2
}

// leftChildIndex get left child index of the given index
func leftChildIndex(i int) int {
	return 2*i + 1
}

// rightChildIndex get right child index of the given index
func rightChildIndex(i int) int {
	return 2*i + 2
}

// swap two elements in the heap
func (h *MaxHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}