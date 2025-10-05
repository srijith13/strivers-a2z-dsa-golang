package main

import (
	"fmt"
)

/*
Given a Matrix, print the given matrix in spiral order.

	Example 1:
	Input: Matrix[][] = { { 1, 2, 3, 4 },
		      			  { 5, 6, 7, 8 },
		      			  { 9, 10, 11, 12 },
	              		  { 13, 14, 15, 16 } }

	Output: 1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10.
	Explanation: The output of matrix in spiral form.

	Example 2:
	Input: Matrix[][] = { { 1, 2, 3 },
						  { 4, 5, 6 },
						  { 7, 8, 9 } }

	Output: 1, 2, 3, 6, 9, 8, 7, 4, 5.
	Explanation: The output of matrix in spiral form.
*/

func bruteForce(arr [][]int) []int {
	var n, m int = len(arr), len(arr[0])
	var left, right, top, bottom int = 0, m - 1, 0, n - 1
	var ans []int
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans = append(ans, arr[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			ans = append(ans, arr[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, arr[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, arr[i][left])
			}
			left++
		}

	}
	return ans
	// TC = O(m x n) Since all the elements are being traversed once and there are total n x m elements ( m elements in each row and total n rows) so the time complexity will be O(n x m).
	// SC =  O(n) Extra Space used for storing traversal in the ans array.
}

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("ResultBF: ", bruteForce(arr))
}
