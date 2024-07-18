package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	// arr := []int{3, 5, 1, 1, 2, 7, 8}
	fmt.Println(arr[:10])
	// fmt.Println(arr)

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	insertionSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Сортировка вставками O(n^2) занимает:", duration) // квадратичная сложность
	fmt.Println(emptyArr[:10])
	// fmt.Println(emptyArr)
}

func createArr() (arr []int) {
	sizeArr := 50000
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key //помещаем элемент справа отменьшего числа
	}
	return arr
}
