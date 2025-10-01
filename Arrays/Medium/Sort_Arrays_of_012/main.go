package main

import (
	"fmt"
)

/*
Given an array consisting of only 0s, 1s, and 2s. Write a program to in-place sort the array without using inbuilt sort functions. ( Expected: Single pass-O(N) and constant space)

	Input: nums = [2,0,2,1,1,0]
	Output: [0,0,1,1,2,2]

	Input: nums = [2,0,1]
	Output: [0,1,2]

	Input: nums = [0]
	Output: [0]
*/

func merge(arr []int, low, mid, high int) {
	var temp []int
	left, right := low, mid+1

	//storing elements in the temporary array in a sorted manner
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	// if elements on the left half are still left
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	//  if elements on the right half are still left
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// transferring all elements from temporary to arr
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}

}

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)    // left half
	mergeSort(arr, mid+1, high) // right half
	merge(arr, low, mid, high)  // merging sorted halves
}

func bruteForce(arr []int) {

	mergeSort(arr, 0, len(arr)-1)
	// TC = O(nlogn)
	// Reason: At each step, we divide the whole array, for that logn and we assume n steps are taken to get sorted array, so overall time complexity will be nlogn
	// SC = O(n)  Reason: We are using a temporary array to store elements in sorted order.
}

func better(arr []int) {
	// Since it has only 3 unique ints
	count0, count1, count2 := 0, 0, 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count0++
		}
		if arr[i] == 1 {
			count1++
		}
		if arr[i] == 2 {
			count2++
		}
	}
	for i := 0; i < count0; i++ {
		arr[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		arr[i] = 1
	}
	for i := count0 + count1; i < len(arr); i++ {
		arr[i] = 2
	}

	// TC = O(N) + O(N), where N = size of the array. First O(N) for counting the number of 0’s, 1’s, 2’s, and second O(N) for placing them correctly in the original array.
	// SC = O(1) as we are not using any extra space.
}

// This problem is a variation of the popular Dutch National flag algorithm.
func optimal(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid < high {
		if arr[mid] == 0 {
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	// TC = O(N), where N = size of the given array.
	// SC = O(1) as we are not using any extra space.
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	arr0 := arr
	bruteForce(arr0)
	fmt.Println("ResultBF: ", arr0) // Any kind of sort other than inbuile with TC = O(N logN)
	arr1 := arr
	better(arr1)
	fmt.Println("ResultB: ", arr1)
	arr2 := arr
	optimal(arr2)
	fmt.Println("ResultO: ", arr2)
}
