package main

import "fmt"

/*

You are given an array of integers, your task is to move all the zeros in the array to the end of the array and move non-negative integers to the front by maintaining their order.
	Example 1:
	Input: 1 ,0 ,2 ,3 ,0 ,4 ,0 ,1
	Output: 1 ,2 ,3 ,4 ,1 ,0 ,0 ,0
	Explanation: All the zeros are moved to the end and non-negative integers are moved to front by maintaining order

	Example 2:
	Input: 1,2,0,1,0,4,0
	Output: 1,2,1,4,0,0,0
	Explanation: All the zeros are moved to the end and non-negative integers are moved to front by maintaining order
*/

func bruteForce(arr []int) {
	var temp []int
	for _, val := range arr {
		if val != 0 {
			temp = append(temp, val)
		}
	}
	nz := len(temp)

	for i := 0; i < nz; i++ {
		arr[i] = temp[i]
	}
	for i := nz; i < len(arr); i++ {
		arr[i] = 0
	}
	// TC =  O(N) + O(X) + O(N-X) ~ O(2*N), where N = total no. of elements,
	// X = no. of non-zero elements, and N-X = total no. of zeros.
	// Reason: O(N) for copying non-zero elements from the original to the temporary array.
	// O(X) for again copying it back from the temporary to the original array.
	// O(N-X) for filling zeros in the original array. So, the total time complexity will be O(2*N).

	// SC = O(N), as we are using a temporary array to solve this problem and the maximum size of the array can be N in the worst case.
}

func optimal(arr []int) {
	j := -1
	for idx, val := range arr {
		if val == 0 {
			j = idx
			break
		}
	}
	for i := j + 1; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}

	// TC = O(N), N = size of the array. Reason: We have used 2 loops and using those loops, we are basically traversing the array once.
	// SC = O(1) as we are not using any extra space to solve this problem.
}

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 0, 1}
	arr2 := []int{1, 2, 0, 1, 0, 4, 0}
	bruteForce(arr2)
	fmt.Println("Result: ", arr2)
	optimal(arr1)
	fmt.Println("Result: ", arr1)
}
