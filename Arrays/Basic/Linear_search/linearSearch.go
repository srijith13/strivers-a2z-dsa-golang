package main

import "fmt"

/*
Given an array, and an element num the task is to find if num is present in the given array or not. If present print the index of the element or print -1.
	Example 1:
	Input: arr[]= 1 2 3 4 5, num = 3
	Output: 2
	Explanation: 3 is present in the 2nd index

	Example 2:
	Input: arr[]= 5 4 3 2 1, num = 5
	Output: 0
	Explanation: 5 is present in the 0th index
*/

func optimal(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
	// TC =  O(n) single traversal
	// SC = O(1)
}

func main() {
	arr1 := []int{5, 4, 6, 7, 8}
	num := 9
	fmt.Println("Result: ", optimal(arr1, num))
}
