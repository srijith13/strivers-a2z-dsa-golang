package main

import (
	"fmt"
	"math"
)

/*
Given an array that contains both negative and positive integers, find the maximum product subarray.

	Example 1:
	Input: Nums = [1,2,3,4,5,0]
	Output: 120
	Explanation:
	In the given array, 1×2×3×4×5 gives maximum product value

	Example 2:
	Input: Nums = [1,2,-3,0,-4,-5]
	Output: 20
	Explanation:
	In the given array, (-4)×(-5) gives maximum product value.

*/

func bruteForce(arr []int) int {
	n := len(arr)
	maxi := math.MinInt
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			prod := 1
			for k := i; k <= j; k++ {
				prod *= arr[k]
			}
			maxi = max(maxi, prod)
		}
	}

	return maxi
	// TC = O(N3), we check the product of all possible subarrays using 3 nested loops.
	// SC = O(1), No extra space is used.
}

func better(arr []int) int {
	n := len(arr)
	maxi := math.MinInt
	for i := 0; i < n; i++ {
		prod := 1
		for j := i; j < n; j++ {
			prod *= arr[j]
			maxi = max(maxi, prod)
		}
	}

	return maxi
	// TC = O(N2), we check the product of all possible subarrays using 2 nested loops.
	// SC = O(1), No extra space is used.
}

func optimal1(arr []int) int {
	n := len(arr)
	maxi := math.MinInt
	pre, suff := 1, 1

	for i := 0; i < n; i++ {
		if pre == 0 {
			pre = 1
		}
		if suff == 0 {
			suff = 1
		}
		pre *= arr[i]
		suff *= arr[n-i-1]
		maxi = max(maxi, max(pre, suff))
	}

	return maxi
	// TC = O(N), every element of array is visited once.
	// SC = : O(1), constant number of variables are used.
}

func optimal2(arr []int) int {
	n := len(arr)
	res, maxProd, minProd := arr[0], arr[0], arr[0]
	for i := 1; i < n; i++ {
		curr := arr[i]
		if curr < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(curr, maxProd*curr)
		minProd = min(curr, minProd*curr)
		res = max(res, maxProd)
	}

	return res
	// TC = O(N), every element of array is visited once.
	// SC = : O(1), constant number of variables are used.
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 0}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO1: ", optimal1(arr))
	fmt.Println("ResultO2: ", optimal2(arr))
}
