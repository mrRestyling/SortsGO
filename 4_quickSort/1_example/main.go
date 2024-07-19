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
	fmt.Println("Быстрая сортировка")
	fmt.Println("Сложность: O(n * log n) - логарифмическая")
	fmt.Println("Отрезок из первых семи чисел начального массива:", arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	emptyArr = quickSort(emptyArr)
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

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	foot := arr[0] // опорный элемент

	var down, up []int

	for _, num := range arr[1:] {
		if num <= foot {
			down = append(down, num)
		} else {
			up = append(up, num)
		}
	}

	result := append(quickSort(down), foot)   // "ДО" мы помещаем элементы, которые меньше опорного + опорный
	result = append(result, quickSort(up)...) // "ПОСЛЕ" мы помещаем элементы, которые больше опорного

	return result
}
