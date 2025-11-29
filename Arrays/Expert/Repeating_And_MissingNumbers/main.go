package main

import (
	"fmt"
)

/*
Given an integer array nums of size n containing values from [1, n] and each value appears exactly once in the array, except for A, which appears twice and B which is missing.
Return the values A and B, as an array of size 2, where A appears in the 0-th index and B in the 1st index.
Note: You are not allowed to modify the original array.

	Example 1:
	Input: nums = [3, 5, 4, 1, 1]
	Output: [1, 2]
	Explanation: 1 appears twice in the array, and 2 is missing from the array. So the output is [1, 2].

	Example 2:
	Input: nums = [1, 2, 3, 6, 7, 5, 7]
	Output: [7, 4]
	Explanation: 7 appears twice in the array, and 4 is missing from the array. So the output is [7, 4].
*/

func bruteForce(arr []int) []int {
	n := len(arr)
	repeating, missing := -1, -1
	for i := 1; i <= n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if arr[j] == i {
				cnt++
			}
		}
		if cnt == 2 {
			repeating = i
		} else if cnt == 0 {
			missing = i
		}
		if repeating != -1 && missing != -1 {
			break
		}
	}

	return []int{repeating, missing}
	// TC = O(N2), where N is the size of the array. This is because we are iterating through the array for each integer from 1 to N, leading to a nested loop.
	// SC = O(1), as we are using a constant amount of space for the variables `repeating` and `missing`, regardless of the input size.
}

func better(arr []int) []int {
	n := len(arr)
	hash := make([]int, n+1)
	for i := 0; i < n; i++ {
		hash[arr[i]]++
	}
	repeating, missing := -1, -1
	for i := 1; i < n; i++ {
		if hash[i] == 2 {
			repeating = i
		} else if hash[i] == 0 {
			missing = i
		}
		if repeating != -1 && missing != -1 {
			break
		}
	}

	return []int{repeating, missing}
	// TC =  O(2*N), where N is the size of the array. This is because we are iterating through the array once to build the hash array and then iterating through the hash
	// array to find the repeating and missing numbers.

	// SC = O(N), as we are using an additional hash array of size N+1 to store the frequency of each element.
}

// Using the following Algorithm
// First, calculate the sum of all elements in the given array, denoted as S, and the sum of natural numbers from 1 to N, denoted as Sn. The formula for Sn is (N * (N + 1)) / 2.
// Now calculate S - Sn. This gives us X - Y, where X is the repeating number and Y is the missing number.
// Next, calculate the sum of squares of all elements in the array, denoted as S2, and the sum of squares of the first N numbers, denoted as S2n. The formula for S2n is (N * (N + 1) * (2N + 1)) / 6.
// Now calculate S2 - S2n. This gives us X2 - Y2.
// From the equations S - Sn = X - Y and S2 - S2n = X2 - Y2, we can compute X + Y by the formula (S2 - S2n) / (S - Sn).
// Using the values of X + Y and X - Y, we can solve for X and Y through simple addition and subtraction.
// Finally, return the values of X (the repeating number) and Y (the missing number).
func optimal1(arr []int) []int {
	n := len(arr)
	var sn int64 = int64((n * (n + 1)) / 2)
	var s2n int64 = int64((n * (n + 1) * (2*n + 1)) / 6)
	var s, s2 int64 = 0, 0
	for i := 0; i < n; i++ {
		s += int64(arr[i])
		s2 += int64(arr[i]) * int64(arr[i])
	}
	val := s - sn
	val2 := s2 - s2n
	val2 = val2 / val
	x := (val + val2) / 2
	y := x - val
	return []int{int(x), int(y)}
	// TC = O(N), where N is the size of the array. This is because we are iterating through the array to calculate the sums and sums of squares.

	// SC = : O(1), as we are using a constant amount of space for variables, regardless of the input size.
}

func optimal2(arr []int) []int {
	n := len(arr)
	xr := 0
	for i := 0; i < n; i++ {
		xr ^= arr[i]
		xr ^= (i + 1)
	}
	bitNo := 0
	num := xr & ^(xr - 1)
	zero, one := 0, 0
	for i := 0; i < n; i++ {
		if (arr[i] & (1 << bitNo)) != 0 {
			one ^= arr[i]
		} else {
			zero ^= arr[i]
		}
	}
	for i := 1; i <= n; i++ {
		if (i & num) != 0 {
			one ^= i
		} else {
			zero ^= i
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if arr[i] == zero {
			cnt++
		}
	}
	if cnt == 2 {
		return []int{zero, one}
	}
	return []int{one, zero}
	// TC = O(N), where N is the size of the array. This is because we are iterating through the array to calculate the sums and sums of squares.

	// SC = : O(1), as we are using a constant amount of space for variables, regardless of the input size.
}

func main() {
	arr := []int{3, 5, 4, 1, 1}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO1: ", optimal1(arr))
	fmt.Println("ResultO2: ", optimal2(arr))
}
