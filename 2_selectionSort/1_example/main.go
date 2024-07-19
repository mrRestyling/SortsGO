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
	fmt.Println("Сортировка выбором/двоичный поиск")
	fmt.Println("Сложность: O(n^2) - квадратичная")
	fmt.Println("Отрезок из первых семи чисел начального массива:", arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	selectionSort(emptyArr)
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

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		// находим индекс минимального элемента
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// помещаем этот минимальный элемент в начало
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
