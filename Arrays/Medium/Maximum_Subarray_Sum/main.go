package main

import (
	"fmt"
	"math"
)

/*
Given an integer array arr, find the contiguous subarray (containing at least one number) which has the largest sum and returns its sum and prints the subarray.

	Example 1:
	Input: arr = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

	Examples 2:
	Input: arr = [1]
	Output: 1
	Explanation: Array has only one element and which is giving positive sum of 1.
*/

func bruteForce(arr []int) int {

	var maxi int = math.MinInt
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			maxi = max(maxi, sum)
		}
	}
	return maxi
	// TC = O(N3), where N = size of the array.
	// Reason: We are using three nested loops, each running approximately N times.
	// SC =  O(1) as we are not using any extra space.
}

func better(arr []int) int {
	var maxi int = math.MinInt
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			maxi = max(maxi, sum)
		}
	}
	return maxi
	// TC = O(N2), where N = size of the array.
	// Reason: We are using three nested loops, each running approximately N times.
	// SC =  O(1) as we are not using any extra space.
}

// Kadane's Algorithm
func optimal(arr []int) int {
	// 	The intuition of the algorithm is not to consider the subarray as a part of the answer if its sum is less than 0. A subarray with a sum less than 0
	// will always reduce our answer and so this type of subarray cannot be a part of the subarray with maximum sum.
	// Here, we will iterate the given array with a single loop and while iterating we will add the elements in a sum variable.
	// Now, if at any point the sum becomes less than 0, we will set the sum as 0 as we are not going to consider any subarray with a negative sum.
	// Among all the sums calculated, we will consider the maximum one.
	// Thus we can solve this problem with a single loop.

	var maxi, sum int = math.MinInt, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		maxi = max(maxi, sum)
		if sum < 0 {
			sum = 0
		}
	}

	// To consider the sum of the empty subarray uncomment the following check:
	// if maxi < 0 {
	// maxi = 0
	// }

	return maxi
	// TC = O(N), where N = size of the given array.
	// SC = O(1) as we are not using any extra space.
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
