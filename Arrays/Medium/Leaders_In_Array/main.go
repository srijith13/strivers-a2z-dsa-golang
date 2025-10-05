package main

import (
	"fmt"
	"math"
	"slices"
)

/*
Given an array, print all the elements which are leaders. A Leader is an element that is greater than all of the elements on its right side in the array.

	Example 1:
	Input: arr = [4, 7, 1, 0]
	Output:7 1 0
	Explanation:
	Rightmost element is always a leader. 7 and 1 are greater than the elements in their right side.

	Example 2:
	Input: arr = [10, 22, 12, 3, 0, 6]
	Output: 22 12 6
	Explanation:
	6 is a leader. In addition to that, 12 is greater than all the elements in its right side (3, 0, 6), also 22 is greater than 12, 3, 0, 6.
*/

func bruteForce(arr []int) []int {
	var ans []int
	for i := 0; i < len(arr); i++ {
		leader := true
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				leader = false
				break
			}
		}
		if leader {
			ans = append(ans, arr[i])
		}
	}
	return ans
	// TC = O(n^2) Since there are nested loops being used, at the worst case n^2 time would be consumed.
	// SC =  O(n) There is no extra space being used in this approach. But, a O(N) of space for ans array will be used in the worst case.
}

func optimal(arr []int) []int {
	var maxi int = math.MinInt
	var ans []int
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > maxi {
			ans = append(ans, arr[i])
		}
		maxi = max(maxi, arr[i])
	}
	// reverse the slice for desc order
	slices.Reverse(ans)
	return ans
	// TC = O(N), where N = size of the array. if sort is added O(nlogn)
	// SC =  O(N) of space for ans array will be used in the worst case.
}

/*
Different approach
	# Last element of an array is always a leader,
    # push into ans array.
    max_elem = arr[n - 1]
    ans.append(arr[n - 1])

    # Start checking from the end whether a number is greater
    # than max no. from right, hence leader.
    for i in range(n - 2, -1, -1):
        if arr[i] > max_elem:
            ans.append(arr[i])
            max_elem = arr[i]
*/

func main() {
	arr := []int{10, 22, 12, 3, 0, 6}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
