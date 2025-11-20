package main

import (
	"fmt"
	"sort"
)

/*
Given two sorted integer arrays nums1 and nums2, merge both the arrays into a single array sorted in non-decreasing order.
The final sorted array should be stored inside the array nums1 and it should be done in-place.
Array nums1 has a length of m + n, where the first m elements denote the elements of nums1 and rest are 0s whereas nums2 has a length of n.

	Example 1:
	Input : nums1 = [-5, -2, 4, 5, 0, 0, 0], nums2 = [-3, 1, 8]
	Output : [-5, -3, -2, 1, 4, 5, 8]
	Explanation : The merged array is: [-5, -3, -2, 1, 4, 5, 8], where [-5, -2, 4, 5] are from nums1 and [-3, 1, 8] are from nums2

	Example 2:
	Input : nums1 = [0, 2, 7, 8, 0, 0, 0], nums2 = [-7, -3, -1]
	Output :  [-7, -3, -1, 0, 2, 7, 8]
	Explanation :  The merged array is: [-7, -3, -1, 0, 2, 7, 8], where [0, 2, 7, 8] are from nums1 and [-7, -3, -1] are from nums2
*/

func bruteForce(nums1, nums2 []int, n, m int) {
	//  In case we have to use the given 2 arrays itself and arr1 does not have size of n+m
	var arr3 []int = make([]int, n+m)
	var left, right, index int = 0, 0, 0

	for left < n && right < m {
		if nums1[left] <= nums2[right] {
			arr3[index] = nums1[left]
			left++
			index++
		} else {
			arr3[index] = nums2[right]
			right++
			index++
		}
	}
	for left < n {
		arr3[index] = nums1[left]
		left++
		index++
	}
	for right < m {
		arr3[index] = nums2[right]
		right++
		index++
	}
	for i := 0; i < n+m; i++ {
		if i < n {
			nums1[i] = arr3[i]
		} else {
			nums2[i-n] = arr3[i]

		}
	}
	// TC = O(n+m) + O(n+m), n,m = size of the given arrays.
	// SC = O(n+m) additional space used to store the non-overlapping intervals.
}

func optimal1(nums1, nums2 []int, n, m int) {
	//  In case we have to use the given 2 arrays itself and arr1 does not have size of n+m
	var left, right int = n - 1, 0
	for left >= 0 && right < m {
		if nums1[left] > nums2[right] {
			nums1[left], nums2[right] = nums2[right], nums1[left]
			left--
			right++
		} else {
			break
		}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	/// TC = O(min(n+m)) + O(nlogn) + O(mlogm), n,m = size of the given arrays.
	// SC = O(1), as we are not using any extra space to solve this problem.
}

func swapIfGreater(arr1, arr2 []int, idx1, idx2 int) {
	if arr1[idx1] > arr2[idx2] {
		arr1[idx1], arr2[idx2] = arr2[idx2], arr1[idx1]
	}
}
func optimal2(nums1, nums2 []int, n, m int) {
	//  In case we have to use the given 2 arrays itself and arr1 does not have size of n+m with using Gap method from Shell Sort
	var length int = n + m
	gap := (length / 2) + (length % 2)
	for gap > 0 {
		left := 0
		right := left + gap
		for right < length {
			if left < n && right >= n {
				swapIfGreater(nums1, nums2, left, right-n)
			} else if left >= n {
				swapIfGreater(nums2, nums2, left-n, right-n)
			} else {
				swapIfGreater(nums1, nums1, left, right)
			}
			left++
			right++
		}
		if gap == 1 {
			break
		}
		gap = (gap / 2) + (gap % 2)
	}

	/// TC = O(min(n+m)) + O(nlogn) + O(mlogm), n,m = size of the given arrays.
	// SC = O(1), as we are not using any extra space to solve this problem.
}

func optimal3(nums1, nums2 []int, n, m int) {
	//  In case we have to use the given 2 arrays itself and arr1 does not have size of n+m with using Gap method from Shell Sort
	i, j := n-1, m-1
	k := n + m - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}

	/// TC = O(N+M), we traverse both the arrays exactly once.
	// SC = O(1), constant extra space is used to store pointers.
}

func main() {
	nums1 := []int{-5, -2, 4, 5}
	nums2 := []int{-3, 1, 8}
	bruteForce(nums1, nums2, len(nums1), len(nums2))
	fmt.Println("ResultBF: ", nums1, nums2) // result ResultBF:  [-5 -3 -2 1] [4 5 8]
	nums1 = []int{-5, -2, 4, 5}
	nums2 = []int{-3, 1, 8}
	optimal1(nums1, nums2, len(nums1), len(nums2))
	fmt.Println("ResultO1: ", nums1, nums2) // result ResultO1:  [-5 -3 -2 1] [4 5 8]
	nums1 = []int{-5, -2, 4, 5}
	nums2 = []int{-3, 1, 8}
	optimal2(nums1, nums2, len(nums1), len(nums2))
	fmt.Println("ResultO2: ", nums1, nums2) // result ResultO2:  [-5 -3 -2 1] [4 5 8]
	nums1 = []int{-5, -2, 4, 5, 0, 0, 0}
	nums2 = []int{-3, 1, 8}
	optimal3(nums1, nums2, 4, 3)     // Actual
	fmt.Println("ResultO3: ", nums1) // result [-5 -3 -2 1 4 5 8]
}
