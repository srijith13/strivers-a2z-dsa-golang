package main

import (
	"fmt"
)

/*
Given an array of size n, sort the array using Merge Sort.

	Example 1:
	Input : N=7,arr[]={3,2,8,5,1,4,23}
	Output : {1,2,3,4,5,8,23}
	Explanation : Given array is sorted in non-decreasing order.

	Example 2:
	Input : N=5, arr[]={4,2,1,6,7}
	Output : {1,2,4,6,7}
	Explanation : Given array is sorted in non-decreasing order.
*/

func countPairs(arr []int, low, mid, high int) int {
	right := mid + 1
	cnt := 0
	for i := low; i <= mid; i++ {
		for right <= high && arr[i] > 2*arr[right] {
			right++
		}
		cnt += right - (mid + 1)
	}
	return cnt
}

func merge(arr []int, low, mid, high int) {
	left, right := low, mid+1

	var temp []int
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
	// TC = O(N*logN), merging two arrays take linear time and array is recursively divided into halves (logN times).

	// SC = : O(N), as in the merge sort We use a temporary array to store elements in sorted order.
}

func main() {
	arr := []int{3, 2, 8, 5, 1, 4, 23}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("ResultO: ", arr)
}
