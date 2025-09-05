package main

import "fmt"

func CelusiusToFahrenheit(c int) string {
	f := ((9 / 5) * c) + 32
	return fmt.Sprintf("%.2d", f)
}
