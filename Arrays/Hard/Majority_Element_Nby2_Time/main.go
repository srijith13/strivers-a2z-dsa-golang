package main

import (
	"fmt"
)

/*
Given an array of N integers, write a program to return an element that occurs more than N/2 times in the given array.
You may consider that such an element always exists in the array.
	Example 1:
	Input Format: N = 3, nums[] = {3,2,3}
	Result: 3
	Explanation: When we just count the occurrences of each number and compare with half of the size of the array, you will get 3 for the above solution.

	Example 2:
	Input Format:  N = 7, nums[] = {2,2,1,1,1,2,2}
	Result: 2
	Explanation: After counting the number of times each element appears and comparing it with half of array size, we get 2 as result.

	Example 3:
	Input Format:  N = 10, nums[] = {4,4,2,4,3,4,4,3,2,4}
	Result: 4
*/

func bruteForce(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[j] == arr[i] {
				count++
			}
		}
		if count > n/2 {
			return arr[i]
		}
	}

	return -1

	// TC = O(N2), where N = size of the given array. Reason: For every element of the array the inner loop runs for N times.
	// And there are N elements in the array. So, the total time complexity is O(N2).
	// SC = O(1) as we are not using any extra space.
}

func better(arr []int) int {
	n := len(arr)
	mpp := make(map[int]int)
	for _, val := range arr {
		mpp[val]++
	}

	for key, val := range mpp {
		if val > (n / 2) {
			return key
		}
	}

	return -1

	// TC = O(N*logN) + O(N), where N = size of the given array.
	// Reason: We are using a map data structure. Insertion in the map takes logN time. And we are doing it for N elements.
	// So, it results in the first term O(N*logN). The second O(N) is for checking which element occurs more than floor(N/2) times.
	// If we use unordered_map instead, the first term will be O(N) for the best and average case and for the worst case, it will be O(N2).
	// SC = O(N) as we are using a map data structure.
}

// Moore’s Voting Algorithm:
func optimal(arr []int) int {
	n, cnt := len(arr), 0
	var el int

	for _, val := range arr {
		if cnt == 0 {
			cnt = 1
			el = val
		} else if val == el {
			cnt++
		} else {
			cnt--
		}

	}
	cnt1 := 0
	for _, val := range arr {
		if val == el {
			cnt1++
		}
	}
	if cnt1 > n/2 {
		return el
	}

	return -1
	// TC =O(N) + O(N), where N = size of the given array.
	// Reason: The first O(N) is to calculate the count and find the expected majority element. The second one is to check if the expected element is the majority one or not.
	// Note: If the question states that the array must contain a majority element, in that case, we do not need the second check. Then the time complexity will boil down to O(N).
	// SC = : O(1) as we are not using any extra space.
}

func main() {
	arr := []int{4, 4, 2, 4, 3, 4, 4, 3, 2, 4}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
