package main

import (
	"fmt"
	"slices"
)

/*
Given an array Arr[] of integers, rearrange the numbers of the given array into the lexicographically next greater permutation of numbers.
If such an arrangement is not possible, it must rearrange to the lowest possible order (i.e., sorted in ascending order).

	Example 1 :
	Input format: Arr[] = {1,3,2}
	Output: Arr[] = {2,1,3}
	Explanation: All permutations of {1,2,3} are {{1,2,3} , {1,3,2}, {2,13} , {2,3,1} , {3,1,2} , {3,2,1}}. So, the next permutation just after {1,3,2} is {2,1,3}.

	Example 2:
	Input format: Arr[] = {3,2,1}
	Output: Arr[] = {1,2,3}
	Explanation: As we see all permutations of {1,2,3}, we find {3,2,1} at the last position. So, we have to return the topmost permutation.
*/

func bruteForce(arr []int) []int {
	return []int{}
	// TC =
	// SC =
}

func optimal(arr []int) []int {
	var n, idx int = len(arr), -1

	for i := n - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			idx = i
			break
		}
	}
	if idx == -1 {
		slices.Reverse(arr)
		return arr
	}

	for i := n - 1; i > idx; i-- {
		if arr[i] > arr[idx] {
			arr[i], arr[idx] = arr[idx], arr[i]\
			break
		}
	}

	slices.Reverse(arr[idx+1:])
	return arr
	// TC = O(3N), where N = size of the given array
	// Finding the break-point, finding the next greater element, and reversal at the end takes O(N) for each, where N is the number of elements in the input array.
	// This sums up to 3*O(N) which is approximately O(3N).
	// SC =  O(N) of space for ans array will be used in the worst case.
}

func main() {
	arr := []int{2, 1, 5, 4, 3, 0, 0}
	// fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
