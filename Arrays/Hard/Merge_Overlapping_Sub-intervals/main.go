package main

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals and return an array of the non-overlapping intervals that cover
all the intervals in the input.

	Example 1:
	Input : intervals=[[1,3],[2,6],[8,10],[15,18]]
	Output : [[1,6],[8,10],[15,18]]
	Explanation : Since intervals [1,3] and [2,6] are overlapping we can merge them to form [1,6] intervals.

	Example 2:
	Input : [[1,4],[4,5]]
	Output :  [[1,5]]
	Explanation :  Since intervals [1,4] and [4,5] are overlapping we can merge them to form [1,5].
*/

func bruteForce(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	var ans [][]int
	for i := 0; i < len(arr); i++ {
		start, end := arr[i][0], arr[i][1]
		if len(ans) != 0 && end <= ans[len(ans)-1][1] {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if arr[j][0] <= end {
				end = max(end, arr[j][1])
			} else {
				break
			}
		}
		ans = append(ans, []int{start, end})

	}
	return ans
	// TC = O(NlogN + N^2), we sort the entire array and then merge them in a single pass.
	// SC = O(N), additional space used to store the non-overlapping intervals.
}

func bruteForce2(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	var ans [][]int

	n := len(arr)
	i := 0
	for i < n {
		start, end := arr[i][0], arr[i][1]
		j := i + 1
		for j < len(arr) && arr[j][0] <= end {
			if arr[j][1] > end {
				end = arr[j][1]
			}
			j++
		}
		ans = append(ans, []int{start, end})
		i = j
	}
	return ans
	// TC = O(NlogN + N^2), we sort the entire array and then merge them in a single pass.
	// SC = O(N), additional space used to store the non-overlapping intervals.
}

func optimal(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	var ans [][]int
	for _, interval := range arr {
		if len(ans) == 0 || ans[len(ans)-1][1] < interval[0] {
			// Append current interval as a new one
			ans = append(ans, interval)
		} else {
			// Overlapping: merge by extending the end
			if ans[len(ans)-1][1] < interval[1] {
				ans[len(ans)-1][1] = interval[1]
			}
		}
	}
	return ans
	// TC = O(NlogN N) we sort the entire array and then merge them in a single pass.
	// SC = O(N), additional space used to store the non-overlapping intervals.
}

func main() {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultBF2: ", bruteForce2(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
