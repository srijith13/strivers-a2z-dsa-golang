package main

import (
	"fmt"
)

/*
Given a matrix m X n, count paths from left-top to the right bottom of a matrix with the constraints that from each cell you can either only move to the rightward direction
or the downward direction.

	Example 1:
	Input Format: m = 2, n= 2
	Output: 2
	Explanation: From the top left corner there are total 2 ways to reach the bottom right corner:
	Step 1: Right ->> Down
	Step 2: Down ->> Right

	Example 2:
	Input Format: m = 2, n= 3
	Output: 3
	Explanation:  From the top left corner there is a total of 3 ways to reach the bottom right corner.
	Step 1: Right ->> Right ->> Down
	Step 2: Right ->> Down ->> Right
	Step 3: Down ->> Right->> Right
*/

func countPaths(i, j, m, n int) int {
	if i == n-1 && j == m-1 {
		return 1
	}
	if i >= n || j >= m {
		return 0
	} else {
		return countPaths(i+1, j, m, n) + countPaths(i, j+1, m, n)
	}
}

func bruteForce(m, n int) int {
	return countPaths(0, 0, m, n)
	// TC = O(2^n) ,complexity  of this recursive solution is exponential.

	// SC = O(2^n) As we are using stack space for recursion so the space complexity will also be exponential.
}

func countPaths2(i, j, m, n int, dp [][]int) int {
	if i == n-1 && j == m-1 {
		return 1
	}
	if i >= n || j >= m {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	} else {
		dp[i][j] = countPaths2(i+1, j, m, n, dp) + countPaths2(i, j+1, m, n, dp)
		return dp[i][j]
	}
}

func better(m, n int) int {
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	num := countPaths2(0, 0, m, n, dp)
	if m == 1 && n == 1 {
		return num
	}

	return dp[0][0]
	// TC =  O(n*m) + O(nm), because at max there can be n*m number of states and a for loop for initializing to -1.
	// Reason: We are using two nested loops here. As each of them is running for exactly N times, the time complexity will be approximately O(N2).
	// SC = O(n*m) As we are using extra space for the dummy matrix
}

func optimal(m, n int) int {

	N := n + m - 2
	r := m - 1
	var res int64 = 1
	// res := 1
	for i := 1; i <= r; i++ {
		res = res * int64((N - r + i)) / int64(i)
		// res = (res * (N - r + i)) / i

	}
	return int(res)
	// return res

	// TC = O(n-1) or  O(m-1) depending on the formula we are using.
	// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N) but if we are using a map data structure,
	// the time complexity will be O(N*logN). The least complexity will be O(N) as we are using a loop to traverse the array.

	// SC = O(1) as we are not using any extra space.
}

func main() {
	m, n := 3, 7
	fmt.Println("ResultBF: ", bruteForce(m, n))
	fmt.Println("ResultB: ", better(m, n))
	fmt.Println("ResultO: ", optimal(m, n))
}
