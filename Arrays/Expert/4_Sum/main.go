package main

import (
	"fmt"
	"sort"
)

/*
Given an array of N integers, your task is to find unique quads that add up to give a target value.
In short, you need to return an array of all the unique quadruplets [arr[a], arr[b], arr[c], arr[d]] such that their sum is equal to a given target.
Note: a, b, c and d are also distinct and lies between 0 to n-1 (both inclusive).

	Example 1:
	Input Format:arr[] = [1,0,-1,0,-2,2], target = 0
	Result: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	Explanation:We have to find unique quadruplets from the array such that the sum of those elements is equal to the target sum given that is 0. The result obtained is such that the sum of the quadruplets yields 0.

	Example 2:
	Input Format: arr[] = [4,3,3,4,4,2,1,2,1,1], target = 9
	Result: [[1,1,3,4],[1,2,2,4],[1,2,3,3]]
	Explanation: The sum of all the quadruplets is equal to the target i.e. 9.
*/

func bruteForce(arr []int, target int64) (ans [][]int) {
	n := len(arr)
	st := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					var sum int64 = int64(arr[i] + arr[j])
					sum += int64(arr[k])
					sum += int64(arr[l])
					if sum == target {
						temp := []int{arr[i], arr[j], arr[k], arr[l]}
						sort.Ints(temp)
						key := fmt.Sprintf("%d,%d,%d,%d", temp[0], temp[1], temp[2], temp[3])
						st[key] = struct{}{}
					}
				}
			}
		}
	}
	for key, _ := range st {
		var a, b, c, d int
		fmt.Sscanf(key, "%d,%d,%d,%d", &a, &b, &c, &d)
		ans = append(ans, []int{a, b, c, d})
	}

	return ans
	// TC =O(N4 * log(no. of unique quads)) or O(N4), where N = size of the array.
	// Reason: Here, we are mainly using 4 nested loops. And inserting triplets into the set takes O(log(no. of unique quads)) time complexity.
	// But we are not considering the time complexity of sorting as we are just sorting 4 elements every time.

	// SC = O(2 * no. of the unique quads) as we are using a set data structure and a list to store the quads.
}

func better(arr []int, target int64) (ans [][]int) {
	n := len(arr)
	st := make(map[string]struct{})
	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {
			hashSet := make(map[int]struct{})
			for k := j + 1; k < n; k++ {
				var sum int64 = int64(arr[i] + arr[j])
				sum += int64(arr[k])
				fourth := int(target - sum)
				if _, ok := hashSet[fourth]; ok {
					temp := []int{arr[i], arr[j], arr[k], fourth}
					sort.Ints(temp)
					key := fmt.Sprintf("%d,%d,%d,%d", temp[0], temp[1], temp[2], temp[3])
					st[key] = struct{}{}
				}
				hashSet[arr[k]] = struct{}{}
			}
		}
	}
	for key, _ := range st {
		var a, b, c, d int
		fmt.Sscanf(key, "%d,%d,%d,%d", &a, &b, &c, &d)
		ans = append(ans, []int{a, b, c, d})
	}

	return ans
	// TC =  O(N3*log(M)), as we are mainly using 3 nested loops, and inside the loops there are some operations on the set data structure which take log(M) time complexity.

	// SC = O(2 * no. of the unique quadruplets) + O(N) as we are using a set data structure and a list to store the quads. This results in the first term.
	// And the second space is taken by the set data structure we are using to store the array elements. At most, the set can contain approximately all the array elements
	// and so the space complexity is O(N).
}

func optimal(arr []int, target int64) (ans [][]int) {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if j != i+1 && arr[j] == arr[j-1] {
				continue
			}
			left, right := j+1, len(arr)-1 // left, right := i+1, len(arr)-1
			for left < right {
				var sum int64 = int64(arr[i] + arr[j])
				sum += int64(arr[left])
				sum += int64(arr[right])
				if sum == target {
					ans = append(ans, []int{arr[i], arr[j], arr[left], arr[right]})
					left++
					right--
					for left < right && arr[left] == arr[left-1] {
						left++
					}
					for left < right && arr[right] == arr[right+1] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return ans
	// TC = O(N3), as Each of the pointers i and j, is running for approximately N times. And both the pointers k and l combined can run for approximately
	// N times including the operation of skipping duplicates. So the total time complexity will be O(N3).

	// SC = : O(no. of quadruplets), as This space is only used to store the answer. We are not using any extra space to solve this problem.
	// So, from that perspective, space complexity can be written as O(1).
}

func main() {
	arr := []int{1, 0, -1, 0, -2, 2}
	var target int64 = 0
	fmt.Println("ResultBF: ", bruteForce(arr, target))
	fmt.Println("ResultB: ", better(arr, target))
	fmt.Println("ResultO: ", optimal(arr, target))
}
