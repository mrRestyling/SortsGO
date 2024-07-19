package main

import (
	"fmt"
	"math/rand"
)

var index, start, end int

func main() {
	arr := RandomInt()
	fmt.Printf("Неотсортированный массив:\t%v\n", arr)
	fmt.Println("")
	sort(arr, 0, len(arr)-1)
	fmt.Println("")
	fmt.Printf("Отсортированный массив:\t%v\n", arr)
}

func sort(s1 []int, start, end int) {
	if start >= end {
		return
	}

	pivot := s1[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > s1[j] {
			s1[i], s1[j] = s1[j], s1[i]
			i++
		}
		fmt.Printf("Сортировка ...:\t%v\n", s1)
	}

	s1[start], s1[i-1] = s1[i-1], s1[start]

	sort(s1, start, i-2)
	sort(s1, i, end)
}

func RandomInt() []int {
	list := rand.Perm(20)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}
