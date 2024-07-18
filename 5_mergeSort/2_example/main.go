package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// noSort := RandomInt() // срез int
	noSort := []int{3, 1, 5, 6, 2, 4, 7, 8}
	fmt.Printf("Unsorted list:\t%v\n", noSort)
	sorted := sort(noSort)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", sorted)
}

func sort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = sort(left)
	right = sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				fmt.Printf("Sorting:\t%v\n", result)
				left = left[1:]
			} else {
				result = append(result, right[0])
				fmt.Printf("Sorting:\t%v\n", result)
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			fmt.Printf("Sorting:\t%v\n", result)
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			fmt.Printf("Sorting:\t%v\n", result)
			right = right[1:]
		}
	}

	return result
}
func RandomInt() []int {
	list := rand.Perm(20)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}
