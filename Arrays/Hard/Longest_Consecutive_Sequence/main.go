package main

import (
	"fmt"
	"math"
	"slices"
)

/*
You are given an array of ‘N’ integers. You need to find the length of the longest sequence which contains the consecutive elements.

	Example 1:
	Input: [100, 200, 1, 3, 2, 4]
	Output: 4
	Explanation: The longest consecutive subsequence is 1, 2, 3, and 4.

	Example 2:
	Input: [3, 8, 5, 7, 6]
	Output: 4
	Explanation: The longest consecutive subsequence is 5, 6, 7, and 8.
*/

func linearSearch(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}

func bruteForce(arr []int) int {

	var longest int = 1
	for _, val := range arr {
		x, cnt := val, 1

		for linearSearch(arr, x+1) {
			x++
			cnt++
		}

		longest = max(longest, cnt)
	}

	return longest
	// TC = O(N2), N = size of the given array.
	// Reason: We are using nested loops each running for approximately N times.

	// SC = O(1) as we are using a list that stores a maximum of 2 elements. The space used is so small that it can be considered constant.
}

func better(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	slices.Sort(arr)
	// sort.Ints(arr)
	cnt, lastSmaller, longest := 0, math.MinInt, 1

	for _, val := range arr {
		if val-1 == lastSmaller {
			cnt++
			lastSmaller = val
		} else if val != lastSmaller {
			cnt = 1
			lastSmaller = val
		}
		longest = max(longest, cnt)
	}

	return longest
	// TC =  O(NlogN) + O(N), N = size of the given array.
	// Reason: O(NlogN) for sorting the array. To find the longest sequence, we are using a loop that results in O(N).
	// SC = O(1), as we are not using any extra space to solve this problem.
}

func optimal(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	longest := 1
	st := make(map[int]struct{})
	for _, val := range arr {
		st[val] = struct{}{}
	}

	for key, _ := range st {
		if _, ok := st[key-1]; !ok {
			x, cnt := key, 1

			for {
				if _, ok := st[x+1]; !ok {
					break
				}
				x++
				cnt++
			}

			longest = max(longest, cnt)
		}
	}

	return longest
	// TC =  O(N) + O(2*N) ~ O(3*N), where N = size of the array.
	// Reason: O(N) for putting all the elements into the set data structure. After that for every starting element, we are finding the consecutive elements. T
	// hough we are using nested loops, the set will be traversed at most twice in the worst case. So, the time complexity is O(2*N) instead of O(N2).
	// Note: The time complexity is computed under the assumption that we are using unordered_set and it is taking O(1) for the set operations.
	// If we consider the worst case the set operations will take O(N) in that case and the total time complexity will be approximately O(N2).
	// And if we use the set instead of unordered_set, the time complexity for the set operations will be O(logN) and the total time complexity will be O(NlogN).

	// SC =O(N), as we are using the set data structure to solve this problem.
}

func main() {
	arr := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
