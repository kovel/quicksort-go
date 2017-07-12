package main

import (
	"fmt"
	"math/rand"
)

func sort3(array []int) {
	quickSort3(array, 0, len(array)-1)
}

func quickSort3(array []int, left, right int) {
	size := right - left + 1
	if size <= 10 {
		insertionSort(array, left, right)
	} else {
		pivot := mediaOf3_2(array, left, right)

		partition := partitionIt3(array, left, right, pivot)
		quickSort3(array, left, partition-1)
		quickSort3(array, partition+1, right)
	}
}
func insertionSort(values []int, left, right int) {
	for i := left + 1; i <= right ; i++ {
		j := i + 1
		temp := values[j]

		for j > left && values[j - 1] >= temp {
			values[j] = values[j - 1]
			j--
		}
		values[j] = temp
	}
}

func mediaOf3_2(values []int, left, right int) int {
	median := (left + right) / 2
	manualSort2(values, left, median, right)
	swap3(values, median, right-1)
	return values[right-1]
}

func partitionIt3(values []int, left, right, pivot int) int {
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
			swap3(values, leftPtr, rightPtr)
		}
	}

	swap3(values, leftPtr, right-1)
	return leftPtr
}

func manualSort2(values []int, left, center, right int) {
	size := right - left + 1
	if size <= 1 {
		return
	} else if size == 2 {
		if values[left] > values[right] {
			swap3(values, left, right)
		}
	} else {
		if values[left] > values[center] {
			swap3(values, left, center)
		}

		if values[left] > values[right] {
			swap3(values, left, right)
		}

		if values[center] > values[right] {
			swap3(values, center, right)
		}
	}
}

func swap3(values []int, first, second int) {
	temp := values[first]
	values[first] = values[second]
	values[second] = temp
}

func printArray3(array []int) {
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

	printArray3(array)
	sort3(array)
	printArray3(array)
}
