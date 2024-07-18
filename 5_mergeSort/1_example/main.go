package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	// arr := []int{3, 1, 5, 6, 2, 4, 7, 8}
	fmt.Println(arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	// fmt.Println(emptyArr)
	start := time.Now()
	// sort.Ints(emptyArr)
	emptyArr = mergeSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Сортировка слиянием O(n * log n) занимает:", duration) // log
	fmt.Println(emptyArr[:7])

}

func createArr() (arr []int) {
	sizeArr := 50000
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return

}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	// fmt.Println(plus(left, right))

	return plus(left, right)
}

func plus(left, right []int) []int {

	var plused []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			plused = append(plused, left[0])
			left = left[1:]
		} else {
			plused = append(plused, right[0])
			right = right[1:]
		}
	}

	plused = append(plused, left...)
	plused = append(plused, right...)
	return plused
}
