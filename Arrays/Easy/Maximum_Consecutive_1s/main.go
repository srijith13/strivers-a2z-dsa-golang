package main

import "fmt"

/*
Given an array that contains only 1 and 0 return the count of maximum consecutive ones in the array.
	Example 1:
	Input: prices = {1, 1, 0, 1, 1, 1}
	Output: 3
	Explanation: There are two consecutive 1’s and three consecutive 1’s in the array out of which maximum is 3.
	Input: prices = {1, 0, 1, 1, 0, 1}
	Output: 2
	Explanation: There are two consecutive 1's in the array.
*/

func optimal(arr []int) int {
	var maxi, cnt = 0, 0
	for _, val := range arr {
		if val == 1 {
			cnt++
			maxi = max(maxi, cnt)
		} else {
			cnt = 0
		}
	}
	return maxi
	// TC = O(N) since the solution involves only a single pass.
	// SC = O(1) as we are not using any extra space to solve this problem.
}

func main() {
	arr1 := []int{1, 1, 0, 1, 1, 1}
	fmt.Println("Result: ", optimal(arr1))

	// no brute force solution for this
}
