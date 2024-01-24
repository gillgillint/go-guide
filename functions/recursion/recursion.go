package recursion

import "fmt"

func main() {

	result := factorial(5)

	fmt.Println(result)

}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
func factorial(number int) int {
	//using loop
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	if number <= 0 {
		return 1
	}
	return number * factorial(number-1)
}
