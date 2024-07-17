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
	bubbleSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Пузырьковая сортировка O(n^2) занимает:", duration) // квадратичная сложность
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
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
