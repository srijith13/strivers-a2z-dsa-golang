package main

import (
	"fmt"
	"math"
)

/*
Given an array of N integers. Find the elements that appear more than N/3 times in the array. If no such element exists, return an empty vector.

	Example 1:
	Input Format: N = 5, array[] = {1,2,2,3,2}
	Result: 2
	Explanation: Here we can see that the Count(1) = 1, Count(2) = 3 and Count(3) = 1.Therefore, the count of 2 is greater than N/3 times. Hence, 2 is the answer.

	Example 2:
	Input Format:  N = 6, array[] = {11,33,33,11,33,11}
	Result: 11 33
	Explanation: Here we can see that the Count(11) = 3 and Count(33) = 3. Therefore, the count of both 11 and 33 is greater than N/3 times. Hence, 11 and 33 is the answer.
*/

func bruteForce(arr []int) []int {
	n := len(arr)
	var ls []int
	for i := 0; i < n; i++ {
		if len(ls) == 0 || ls[0] != arr[i] {
			cnt := 0
			for j := 0; j < n; j++ {
				if arr[j] == arr[i] {
					cnt++
				}
			}
			if cnt > (n / 3) {
				ls = append(ls, arr[i])
			}
		}
		if len(ls) == 2 {
			break
		}
	}

	return ls
	// TC = O(N2), where N = size of the given array.
	// Reason: For every element of the array the inner loop runs for N times. And there are N elements in the array. So, the total time complexity is O(N2).

	// SC = O(1) as we are using a list that stores a maximum of 2 elements. The space used is so small that it can be considered constant.
}

func better(arr []int) []int {
	n := len(arr)
	mini := (n / 3) + 1
	var ls []int
	mpp := make(map[int]int)
	for i := 0; i < n; i++ {
		mpp[arr[i]]++
		if mpp[arr[i]] == mini {
			ls = append(ls, arr[i])
		}
		if len(ls) == 2 {
			break
		}
	}
	return ls
	// TC =  O(N*logN), where N = size of the given array.
	// Reason: We are using a map data structure. Insertion in the map takes logN time. And we are doing it for N elements. So, it results in the first term O(N*logN).
	// If we use unordered_map instead, the first term will be O(N) for the best and average case and for the worst case, it will be O(N2).
	// SC = O(N) as we are using a map data structure. We are also using a list that stores a maximum of 2 elements. That space used is so small that it can be considered constant.
}

func optimal(arr []int) []int {
	n, cnt1, cnt2 := len(arr), 0, 0
	var el1, el2 int = math.MinInt, math.MinInt

	for i := 0; i < n; i++ {
		if cnt1 == 0 && el2 != arr[i] {
			cnt1 = 1
			el1 = arr[i]
		} else if cnt2 == 0 && el1 != arr[i] {
			cnt2 = 1
			el2 = arr[i]
		} else if arr[i] == el1 {
			cnt1++
		} else if arr[i] == el2 {
			cnt2++
		} else {
			cnt1--
			cnt2--
		}

	}
	var ls []int

	cnt1, cnt2 = 0, 0
	for _, val := range arr {
		if val == el1 {
			cnt1++
		}
		if val == el2 {
			cnt2++
		}
	}
	mini := (n / 3) + 1
	if cnt1 >= mini {
		ls = append(ls, el1)
	}
	if cnt2 >= mini {
		ls = append(ls, el2)
	}

	// Uncomment the following line
	// if it is told to sort the answer array:
	// 	slices.Sort(ls) TC --> O(2*log2) ~ O(1);

	return ls
	// TC =  O(N) + O(N), where N = size of the given array.
	// Reason: The first O(N) is to calculate the counts and find the expected majority elements. The second one is to check if the calculated elements are the majority ones or not.
	// SC =  O(1) as we are only using a list that stores a maximum of 2 elements. The space used is so small that it can be considered constant.
}

func main() {
	arr := []int{11, 33, 33, 11, 33, 11}

	fmt.Println("ResultBF: ", bruteForce(arr))
	fmt.Println("ResultB: ", better(arr))
	fmt.Println("ResultO: ", optimal(arr))
}
