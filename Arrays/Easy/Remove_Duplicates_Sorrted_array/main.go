package main

import "fmt"

/*
Given an integer array sorted in non-decreasing order, remove the duplicates in place such that each unique element appears only once. The relative order of the elements should be kept the same.
If there are k elements after removing the duplicates, then the first k elements of the array should hold the final result. It does not matter what you leave beyond the first k elements.
Note: Return k after placing the final result in the first k slots of the array.

Example 1:
Input: arr[1,1,2,2,2,3,3]
Output: arr[1,2,3,_,_,_,_]
Explanation: Total number of unique elements are 3, i.e[1,2,3] and Therefore return 3 after assigning [1,2,3] in the beginning of the array.

Example 2:
Input: arr[1,1,1,2,2,3,3,3,3,4,4]
Output: arr[1,2,3,4,_,_,_,_,_,_,_]
Explanation: Total number of unique elements are 4, i.e[1,2,3,4] and Therefore return 4 after assigning [1,2,3,4] in the beginning of the array.
*/

func bruteForce(arr []int) []int {
	set := make(map[int]struct{})
	for _, val := range arr {
		set[val] = struct{}{}
	}
	idx := 0
	for key, _ := range set {
		arr[idx] = key
		idx++
	}
	return arr
	// TC =  O(n*log(n)) + O(n)  for iterating on map and normal iteration
	// SC = O(n)   map might have same size as arr
}

func optimal(arr []int) []int {
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return arr
	// TC =  O(n)  for iterating
	// SC = O(1)
}

func main() {
	// arr1 := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}
	arr2 := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}
	// fmt.Println("Result: ", bruteForce(arr1))
	//  In golang the Brute force doesnt work as golang doesnt have set  to do the brute force and map is unordered so gives random answer

	fmt.Println("Result: ", optimal(arr2))

}
