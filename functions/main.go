package main

import (
	"fmt"
	"strings"
)

func main() {

	var numbs int
	fmt.Print("Enter the number of Stars: N=")
	fmt.Scan(&numbs)

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Factorial of %d is %d\n", numbs, Factorial(numbs))
	fmt.Println(strings.Repeat("--", 20))

	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Factorial of %d is %s\n", numbs, isEvenOdd(numbs))
	fmt.Println(strings.Repeat("--", 20))

	var celesius int
	fmt.Print("Enter the value of celesius: c=")
	fmt.Scan(&celesius)

	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Fahrenheit of %d is %s\n", numbs, CelusiusToFahrenheit(celesius))
	fmt.Println(strings.Repeat("--", 20))

}

func isEvenOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
