package main

import (
	"fmt"
	"sort"
)

/*
Given an array of N integers, your task is to find unique triplets that add up to give a sum of zero. In short, you need to return an array
of all the unique triplets [arr[a], arr[b], arr[c]] such that i!=j, j!=k, k!=i, and their sum is equal to zero.

	Example 1:
	Input: nums = [-1,0,1,2,-1,-4]
	Output: [[-1,-1,2],[-1,0,1]]
	Explanation: Out of all possible unique triplets possible, [-1,-1,2] and [-1,0,1] satisfy the condition of summing up to zero with i!=j!=k


	Example 2:
	Input: nums=[-1,0,1,0]
	Output: Output: [[-1,0,1],[-1,1,0]]
	Explanation: Out of all possible unique triplets possible, [-1,0,1] and [-1,1,0] satisfy the condition of summing up to zero with i!=j!=k
*/

func bruteForce(arr []int) (ans [][]int) {
	n := len(arr)
	st := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					temp := []int{arr[i], arr[j], arr[k]}
					sort.Ints(temp)
					key := fmt.Sprintf("%d,%d,%d", temp[0], temp[1], temp[2])
					st[key] = struct{}{}
				}
			}
		}
	}
	for key, _ := range st {
		var a, b, c int
		fmt.Sscanf(key, "%d,%d,%d", &a, &b, &c)
		ans = append(ans, []int{a, b, c})
	}

	return ans
	// TC = O(N3 * log(no. of unique triplets)) or O(N3 * N * log(m)), where N = size of the array.
	// Reason: Here, we are mainly using 3 nested loops. And inserting triplets into the set takes O(log(no. of unique triplets)) time complexity.
	// But we are not considering the time complexity of sorting as we are just sorting 3 elements every time.

	// SC = O(2 * no. of the unique triplets) as we are using a set data structure and a list to store the triplets.
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println("ResultBF: ", bruteForce(arr))

}
