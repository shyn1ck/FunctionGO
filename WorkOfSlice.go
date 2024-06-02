package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите количество элементов в слайсе:")
	fmt.Scan(&n)
	slice := make([]int, n)
	fmt.Println("Введите элементы слайса:")
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	result := WorkOfSlice(slice)
	fmt.Printf("Произведение чисел в слайсе %v равно %d\n", slice, result)
}

func WorkOfSlice(slice []int) int {
	number := 1
	for _, slice := range slice {
		number *= slice
	}
	return number
}
