package main

import (
	"fmt"
	"sort"
)

/*
Given two sorted arrays, arr1, and arr2 of size n and m. Find the union of two sorted arrays.
The union of two arrays can be defined as the common and distinct elements in the two arrays.NOTE: Elements in the union should be in ascending order.

	Example 1:
	Input:n = 5,m = 5.
	arr1[] = {1,2,3,4,5}
	arr2[] = {2,3,4,4,5}
	Output:
	{1,2,3,4,5}

	Explanation:
	Common Elements in arr1 and arr2  are:  2,3,4,5
	Distnict Elements in arr1 are : 1
	Distnict Elemennts in arr2 are : No distinct elements.
	Union of arr1 and arr2 is {1,2,3,4,5}

	Example 2:
	Input:n = 10,m = 7.
	arr1[] = {1,2,3,4,5,6,7,8,9,10}
	arr2[] = {2,3,4,4,5,11,12}
	Output: {1,2,3,4,5,6,7,8,9,10,11,12}
	Explanation:
	Common Elements in arr1 and arr2  are:  2,3,4,5
	Distnict Elements in arr1 are : 1,6,7,8,9,10
	Distnict Elemennts in arr2 are : 11,12
	Union of arr1 and arr2 is {1,2,3,4,5,6,7,8,9,10,11,12}

*/

func bruteForce(arr1, arr2 []int) []int {
	set := make(map[int]struct{})

	// Insert all elements from arr1
	for _, v := range arr1 {
		set[v] = struct{}{}
	}
	// Insert all elements from arr2
	for _, v := range arr2 {
		set[v] = struct{}{}
	}
	// Collect keys into a slice
	result := make([]int, 0, len(set))
	for k := range set {
		result = append(result, k)
	}
	// Sort the result (since map order is random)
	sort.Ints(result)
	return result
	// TC = O((n+m) log(n+m)), Insert into map → O(n + m) average case Extract + sort → O((n+m) log(n+m)).
	// Inserting a key in map takes logN times, where N is no of elements in map.
	// At max map can store m+n elements {when there are no common elements and elements in arr,arr2 are distntict}.
	// So Inserting m+n th element takes log(m+n) time. Upon approximation across insertion of all elements in worst it would take O((m+n)log(m+n) time.
	// Using HashMap also takes the same time, On average insertion in unordered_map takes O(1) time but sorting the union vector takes O((m+n)log(m+n))  time.
	//  Because at max union vector can have m+n elements.

	// SC = O(n+m) for the map.

}

func optimal(arr1, arr2 []int) []int {

	i, j := 0, 0
	var union []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}
	for i < len(arr1) {
		if union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}
	for j < len(arr2) {
		if union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}

	return union
	// TC = O(m+n), Because at max i runs for n times and j runs for m times. When there are no common elements in arr1 and arr2 and all elements in arr1, arr2 are distinct.
	// SC = O(m+n) {If Space of Union ArrayList is considered}
}
func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 3, 4, 4, 5, 11, 12}
	fmt.Println("ResultBF: ", bruteForce(arr1, arr2))

	fmt.Println("ResultO: ", optimal(arr1, arr2))
}
