package main

import (
	"fmt"
)

/*
Given an array and a sum k, we need to print the length of the longest subarray that sums to k.

	Example 1:
	Input Format: N = 3, k = 5, array[] = {2,3,5}
	Result: 2
	Explanation: The longest subarray with sum 5 is {2, 3}. And its length is 2.

	Example 2:
	Input Format: N = 3, k = 1, array[] = {-1, 1, 1}
	Result: 3
	Explanation: The longest subarray with sum 1 is {-1, 1, 1}. And its length is 3.
*/

func bruteForce(arr []int, kVal int) int {

	var longest int = 0
	for i := 0; i < len(arr); i++ {

		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			if sum == kVal {
				longest = max(longest, j-i+1)
			}
		}
	}

	return longest
	// TC = O(N3) approx., where N = size of the array.
	// Reason: We are using three nested loops, each running approximately N times.

	// SC = O(1) as we are not using any extra space.
}

func better(arr []int, kVal int) int {

	var longest int = 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == kVal {
				longest = max(longest, j-i+1)
			}
		}
	}
	return longest
	// TC = O(N2) approx., where N = size of the array.
	// Reason: We are using three nested loops, each running approximately N times.

	// SC = O(1) as we are not using any extra space.
}

func optimal(arr []int, kVal int) int {
	longest, sum := 0, 0
	preSum := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == kVal {
			longest = max(longest, i+1)
		}
		rem := sum - kVal
		if _, ok := preSum[rem]; ok {
			longest = max(longest, i-preSum[rem])
		}
		if _, ok := preSum[sum]; !ok {
			preSum[sum] = i
		}
	}

	return longest
	// TC =  O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
	// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N)(though in the worst case,
	// unordered_map takes O(N) to find an element and the time complexity becomes O(N2)) but if we are using a map data structure, the time complexity will be O(N*logN).
	// The least complexity will be O(N) as we are using a loop to traverse the array.

	// SC = O(N) as we are using a map data structure.
}

func main() {
	arr := []int{-1, 1, 1}
	k := 1
	fmt.Println("ResultBF: ", bruteForce(arr, k))
	fmt.Println("ResultB: ", better(arr, k))
	fmt.Println("ResultO: ", optimal(arr, k))
}
