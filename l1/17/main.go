package main

import (
	"fmt"
	"sort"
)

func BinarSearche(nums []int, x int) int {
	// начальные значения для указателей
	l := 0
	r := len(nums) - 1
	// пока указатели не замкнутся
	for l != r {
		// среднее значачения индекс
		mid := (l + r) / 2
		// если мы нашли x, то вернем индекс
		if nums[mid] == x {
			return mid
		}
		// если меньши нашего x, то меняем границу левого, так как там очивидно нет
		if nums[mid] < x {
			l = mid + 1
		}
		// если больше нашего x, то меняем границу правого, так как там очивидно нет
		if nums[mid] > x {
			r = mid - 1
		}
	}
	// крайнии случай
	if nums[l] == x {
		return l
	}
	// если x нет возращаем -1
	return -1
}

func main() {
	a := []int{3, 5, 2, 4}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	fmt.Println(a)

	fmt.Print(BinarSearche(a, 5))
}
