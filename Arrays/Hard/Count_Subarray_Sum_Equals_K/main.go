package main

import (
	"fmt"
)

/*
Given an array of integers and an integer k, return the total number of subarrays whose sum equals k.
A subarray is a contiguous non-empty sequence of elements within an array.

	Example 1:
	Input Format: N = 4, array[] = {3, 1, 2, 4}, k = 6
	Result: 2
	Explanation: The subarrays that sum up to 6 are [3, 1, 2] and [2, 4].

	Example 2:
	Input Format: N = 3, array[] = {1,2,3}, k = 3
	Result: 2
	Explanation: The subarrays that sum up to 3 are [1, 2], and [3].
*/

func bruteForce(arr []int, sumK int) int {
	cnt, n := 0, len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k < n; k++ {
				sum += arr[k]
			}
			if sum == sumK {
				cnt++
			}
		}
	}

	return cnt
	// TC = O(N3), where N = size of the array.
	// Reason: We are using three nested loops here. Though all are not running for exactly N times, the time complexity will be approximately O(N3).

	// SC = O(1) as we are not using any extra space.
}

func better(arr []int, sumK int) int {
	cnt, n := 0, len(arr)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == sumK {
				cnt++
			}
		}
	}

	return cnt
	// TC =  O(N2), where N = size of the array.
	// Reason: We are using two nested loops here. As each of them is running for exactly N times, the time complexity will be approximately O(N2).
	// SC = O(1) as we are not using any extra space.
}

func optimal(arr []int, sumK int) int {
	cnt, preSum := 0, 0
	mpp := make(map[int]int)
	mpp[0] = 1
	for _, val := range arr {
		preSum += val
		remove := preSum - sumK
		cnt += mpp[remove]
		mpp[preSum] += 1
	}

	return cnt
	// TC = O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
	// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N) but if we are using a map data structure,
	// the time complexity will be O(N*logN). The least complexity will be O(N) as we are using a loop to traverse the array.

	// SC = O(N) as we are using a map data structure.
}

func main() {
	arr := []int{3, 1, 2, 4}
	sumK := 6
	fmt.Println("ResultBF: ", bruteForce(arr, sumK))
	fmt.Println("ResultB: ", better(arr, sumK))
	fmt.Println("ResultO: ", optimal(arr, sumK))
}
