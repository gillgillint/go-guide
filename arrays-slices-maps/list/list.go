package list

import "fmt"

// practice
// Time to practice what you learned!

//  1. Create a new array (!) that contains three hobbies you have
//     Output (print) that array in the command line.
//  2. Also output more data about that array:
//     - The first element (standalone)
//     - The second and third element combined as a new list
//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.

type Product struct {
	id    string
	title string
	price float64
}

// func main() {
// 	// 1)
// 	hobbies := [3]string{"Sports", "Cookies", "Reading"}
// 	fmt.Println(hobbies)

// 	// 2)
// 	fmt.Println(hobbies[0])
// 	fmt.Println(hobbies[1:])

// 	// 3)
// 	mainHobbies := hobbies[:2]
// 	fmt.Println(mainHobbies)

// 	// 4)
// 	fmt.Println(cap(mainHobbies))
// 	mainHobbies = mainHobbies[1:3]
// 	fmt.Println(mainHobbies)

// 	// 5)
// 	courseGoals := []string{"Learn Go", "Learn basic"}
// 	fmt.Println(courseGoals)

// 	// 6)
// 	courseGoals[1] = "Learn Details"
// 	courseGoals = append(courseGoals, "Learn basic")
// 	fmt.Println(courseGoals)

// 	// 7)
// 	products := []Product{{"1", "-1", 3.33}, {"2", "product-2", 12.99}}
// 	fmt.Println(products)

// 	newProduct := Product{"3", "product-3", 10.49}
// 	products = append(products, newProduct)
// 	fmt.Println(products)

// }

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	discountPrices := []float64{2.99, 1.59, 3.99}
	updatedPrices = append(updatedPrices, discountPrices...)
	fmt.Println(updatedPrices)
}

// func main() {
// 	var productNames [4]string
// 	productNames = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[3])

// 	featuredPrices := prices[1:]

// 	highlightFeatured := featuredPrices[:1]
// 	fmt.Println(highlightFeatured)
// 	fmt.Println(featuredPrices)

// 	fmt.Println(cap(prices))
// }
