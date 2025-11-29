package main

import (
	"fmt"
)

/*
Given an array of numbers, you need to return the count of reverse pairs. Reverse Pairs are those pairs where i<j and arr[i]>2*arr[j].

	Example 1:
	Input: N = 5, array[] = {1,3,2,3,1)
	Output: 2
	Explanation: The pairs are (3, 1) and (3, 1) as from both the pairs the condition arr[i] > 2*arr[j] is satisfied.

	Example 2:
	Input: N = 4, array[] = {3,2,1,4}
	Output: 1
	Explanation: There is only 1 pair  ( 3 , 1 ) that satisfy the condition arr[i] > 2*arr[j]

*/

func bruteForce(arr []int) int {
	n, cnt := len(arr), 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > 2*arr[j] {
				cnt++
			}
		}
	}

	return cnt
	// TC = O(N2), We are using nested loops here and those two loops roughly run for N times.
	// SC = O(1), since no extra space is used apart from variables.
}

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

func mergeSort(arr []int, low, high int) int {
	cnt := 0
	if low >= high {
		return cnt
	}
	mid := (low + high) / 2
	cnt += mergeSort(arr, low, mid)
	cnt += mergeSort(arr, mid+1, high)
	cnt += countPairs(arr, low, mid, high)
	merge(arr, low, mid, high)
	return cnt
}

func optimal(arr []int) int {
	n := len(arr)
	return mergeSort(arr, 0, n-1)
	// TC = O(2N*logN), Inside the mergeSort() we call merge() and countPairs() except mergeSort() itself.
	// Now, inside the function countPairs(), though we are running a nested loop, we are actually iterating the left half once and the right half once in total.
	// That is why, the time complexity is O(N). And the merge() function also takes O(N). The mergeSort() takes O(logN) time complexity.
	// Therefore, the overall time complexity will be O(logN * (N+N)) = O(2N*logN).

	// SC = : O(N), as in the merge sort We use a temporary array to store elements in sorted order.
}

func main() {
	arr := []int{1, 3, 2, 3, 1}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
