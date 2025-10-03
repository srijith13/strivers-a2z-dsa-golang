package main

import (
	"fmt"
)

/*
Variety-1

There’s an array ‘A’ of size ‘N’ with an equal number of positive and negative elements. Without altering the relative order of positive and negative elements,
you must return an array of alternately positive and negative values.
Note: Start the array with positive elements.

	Example 1:
	Input:
	arr[] = {1,2,-4,-5}, N = 4
	Output: 1 -4 2 -5
	Explanation:
	Positive elements = 1,2
	Negative elements = -4,-5
	To maintain relative ordering, 1 must occur before 2, and -4 must occur before -5.

	Example 2:
	Input:
	arr[] = {1,2,-3,-1,-2,-3}, N = 6
	Output: 1 -3 2 -1 3 -2
	Explanation:
	Positive elements = 1,2,3
	Negative elements = -3,-1,-2
	To maintain relative ordering, 1 must occur before 2, and 2 must occur before 3.
	Also, -3 should come before -1, and -1 should come before -2.

Variety-2

Problem Statement:
There’s an array ‘A’ of size ‘N’ with positive and negative elements (not necessarily equal). Without altering the relative order of positive and negative elements,
you must return an array of alternately positive and negative values. The leftover elements should be placed at the very end in the same order as in array A.
Note: Start the array with positive elements.

	Example 1:
	Input:
	arr[] = {1,2,-4,-5,3,4}, N = 6
	Output:1 -4 2 -5 3 4
	Explanation:
	Positive elements = 1,2
	Negative elements = -4,-5
	To maintain relative ordering, 1 must occur before 2, and -4 must occur before -5.
	Leftover positive elements are 3 and 4 which are then placed at the end of the array.

	Example 2:
	Input:
	arr[] = {1,2,-3,-1,-2,-3}, N = 6
	Output:1 -3 2 -1 3 -2
	Explanation:
	Positive elements = 1,2
	Negative elements = -3,-1,-2,-4
	To maintain relative ordering, 1 must occur before 2.
	Also, -3 should come before -1, and -1 should come before -2.
	After alternate ordering, -2 and -4 are left, which would be placed at the end of the ans array.
*/

func bruteForce1(arr []int) []int {
	var pos, neg []int

	for _, val := range arr {
		if val >= 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[2*i] = pos[i]
		arr[2*i+1] = neg[i]
	}
	return arr
	// TC = O(N+N/2) of running time due to multiple traversals which we’ll try to optimize in the optimized approach given below.
	// SC =  O(N) as we are using extra space for arrays which adds upto len of original array.
}

func optimal1(arr []int) []int {
	ans := make([]int, len(arr))
	posIdx, negIdx := 0, 1
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			ans[negIdx] = arr[i]
			negIdx += 2
		} else {
			ans[posIdx] = arr[i]
			posIdx += 2
		}
	}
	return ans
	// TC = O(N), where N = size of the array.
	// Reason: We are using three nested loops, each running approximately N times.
	// SC =  O(N) as we are using an extra space.
}

func bruteForce2(arr []int) []int {
	var pos, neg []int

	for _, val := range arr {
		if val >= 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}

	if len(pos) < len(neg) {
		for i := 0; i < len(pos); i++ {
			arr[2*i] = pos[i]
			arr[2*i+1] = neg[i]
		}
		idx := len(pos) * 2
		for i := len(pos); i < len(neg); i++ {
			arr[idx] = neg[i]
			idx++
		}
	} else {
		for i := 0; i < len(neg); i++ {
			arr[2*i] = pos[i]
			arr[2*i+1] = neg[i]
		}
		idx := len(neg) * 2
		for i := len(neg); i < len(pos); i++ {
			arr[idx] = pos[i]
			idx++
		}
	}

	return arr
	// TC = O(2*N)  The worst case complexity is O(2*N) which is a combination of O(N) of traversing the array to segregate into neg and pos array
	// and O(N) for adding the elements alternatively to the main array
	// SC =  O(N) as we are using extra space for arrays which adds upto len of original array.
}

func main() {
	arr := []int{1, 2, -4, -5}
	fmt.Println("ResultBF: ", bruteForce1(arr))
	fmt.Println("ResultO: ", optimal1(arr))
	arr1 := []int{1, 2, -3, -1, -2, -3}
	fmt.Println("ResultBF2: ", bruteForce2(arr1))
}
