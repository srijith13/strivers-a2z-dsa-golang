package main

import (
	"fmt"
)

/*
Given an array of integers, rotating array of elements by k elements either left or right.
	Example 1:
	Input: N = 7, array[] = {1,2,3,4,5,6,7} , k=2 , right
	Output: 6 7 1 2 3 4 5
	Explanation: array is rotated to right by 2 position .

	Example 2:
	Input: N = 6, array[] = {3,7,8,9,10,11} , k=3 , left
	Output: 9 10 11 3 7 8
	Explanation: Array is rotated to right by 3 position.
*/

func rightRotate(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	d = d % n
	if d > n {
		return nil
	}
	temp := make([]int, d)
	for i := n - d; i < n; i++ {
		temp[i-n+d] = arr[i]
	}
	for i := n - d - 1; i >= 0; i-- {
		arr[i+d] = arr[i]
	}
	for i := 0; i < d; i++ {
		arr[i] = temp[i]
	}
	return arr
}
func leftRotate(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	d = d % n
	if d > n {
		return nil
	}
	temp := make([]int, d)
	for i := 0; i < d; i++ {
		temp[i] = arr[i]
	}
	for i := 0; i < n-d; i++ {
		arr[i] = arr[i+d]
	}
	for i := n - d; i < n; i++ {
		arr[i] = temp[i-n+d]
	}
	return arr
}

func bruteForce(arr []int, d int, side string) []int {
	switch side {
	case "right":
		arr = rightRotate(arr, d)
	case "left":
		arr = leftRotate(arr, d)
	}
	return arr
	// TC =  O(n)
	// SC = O(n) as we are using another array of size, same as the given array.
}

func Reverse(arr []int, start, end int) {
	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

// Function to Rotate d elements to right
func rightReverseRotate(arr []int, n, d int) {
	// Reverse first n-d elements
	Reverse(arr, 0, n-d-1)
	// Reverse last d elements
	Reverse(arr, n-d, n-1)
	// Reverse whole array
	Reverse(arr, 0, n-1)
}

func leftReverseRotate(arr []int, n, d int) {
	// Reverse first d elements
	Reverse(arr, 0, d-1)
	// Reverse last n-d elements
	Reverse(arr, d, n-1)
	// Reverse whole array
	Reverse(arr, 0, n-1)
}

func optimal(arr []int, d int, side string) []int {
	switch side {
	case "right":
		rightReverseRotate(arr, len(arr), d)
	case "left":
		leftReverseRotate(arr, len(arr), d)
	}
	return arr
	// TC =  O(n) single traversal
	// SC = O(1)  as no extra space is used
}

func main() {
	arr1 := []int{3, 7, 8, 9, 10, 11}
	d := 3
	fmt.Println("Result: ", bruteForce(arr1, d, "right"))
	d = 5
	fmt.Println("Result: ", optimal(arr1, d, "left"))

}
