package main

import (
	"fmt"
)

/*
Given an array of N integers, left rotate the array by one place.
	Example 1:
	Input: N = 5, array[] = {1,2,3,4,5}
	Output: 2,3,4,5,1
	Explanation:
	Since all the elements in array will be shifted
	toward left by one so ‘2’ will now become the
	first index and and ‘1’ which was present at
	first index will be shifted at last.


	Example 2:
	Input: N = 1, array[] = {3}
	Output: 3
	Explanation: Here only element is present and so
	the element at first index will be shifted to
	last index which is also by the way the first index.
*/

func bruteForce(arr []int) []int {
	tempArr := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		tempArr[i-1] = arr[i]
	}
	tempArr[len(arr)-1] = arr[0]
	return tempArr
	// TC =  O(n) as we iterate through the array only once.
	// SC = O(n) as we are using another array of size, same as the given array.
}

func optimal(arr []int) []int {
	temp := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
	return arr
	// TC =  O(n) single traversal
	// SC = O(1)  as no extra space is used
}

func main() {
	arr1 := []int{5, 4, 6, 7, 8}
	fmt.Println("Result: ", bruteForce(arr1))
	fmt.Println("Result: ", optimal(arr1))

}
