/* ***********************************************************************
 *
 *                           Bubble Sort
 *
 *  Except for the classic 'hello world!' application, this is the first
 *  application I have written in Go.
 *
 *  Bubble sort is not useful for production purposes, but it is a useful
 *  example in the classroom for teaching sorting algorithms.
 *
 *  Bubble sort makes multiple passes through the array, bubbling the
 *  next highest value to the next top-most position by continually
 *  comparing neigboring values in the array.
 *
 *  Runs in O(n^2).
 *
 * ******************************************************************** */

package main

import "fmt"

func bubbleSort(array []int, arrayLength int) {
	j := arrayLength - 1

	for j > 0 {
		for i := 0; i < j; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		j--
	}
}

func main() {
	var array = []int{7, 2, 1, 0, 10, 8, 6, 5, 3, 4, 9}
	arrayLength := 11

	bubbleSort(array, arrayLength)

	for i := 0; i < arrayLength; i++ {
		fmt.Printf("% d", array[i])
	}
}
