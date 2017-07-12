package main

import (
	"fmt"
	"math/rand"
)

func sort2(array []int) {
	quickSort2(array, 0, len(array)-1)
}

func quickSort2(array []int, left, right int) {
	size := right - left + 1
	if size <= 3 {
		manualSort(array, left, right-1, right)
	} else {
		median := (left + right) / 2
		manualSort(array, left, median, right)
		swap2(array, median, right-1)
		pivot := array[right-1]

		partition := partitionIt2(array, left, right, pivot)
		quickSort2(array, left, partition-1)
		quickSort2(array, partition+1, right)
	}
}

func partitionIt2(values []int, left, right, pivot int) int {
	leftPtr := left
	rightPtr := right - 1

	for true {
		for leftPtr++; values[leftPtr] < pivot; leftPtr++ {
		}

		for rightPtr--; values[rightPtr] > pivot && rightPtr > 0; rightPtr-- {
		}

		if leftPtr >= rightPtr {
			break
		} else {
			swap2(values, leftPtr, rightPtr)
		}
	}

	swap2(values, leftPtr, right-1)
	return leftPtr
}

func manualSort(values []int, left, center, right int) {
	size := right - left + 1
	if size <= 1 {
		return
	} else if size == 2 {
		if values[left] > values[right] {
			swap2(values, left, right)
		}
	} else {
		if values[left] > values[center] {
			swap2(values, left, center)
		}

		if values[left] > values[right] {
			swap2(values, left, right)
		}

		if values[center] > values[right] {
			swap2(values, center, right)
		}
	}
}

func swap2(values []int, first, second int) {
	temp := values[first]
	values[first] = values[second]
	values[second] = temp
}

func printArray2(array []int) {
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	const N = 16
	array := make([]int, N)
	for k, _ := range array {
		array[k] = rand.Intn(N * 10)
	}

	printArray2(array)
	sort2(array)
	printArray2(array)
}
