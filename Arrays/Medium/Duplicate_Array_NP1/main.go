package main

import (
	"fmt"
	"sort"
)

/*
Given an array of N + 1 size, where each element is between 1 and N. Assuming there is only one duplicate number, your task is to find the duplicate number.

	Example 1:
	Input: arr=[1,3,4,2,2]
	Output: 2
	Explanation: Since 2 is the duplicate number the answer will be 2.

	Example 2:
	Input: [3,1,3,4,2]
	Output:3
	Explanation: Since 3 is the duplicate number the answer will be 3.
*/

func bruteForce(arr []int) int {
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return arr[i]
		}
	}
	return 0
	// TC = O(Nlogn + N), NlogN for sorting the array and O(N) for traversing through the array and checking if adjacent elements are equal or not. But this will distort the array.
	// SC = O(1) as we are not using any extra space.
}

func better(arr []int) int {
	freq := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		if freq[arr[i]] == 0 {
			freq[arr[i]]++
		} else {
			return arr[i]
		}
	}
	return 0
	// TC =  O(N), as we are traversing through the array only once.
	// SC = O(N), as we are using extra space for frequency array.
}

func optimal(arr []int) int {
	slow, fast := arr[0], arr[0]
	for {
		slow = arr[slow]
		fast = arr[arr[fast]]
		if slow == fast {
			break
		}
	}

	fast = arr[0]
	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}

	return slow
	// TC = O(N). Since we traversed through the array only once.
	// SC = O(1) as we are not using any extra space.
}

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
