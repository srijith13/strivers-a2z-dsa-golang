package main

import (
	"fmt"
)

/*
Variation 1: Given row number r and column number c. Print the element at position (r, c) in Pascal’s triangle.
 The element at position (r,c) = (5,3) is: 6
*/

func nCr(n, r int) int64 {
	var res int64 = 1

	// calculating nCr:
	for i := 0; i < r; i++ {
		res = res * int64((n - i))
		res = res / int64((i + 1))
	}
	return res
}

func optimalV1(r, c int) int64 {
	element := nCr(r-1, c-1)
	return element
	// TC = O(n) or O(c), where c = given column number.Reason: We are running a loop for r times, where r is c-1.
	// SC = O(1) as we are not using any extra space.
}

/*
Given the row number n. Print the n-th row of Pascal’s triangle.
Our first observation regarding Pascal’s triangle should be that the n-th row of the triangle has exactly n elements. With this observation, we will proceed to solve this problem.

	Example
	Input = 5
	Output = 1 4 6 4 1
*/

func bruteForceV2(n int) {
	fmt.Printf("ResultBF V2: ")
	for c := 1; c <= n; c++ {
		fmt.Printf("%d ", nCr(n-1, c-1))
	}
	fmt.Printf("\n")
	// TC = O(n^2) orO(n*r), where n is the given row number, and r is the column index which can vary from 0 to n-1.
	// Reason: We are calculating the element for each column. Now, there are total n columns, and for each column,
	// the calculation of the element takes O(r) time where r is the column index.
	// SC = O(1) as we are not using any extra space.
}

func optimalV2(n int) {
	var res int64 = 1

	fmt.Printf("ResultO V2: %d ", res)
	for i := 1; i < n; i++ {
		res = res * int64((n - i))
		res = res / int64((i))
		fmt.Printf("%d ", res)
	}
	fmt.Printf("\n")

	// TC = O(N) where N = given row number. Here we are using only a single loop.
	// SC = O(1) as we are not using any extra space.
}

/*
Given the row number n. Print the Pascal’s triangle till the nth row.

	Example
	Input = 5
	Output =
	1
	1 1
	1 2 1
	1 3 3 1
	1 4 6 4 1

*/

func bruteForceV3(n int) [][]int64 {
	// combination of variation 1 and variation 2
	var ans [][]int64
	for row := 1; row <= n; row++ {
		var temp []int64
		for col := 1; col <= row; col++ {
			temp = append(temp, nCr(row-1, col-1))
		}
		ans = append(ans, temp)
	}
	return ans
	// TC = O(n*n*r) ~ O(n3), where n = number of rows, and r = column index.
	// Reason: The row loop will run for approximately n times. And generating a row using the naive approach of variation 2 takes O(n*r) time complexity.
	// SC = O(n)  In this case, we are only using space to store the answer. That is why space complexity can be still considered as O(1).
}

func generateRow(row int) []int64 {
	var ans int64 = 1
	var ansRow []int64
	ansRow = append(ansRow, 1) //inserting the 1st element

	for col := 1; col < row; col++ {
		ans = ans * int64((row - col))
		ans = ans / int64((col))
		ansRow = append(ansRow, ans)
	}

	return ansRow
}

func optimalV3(n int) [][]int64 {
	var ans [][]int64
	for row := 1; row <= n; row++ {
		ans = append(ans, generateRow(row))
	}

	return ans
	// TC = O(n2), where n = number of rows(given).
	// Reason: We are generating a row for each single row. The number of rows is n. And generating an entire row takes O(n) time complexity.
	// SC = O(n)  In this case, we are only using space to store the answer. That is why space complexity can be still considered as O(1).
}

func main() {
	// fmt.Println("ResultBF V1: ", bruteForceV1(5, 3)) find using the extended nCr formula n!, r!, (n-r)!
	fmt.Println("ResultO V1: ", optimalV1(5, 3))

	bruteForceV2(5)
	optimalV2(5)

	fmt.Println("ResultBF V3: ", bruteForceV3(5))
	fmt.Println("ResultO V3: ", optimalV3(5))

}
