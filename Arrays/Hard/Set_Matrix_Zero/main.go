package main

import (
	"fmt"
)

/*
Given a matrix if an element in the matrix is 0 then you will have to set its entire column and row to 0 and then return the matrix.

	Examples 1:
	Input: matrix=[[1,1,1],[1,0,1],[1,1,1]]
	Output: [[1,0,1],[0,0,0],[1,0,1]]
	Explanation: Since matrix[2][2]=0.Therfore the 2nd column and 2nd row wil be set to 0.

	Input: matrix=[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	Output:[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
	Explanation:Since matrix[0][0]=0 and matrix[0][3]=0. Therefore 1st row, 1st column and 4th column will be set to 0
*/

func markRow(arr [][]int, row, i int) {
	for j := 0; j < row; j++ {
		if arr[i][j] != 0 {
			arr[i][j] = -1
		}
	}
}

func markCol(arr [][]int, col, j int) {
	for i := 0; i < col; i++ {
		if arr[i][j] != 0 {
			arr[i][j] = -1
		}
	}
}

func bruteForce(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				markRow(arr, m, j)
				markCol(arr, n, i)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == -1 {
				arr[i][j] = 0
			}
		}
	}
	return arr
	// TC = O((N*M)*(N + M)) + O(N*M), where N = no. of rows in the matrix and M = no. of columns in the matrix.
	// Reason: Firstly, we are traversing the matrix to find the cells with the value 0. It takes O(N*M). Now, whenever we find any such cell we mark that row and column with -1. This process takes O(N+M). So, combining this the whole process, finding and marking, takes O((N*M)*(N + M)).
	// Another O(N*M) is taken to mark all the cells with -1 as 0 finally.

	// SC = O(1) as we are not using any extra space.
}

func better(arr [][]int) [][]int {
	var row, col []int = make([]int, len(arr)), make([]int, len(arr[0]))
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i] == 1 || col[j] == 1 {
				arr[i][j] = 0
			}
		}
	}
	return arr
	// TC =  O(2*(N*M)), where N = no. of rows in the matrix and M = no. of columns in the matrix.
	// Reason: In this approach, we are also traversing the entire matrix 2 times and each traversal is taking O(N*M) time complexity.
	// SC = O(1) as we are not using any extra space.
}

func optimal(arr [][]int) [][]int {
	n, m, col0 := len(arr), len(arr[0]), 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				arr[i][0] = 0

				// mark j-th column:
				if j != 0 {
					arr[0][j] = 0
				} else {
					col0 = 0
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if arr[i][j] != 0 {
				if arr[i][0] == 0 || arr[0][j] == 0 {
					arr[i][j] = 0
				}
			}
		}
	}

	if arr[0][0] == 0 {
		for j := 0; j < m; j++ {
			arr[0][j] = 0
		}
	}
	if col0 == 0 {
		for i := 0; i < n; i++ {
			arr[i][0] = 0
		}
	}

	return arr
	// TC = O(3N), where N = size of the given array
	// Finding the break-point, finding the next greater element, and reversal at the end takes O(N) for each, where N is the number of elements in the input array.
	// This sums up to 3*O(N) which is approximately O(3N).
	// SC =  O(N) of space for ans array will be used in the worst case.
}

func main() {
	arr := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	arr1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	arr2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr1))
	fmt.Println("ResultO: ", optimal(arr2))
}
