package main

import "fmt"

// Перебираем ключи у первого множества если этот ключ есть во втором множестве то заносем
func IntersectionofSets(set1 map[int]interface{}, set2 map[int]interface{}) (IntersectedSets []int) {
	for key := range set1 {
		if _, ok := set2[key]; ok {
			IntersectedSets = append(IntersectedSets, key)
		}
	}

	return
}

func main() {
	set1 := make(map[int]interface{})
	set2 := make(map[int]interface{})

	set1[5] = nil
	set1[2] = nil
	set1[3] = nil
	set1[5] = nil
	set1[11] = nil

	set2[10] = nil
	set2[11] = nil
	set2[2] = nil
	set2[5] = nil

	fmt.Print(IntersectionofSets(set1, set2))
}
