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

func better(arr []int) (ans [][]int) {
	n := len(arr)
	st := make(map[string]struct{})
	for i := 0; i < n; i++ {
		hashSet := make(map[int]struct{})
		for j := i + 1; j < n; j++ {
			third := -(arr[i] + arr[j])
			if _, ok := hashSet[third]; ok {
				temp := []int{arr[i], arr[j], third}
				sort.Ints(temp)
				key := fmt.Sprintf("%d,%d,%d", temp[0], temp[1], temp[2])
				st[key] = struct{}{}
			}
			hashSet[arr[j]] = struct{}{}
		}
	}
	for key, _ := range st {
		var a, b, c int
		fmt.Sscanf(key, "%d,%d,%d", &a, &b, &c)
		ans = append(ans, []int{a, b, c})
	}

	return ans
	// TC =  O(N2 * log(no. of unique triplets)), as we are mainly using 3 nested loops. And inserting triplets into the set takes
	// O(log(no. of unique triplets)) time complexity. But we are not considering the time complexity of sorting as we are just sorting 3 elements every time.

	// SC = O(2 * no. of the unique triplets) + O(N) as we are using a set data structure and a list to store the triplets and extra
	// O(N) for storing the array elements in another set.
}

func optimal(arr []int) (ans [][]int) {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		j, k := i+1, len(arr)-1 // left, right := i+1, len(arr)-1
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				ans = append(ans, []int{arr[i], arr[j], arr[k]})
				j++
				k--
				for j < k && arr[j] == arr[j-1] {
					j++
				}
				for j < k && arr[k] == arr[k+1] {
					k--
				}
			}
		}
	}

	return ans
	// TC = O(NlogN)+O(N2), as The pointer i, is running for approximately N times. And both the pointers j and k combined can run for approximately N times
	// including the operation of skipping duplicates. So the total time complexity will be O(N2).

	// SC = : O(no. of quadruplets), This space is only used to store the answer. We are not using any extra space to solve this problem.
	// So, from that perspective, space complexity can be written as O(1).
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
