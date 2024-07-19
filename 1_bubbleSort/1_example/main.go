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
	fmt.Println("Пузырьковая сортировка")
	fmt.Println("Сложность: O(n^2) - квадратичная")
	fmt.Println("Отрезок из первых семи чисел начального массива:", arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	bubbleSort(emptyArr)
	fmt.Println("Отрезок из первых семи чисел отсортированного массива:", emptyArr[:7])
	duration := time.Since(start)

	fmt.Printf("Время сортировки из %d чисел занимает: %v\n", size, duration) // квадратичная сложность

}

func makeArr(vol int) (arr []int) {
	size := vol
	arr = make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return
}

// Bubble sort (сортировка пузырьком) -
// данный алгоритм меняет местами два соседних элемента, если первый элемент массива больше второго.
// Так происходит до тех пор, пока алгоритм не обменяет местами все неотсортированные элементы.
// Сложность данного алгоритма сортировки равна O(n^2).
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
