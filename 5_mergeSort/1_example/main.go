package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	size := 100000

	arr := makeArr(size)
	// arr := []int{3, 5, 1, 6, 2, 7, 8}

	fmt.Println()
	fmt.Println("Сортировка слиянием")
	fmt.Println("Сложность: O(n * log n) - логарифмическая")
	fmt.Println("Отрезок из первых семи чисел начального массива:", arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	emptyArr = mergeSort(emptyArr)
	fmt.Println("Отрезок из первых семи чисел отсортированного массива:", emptyArr[:7])
	duration := time.Since(start)

	fmt.Printf("Время сортировки из %d чисел занимает: %v\n", size, duration)

}

func makeArr(vol int) (arr []int) {
	size := vol
	arr = make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
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
