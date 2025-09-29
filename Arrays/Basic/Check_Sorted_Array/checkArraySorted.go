package main

import (
	"fmt"
)

/*
Given an array of size n, write a program to check if the given array is sorted in (ascending / Increasing / Non-decreasing) order or not. If the array is sorted then return True, Else return False.

Note: Two consecutive equal values are considered to be sorted.
	Example 1:
	Input: N = 5, array[] = {1,2,3,4,5}
	Output: True.
	Explanation: The given array is sorted i.e Every element in the array is smaller than or equals to its next values, So the answer is True.


	Example 2:
	Input: N = 5, array[] = {5,4,6,7,8}
	Output: False.
	Explanation: The given array is Not sorted i.e Every element in the array is not smaller than or equal to its next values, So the answer is False.

	Here element 5 is not smaller than or equal to its future elements.

*/

// Brute Force is to use 2 separate loops and compare and give the result

func optimal(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
	// TC =  O(n) single traversal
	// SC = O(1)
}

func main() {
	arr1 := []int{5, 4, 6, 7, 8}
	fmt.Println("Result: ", optimal(arr1))
}
