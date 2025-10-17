package main

import (
	"fmt"
)

/*
You have been given a 2-D array 'mat' of size 'N x M' where 'N' and 'M' denote the number of rows and columns, respectively.
The elements of each row are sorted in non-decreasing order. Moreover, the first element of a row is greater than the last element of the previous row (if it exists).
ou are given an integer ‘target’, and your task is to find if it exists in the given 'mat' or not.

	Example 1:
	Input Format: N = 3, M = 4, target = 8,
	mat[] =
	1 2 3 4
	5 6 7 8
	9 10 11 12
	Result: true
	Explanation: The ‘target’  = 8 exists in the 'mat' at index (1, 3).

	Example 2:
	Input Format: N = 3, M = 3, target = 78,
	mat[] =
	1 2 4
	6 7 8
	9 10 34
	Result: false
	Explanation: The ‘target' = 78 does not exist in the 'mat'. Therefore in the output, we see 'false'.

*/

func bruteForce(arr [][]int, target int) bool {
	n, m := len(arr), len(arr[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == target {
				return true
			}
		}
	}
	return false

	// TC =   O(N X M), where N = given row number, M = given column number.
	// Reason: In order to traverse the matrix, we are using nested loops running for n and m times respectively.
	// SC = O(1) as we are not using any extra space.
}

func binarySearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false

}

func better(arr [][]int, target int) bool {
	n, m := len(arr), len(arr[0])

	for i := 0; i < n; i++ {
		if arr[i][0] <= target && target <= arr[i][m-1] {
			return binarySearch(arr[i], target)
		}
	}
	return false

	// TC = O(N + logM), where N = given row number, M = given column number.
	// Reason: We are traversing all rows and it takes O(N) time complexity. But for all rows, we are not applying binary search rather we are only applying it once for a
	// particular row. That is why the time complexity is O(N + logM) instead of O(N*logM).
	// SC = O(1) as we are not using any extra space.
}

func optimal(arr [][]int, target int) bool {
	n, m := len(arr), len(arr[0])
	low, high := 0, n*m-1

	for low <= high {
		mid := (low + high) / 2
		row, col := mid/m, mid%m
		if arr[row][col] == target {
			return true
		} else if arr[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false

	// TC = O(log(NxM)), where N = given row number, M = given column number.
	// Reason: We are applying binary search on the imaginary 1D array of size NxM.
	// SC =  O(1) as we are not using any extra space.
}

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	target := 8
	fmt.Println("ResultBF: ", bruteForce(arr, target))
	fmt.Println("ResultB: ", better(arr, target))
	fmt.Println("ResultO: ", optimal(arr, target))
}
