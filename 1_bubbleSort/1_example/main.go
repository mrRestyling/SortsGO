package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	// arr := []int{1, 5, 3, 6, 2, 7, 8}
	fmt.Println(arr[:10])
	// fmt.Println(arr)

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	bubbleSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Пузырьковая сортировка O(n^2) занимает:", duration) // квадратичная сложность
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

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ { // перебираем все пары, меняем местами
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
