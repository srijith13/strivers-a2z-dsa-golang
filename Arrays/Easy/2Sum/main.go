package main

import (
	"fmt"
	"sort"
)

/*
Given an array of integers arr[] and an integer target.
1st variant: Return YES if there exist two numbers such that their sum is equal to the target. Otherwise, return NO.
2nd variant: Return indices of the two numbers such that their sum is equal to the target. Otherwise, we will return {-1, -1}.
Note: You are not allowed to use the same element twice. Example: If the target is equal to 6 and num[1] = 3, then nums[1] + nums[1] = target is not a solution.

	Example 1:
	Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 14
	Result: YES (for 1st variant)
		[1, 3] (for 2nd variant)
	Explanation: arr[1] + arr[3] = 14. So, the answer is “YES” for the first variant and [1, 3] for 2nd variant.

	Example 2:
	Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 15
	Result: NO (for 1st variant)
		[-1, -1] (for 2nd variant)
	Explanation: There exist no such two numbers whose sum is equal to the target.
*/

func bruteForce(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i+1; j++ {
			if arr[i]+arr[j] == target {
				fmt.Println("BF YES")
				return []int{i, j}
			}
		}
	}
	fmt.Println("BF NO")
	return []int{-1, -1}
	// TC = O(N^2), two loops.
	// SC = O(1) as we are not using any extra space to solve this problem.
}

func better(arr []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		more := target - num
		if val, ok := hash[more]; ok {
			fmt.Println("BT YES")
			return []int{i, val}
		}
		hash[arr[i]] = i
	}
	fmt.Println("BT NO")
	return []int{-1, -1}
	// TC =O(N), where N = size of the array.
	// Reason: The loop runs N times in the worst case and searching in a hashmap takes O(1) generally. So the time complexity is O(N).
	// Note: In the worst case(which rarely happens), the unordered_map takes O(N) to find an element. In that case, the time complexity will be O(N2).
	// If we use map instead of unordered_map, the time complexity will be O(N* logN) as the map data structure takes logN time to find an element.
	// SC = O(N) as we use the map data structure.
}

// not suitable if we need to return index
func optimal(arr []int, target int) string {
	sort.Ints(arr)
	var left, right int = 0, len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			fmt.Println("OP YES")
			return "Yes"
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	fmt.Println("OP NO")
	return "No"
	// TC = O(N) + O(N*logN), where N = size of the array.
	// Reason: The loop will run at most N times. And sorting the array will take N*logN time complexity.
	// SC = O(1) as we are not using any extra space.
}

func main() {
	arr1 := []int{2, 6, 5, 8, 11}
	target := 14
	fmt.Println("ResultBF: ", bruteForce(arr1, target))
	fmt.Println("ResultB1: ", better(arr1, target))

	fmt.Println("ResultB1: ", optimal(arr1, target)) // not suitable if we need to return index
}
