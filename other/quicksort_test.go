// Problem:
// Implement quicksort.
//
// Mechanic:
// Recursively divide the input into two smaller arrays around a pivot, where
// one half has items smaller than the pivot, other half has items bigger than
// the pivot.
//
// Cost:
// O(nlogn) time and O(nlogn) space.

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		selectionSort(tt.in)
		common.Equal(t, tt.expected, tt.in)
	}
}

func quickSort(in []int, start, end int) {
	if start < end {
		// pi is the pivot/partition index.
		pi := partition(in, start, end)

		// sort the items before and after partition.
		quickSort(in, start, pi-1)
		quickSort(in, pi+1, end)
	}
}

func partition(in []int, start, end int) int {
	pivot := in[end]

	left := start
	right := end - 1

	for left <= right {
		// keep going until we find something on the left that belongs to the
		// right.
		for left <= end && in[left] < pivot {
			left++
		}

		// keep going until we find something on the right that belongs to the
		// left.
		for right >= start && in[right] >= pivot {
			right--
		}

		// by swapping the item at left and right index, we move the item that
		// is smaller than the pivot to the left half and vice versa.
		if left < right {
			tmp := in[left]
			in[left] = in[right]
			in[right] = tmp
		} else {
			// once the partition is finished, move the pivot back to its final
			// position by swapping the item at left and end index.
			tmp := in[left]
			in[left] = in[end]
			in[end] = tmp
		}
	}

	return left
}
