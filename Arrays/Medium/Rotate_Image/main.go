package main

import (
	"fmt"
	"slices"
)

/*
Given a matrix, your task is to rotate the matrix 90 degrees clockwise.
	Example 1:
	Input: [[1,2,3],[4,5,6],[7,8,9]]
	Output: [[7,4,1],[8,5,2],[9,6,3]]
	Explanation: Rotate the matrix simply by 90 degree clockwise and return the matrix.

	Example 2:
	Input: [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
	Output:[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
	Explanation: Rotate the matrix simply by 90 degree clockwise and return the matrix

Note: Rotate matrix 90 degrees anticlockwise
	Example 1:
	Input:   {{1,2,3},
			{4,5,6},
			{7,8,9}}
	Output:
			3 6 9
			2 5 8
			1 4 7
	Explanation: Rotate the matrix anti-clockwise by 90 degrees and return it.

	Example 2:
	Input:    {{1,2,3,4},
			{5,6,7,8},
			{9,10,11,12},
			{13,14,15,16}}
	Output:
			4 8 12 16
			3 7 11 15
			2 6 10 14
			1 5 9 13
	Explanation: Rotate the matrix anti-clockwise by 90 degrees and return it.

*/

func bruteForce(arr [][]int) [][]int {
	var n int = len(arr)
	ans := make([][]int, n)
	// with extra for loop
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := range n {
		for j := range n {
			ans[j][n-i-1] = arr[i][j]
			// ans[i][j] = arr[j][n-i-1]  // anticlockwise rotation

		}
	}

	// Alternative way without for loop but with extra array
	// for j := 0; j < n; j++ {
	// 	row := make([]int, 0, n)
	// 	for i := n - 1; i >= 0; i-- {
	// 		row = append(row, arr[i][j])
	// 	}
	// 	ans = append(ans, row)
	// }

	return ans
	// TC =  O(N*N) to linearly iterate and put it into some other matrix.
	// SC =  O(N*N) to copy it into some other matrix.
}

func optimal(arr [][]int) {
	var n int = len(arr)

	for i := range n {
		for j := range i {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}

	// for i := range n {   // for clockwise uncomment this
	// 	slices.Reverse(arr[i])
	// }
	slices.Reverse(arr) // anticlockwise rotation add only this and remove the for loop

	// TC = O(N*N) + O(N*N).One O(N*N) is for transposing the matrix and the other is for reversing the matrix.
	// SC =  O(1).
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("ResultBF: ", bruteForce(arr))
	optimal(arr)
	fmt.Println("ResultO: ", arr)
}
