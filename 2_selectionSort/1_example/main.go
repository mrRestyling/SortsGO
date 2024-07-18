package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	fmt.Println(arr[:10])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	selectionSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Сортировка выбором/двоичный поиск O(n^2) занимает:", duration) // квадратичная сложность
	fmt.Println(emptyArr[:10])
}

func createArr() (arr []int) {
	sizeArr := 50000
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
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
