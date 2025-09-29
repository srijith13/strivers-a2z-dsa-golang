package main

import (
	"fmt"
	"math"
	"slices"
)

/*
Given an array, find the second smallest and second largest element in the array. Print ‘-1’ in the event that either of them doesn’t exist.
	Example 1:
	Input: [1,2,4,7,7,5]
	Output: Second Smallest : 2
		Second Largest : 5
	Explanation: The elements are as follows 1,2,3,5,7,7 and hence second largest of these is 5 and second smallest is 2

	Example 2:
	Input: [1]
	Output: Second Smallest : -1
		Second Largest : -1
	Explanation: Since there is only one element in the array, it is the largest and smallest element present in the array. There is no second largest or second smallest element present.
*/

func bruteForce(arr []int) {
	n := len(arr)
	if n == 0 || n == 1 {
		fmt.Println(-1, -1)
	}
	slices.Sort(arr)
	fmt.Println("Second Smallest  ", arr[1])
	fmt.Println("Second Largest ", arr[n-2])
	// TC =  O(n log n)
	// SC = O(1)
}

func better(arr []int) {
	n := len(arr)
	var small, secSmall, largest, secLarge = math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	if n == 0 || n == 1 {
		fmt.Println(-1, -1)
	}
	for _, num := range arr {
		small = min(small, num)
		largest = max(largest, num)
	}
	for _, num := range arr {
		if num < secSmall && num != small {
			secSmall = num
		}
		if num > secLarge && num != largest {
			secLarge = num
		}
	}
	fmt.Println("Second Smallest  ", secSmall)
	fmt.Println("Second Largest ", secLarge)
	// TC =  O(n) two traversal
	// SC = O(1)
}

func secSmall(arr []int) int {
	n := len(arr)
	var small, secSmall = math.MaxInt, math.MaxInt
	if n < 2 {
		return -1
	}
	for _, num := range arr {
		if num < small {
			secSmall = small
			small = num
		} else if num < secSmall && num != small {
			secSmall = num
		}
	}
	return secSmall
}

func secLarge(arr []int) int {
	n := len(arr)
	var largest, secLarge = math.MinInt, math.MinInt
	if n < 2 {
		return -1
	}

	for _, num := range arr {
		if num > largest {
			secLarge = largest
			largest = num
		} else if num > secLarge && num != largest {
			secLarge = num
		}
	}
	return secLarge
}

func optimal(arr []int) {
	n := len(arr)
	if n == 0 || n == 1 {
		fmt.Println(-1, -1)
	}
	slices.Sort(arr)
	fmt.Println("Second Smallest  ", secSmall(arr))
	fmt.Println("Second Largest ", secLarge(arr))
	// TC =  O(n) single traversal
	// SC = O(1)
}

func main() {
	arr1 := []int{1, 2, 4, 7, 7, 5}
	bruteForce(arr1)
	better(arr1)
	optimal(arr1)
}
