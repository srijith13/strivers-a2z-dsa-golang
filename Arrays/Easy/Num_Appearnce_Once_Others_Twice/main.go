package main

import (
	"fmt"
)

/*
Given a non-empty array of integers arr, every element appears twice except for one. Find that single one.
	Example 1:
	Input Format: arr[] = {2,2,1}
	Result: 1
	Explanation: In this array, only the element 1 appear once and so it is the answer.

	Example 2:
	Input Format: arr[] = {4,1,2,1,2}
	Result: 4
	Explanation: In this array, only element 4 appear once and the other elements appear twice. So, 4 is the answer.
*/

func bruteForce(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		num := arr[i]
		cnt := 0
		for j := 0; j < n; j++ {
			if arr[j] == num {
				cnt++
			}
		}
		if cnt == 1 {
			return num
		}
	}
	return -1
	// TC = O(N^2), two loops.
	// SC = O(1) as we are not using any extra space to solve this problem.
}

func better2(arr []int) int {
	n := len(arr)

	hash := make(map[int]int)
	for i := 0; i < n; i++ {
		hash[arr[i]]++
	}
	for i := 0; i < n; i++ {
		if hash[arr[i]] == 1 {
			return arr[i]
		}
	}
	return -1
	// TC = O(N*logM) + O(M), where M = size of the map i.e. M = (N/2)+1. N = size of the array.
	// Reason: We are inserting N elements in the map data structure and insertion takes logM time(where M = size of the map).
	// So this results in the first term O(N*logM). The second term is to iterate the map and search the single element.
	// In Java, HashMap generally takes O(1) time complexity for insertion and search. In that case, the time complexity will be O(N) + O(M).
	// SC = O(M) as we are using a map data structure. Here M = size of the map i.e. M = (N/2)+1.
}

// another better solution is to use an array

func better1(arr []int) int {
	n := len(arr)
	maxi := arr[0]
	for i := 0; i < n; i++ {
		maxi = max(maxi, arr[i])
	}
	hash := make([]int, maxi+1)
	for i := 0; i < n; i++ {
		hash[arr[i]]++
	}
	for i := 0; i < n; i++ {
		if hash[arr[i]] == 1 {
			return arr[i]
		}
	}
	return -1
	// TC = O(N)+O(N)+O(N), where N = size of the array
	// Reason: One O(N) is for finding the maximum, the second one is to hash the elements and the third one is to search the single element in the array.
	// SC = O(maxElement+1) where maxElement = the maximum element of the array.
}

func optimal(arr []int) int {
	n := len(arr)
	var xorr int = 0
	for i := 0; i < n; i++ {
		xorr = xorr ^ arr[i]
	}

	return xorr
	// TC = O(N), where N = size of the array.
	// SC = O(1) as we are not using any extra space.
}

func main() {
	arr1 := []int{4, 1, 2, 1, 2, 4}
	fmt.Println("ResultBF: ", bruteForce(arr1))
	fmt.Println("ResultB1: ", better1(arr1))
	fmt.Println("ResultB2: ", better2(arr1))
	fmt.Println("ResultO: ", optimal(arr1))
}
