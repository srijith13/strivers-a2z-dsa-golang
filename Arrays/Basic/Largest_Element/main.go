package main

import (
	"fmt"
	"slices"
)

/*
 Given an array, we have to find the largest element in the array.
 	Example 1:
	Input: arr[] = {2,5,1,3,0};
	Output: 5
	Explanation: 5 is the largest element in the array.

	Example2:
	Input: arr[] = {8,10,5,7,9};
	Output: 10
	Explanation: 10 is the largest element in the array.
*/

func bruteForce(arr []int) int {
	slices.Sort(arr)
	return arr[len(arr)-1]
	// TC =  O(n log n)
	// SC = O(1)
}

func optimal(arr []int) int {
	largest := arr[0]
	for _, i := range arr {
		if i > largest {
			largest = i
		}
	}
	return largest
	// TC =  O(n)
	// SC = O(1)
}

func main() {
	arr1 := []int{8, -11, 5, 7, 9}
	fmt.Println("bruteForce ", bruteForce(arr1))
	fmt.Println("optimal ", optimal(arr1))
}
