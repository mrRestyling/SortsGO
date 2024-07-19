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
	fmt.Println("Сортировка вставками")
	fmt.Println("Сложность: O(n^2) - квадратичная")
	fmt.Println("Отрезок из первых семи чисел начального массива:", arr[:7])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	insertionSort(emptyArr)
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

// Insertion sort (сортировка вставками) -
// На каждой итерации берется элемент и сравнивается с каждым элементом в уже отсортированной части массива,
// таким образом находя «свое место», после чего элемент вставляется на свою позицию.
// Так происходит до тех пор, пока алгоритм не пройдет по всему массиву.
// Сложность данного алгоритма равна O(n^2).
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
