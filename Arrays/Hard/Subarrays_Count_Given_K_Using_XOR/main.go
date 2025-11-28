package main

import (
	"fmt"
)

/*
Given an array of integers A and an integer B. Find the total number of subarrays having bitwise XOR of all elements equal to k.

	Example 1:
	Input: A = [4, 2, 2, 6, 4] , k = 6
	Output: 4
	Explanation: The subarrays having XOR of their elements as 6 are  [4, 2], [4, 2, 2, 6, 4], [2, 2, 6], [6]

	Example 2:
	Input: A = [5, 6, 7, 8, 9], k = 5
	Output: 2
	Explanation: The subarrays having XOR of their elements as 5 are [5] and [5, 6, 7, 8, 9]
*/

func bruteForce(arr []int, K int) int {
	cnt, n := 0, len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			xorVal := 0
			for k := i; k <= j; k++ {
				xorVal = xorVal ^ arr[k]
			}
			if xorVal == K {
				cnt++
			}
		}
	}

	return cnt
	// TC = O(N3), where N = size of the array.
	// Reason: We are using three nested loops here. Though all are not running for exactly N times, the time complexity will be approximately O(N3).

	// SC = O(1) as we are not using any extra space.
}

func better(arr []int, K int) int {
	cnt, n := 0, len(arr)

	for i := 0; i < n; i++ {
		xorVal := 0
		for j := i; j < n; j++ {
			xorVal ^= arr[j]
			if xorVal == K {
				cnt++
			}
		}
	}

	return cnt
	// TC =  O(N2), where N = size of the array.
	// Reason: We are using two nested loops here. As each of them is running for exactly N times, the time complexity will be approximately O(N2).
	// SC = O(1) as we are not using any extra space.
}

func optimal(arr []int, K int) int {
	cnt, xr := 0, 0
	mpp := make(map[int]int)
	mpp[xr] = 1
	for _, val := range arr {
		xr ^= val
		x := xr ^ K
		cnt += mpp[x]
		mpp[xr]++
	}

	return cnt
	// TC = O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
	// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N) but if we are using a map data structure,
	// the time complexity will be O(N*logN). The least complexity will be O(N) as we are using a loop to traverse the array.

	// SC = O(N) as we are using a map data structure.
}

func main() {
	arr := []int{5, 6, 7, 8, 9}
	K := 5
	fmt.Println("ResultBF: ", bruteForce(arr, K))
	fmt.Println("ResultB: ", better(arr, K))
	fmt.Println("ResultO: ", optimal(arr, K))
}
