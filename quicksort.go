package main

import (
	"math/rand"
	"fmt"
)

func sort(values []int) {
	quickSort(values, 0, len(values) - 1)
}

func quickSort(values []int, left, right int) {
	if (right - left <= 0) {
		return
	} else {
		pivot := values[right]

		partition := partitionIt(values, left, right, pivot)
		quickSort(values, left, partition - 1)
		quickSort(values, partition + 1, right)
	}
}

func partitionIt(values []int, left int, right, pivot int) int {
	leftPtr := left - 1
	rightPtr := right

	for true {
		leftPtr++
		for ; values[leftPtr] < pivot; leftPtr++ {}

		rightPtr--
		for ; rightPtr > 0 && values[rightPtr] > pivot; rightPtr-- {}

		if leftPtr >= rightPtr {
			break
		} else {
			swap(values, leftPtr, rightPtr)
		}
	}
	swap(values, leftPtr, right)
	return leftPtr
}
func swap(values []int, i int, j int) {
	temp := values[i]
	values[i] = values[j]
	values[j] = temp
}

func printArray(values []int) {
	for _, v := range values {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	const N int = 16;

	values := make([]int, N)
	for k, _ := range values {
		values[k] = rand.Intn(N * 10)
	}

	printArray(values)
	sort(values)
	printArray(values)
}