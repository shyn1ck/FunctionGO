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
	fmt.Println("Ваш начальный слайс: ", slice)
	evenNumbers := filterEvenNumbers(slice)
	fmt.Printf("Четные числа в слайсе: %v\n", evenNumbers)
}

func filterEvenNumbers(slice []int) []int {
	evenNumbers := []int{}
	for _, slice := range slice {
		if slice%2 == 0 {
			evenNumbers = append(evenNumbers, slice)
		}
	}
	return evenNumbers
}
