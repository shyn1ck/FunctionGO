package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	fmt.Println("Введите данные для map1 (ключ:значение), введите 'exit' для завершения:")
	for {
		var key string
		var value int
		fmt.Scanln(&key, &value)
		if key == "exit" {
			break
		}
		map1[key] = value
	}

	fmt.Println("Введите данные для map2 (ключ:значение), введите 'exit' для завершения:")
	for {
		var key string
		var value int
		fmt.Scanln(&key, &value)
		if key == "exit" {
			break
		}
		map2[key] = value
	}

	mergedMap := mergeMaps(map1, map2)
	fmt.Println("Объединенная карта:", mergedMap)
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)
	for key, value := range map1 {
		mergedMap[key] = value
	}
	for key, value := range map2 {
		if _, ok := mergedMap[key]; ok {
			mergedMap[key] += value
		} else {
			mergedMap[key] = value
		}
	}
	return mergedMap
}
