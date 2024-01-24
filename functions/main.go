package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := sumUp(1, 10, 20, 30, -10)
	anotherResult := sumUp(1, numbers...)

	fmt.Println(result)
	fmt.Println(anotherResult)
}

func sumUp(startingValue int, numbers ...int) int {
	result := 0

	for _, val := range numbers {
		result += val
	}

	result += startingValue
	return result
}
