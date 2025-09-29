package main

import (
	"fmt"
)

/*
Given an integer N and an array of size N-1 containing N-1 numbers between 1 to N. Find the number(between 1 to N), that is not present in the given array.

	Example 1:
	Input Format: N = 5, array[] = {1,2,4,5}
	Result: 3
	Explanation: In the given array, number 3 is missing. So, 3 is the answer.

	Example 2:
	Input Format: N = 3, array[] = {1,3}
	Result: 2
	Explanation: In the given array, number 2 is missing. So, 2 is the answer.

*/

func bruteForce(arr []int, n int) int {
	for i := 1; i < n; i++ {
		flag := 0
		for j := 0; j < n-1; j++ {
			if arr[j] == i {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return i
		}
	}
	return -1
	// TC = O(N2), where N = size of the array+1.
	// SC = O(1) as we are not using any extra space to solve this problem.
}

func better(arr []int, n int) int {
	hash := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		hash[arr[i]]++
	}
	for i := 1; i < n; i++ {
		if hash[i] == 0 {
			return i
		}
	}
	return -1
	// TC =  O(N) + O(N) ~ O(2*N),  where N = size of the array+1. Reason: For storing the frequencies in the hash array,
	// the program takes O(N) time complexity and for checking the frequencies in the second step again O(N) is required. So, the total time complexity is O(N) + O(N).
	// SC =  O(N), where N = size of the array+1. Here we are using an extra hash array of size N+1.
}

func optimal1(arr []int, n int) int {

	summ := (n * (n + 1)) / 2
	s2 := 0
	for i := 0; i < n-1; i++ {
		s2 += arr[i]
	}
	missingNum := summ - s2
	return missingNum
	// TC =  O(N), where N = size of array+1.
	// Reason: Here, we need only 1 loop to get the sum of the array elements. The loop runs for approx. N times. So, the time complexity is O(N)
	// SC =  O(1) as we are not using any extra space.

}

func optimal2(arr []int, n int) int {
	var xor1, xor2 = 0, 0
	for i := 0; i < n-1; i++ {
		xor2 = xor2 ^ arr[i]
		xor1 = xor1 ^ (i + 1)
	}
	xor1 = xor1 ^ n

	return xor1 ^ xor2
	// TC =  O(N), where N = size of array+1.
	// Reason: Here, we need only 1 loop to calculate the XOR. The loop runs for approx. N times. So, the time complexity is O(N).
	// SC =  O(1) as we are not using any extra space.
}

func main() {
	arr1 := []int{1, 2, 4, 5}
	n := 5
	fmt.Println("ResultBF: ", bruteForce(arr1, n))
	fmt.Println("ResultB: ", better(arr1, n))
	fmt.Println("ResultO1: ", optimal1(arr1, n))
	fmt.Println("ResultO2: ", optimal2(arr1, n))

	// no brute force solution for this
}
