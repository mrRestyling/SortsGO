package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	// arr := []int{3, 5, 1, 6, 2, 7, 8}
	fmt.Println(arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	emptyArr = quickSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Быстрая сортировка O(n * log n) занимает:", duration) // log
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
