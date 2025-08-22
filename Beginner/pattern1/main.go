package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 1; i <= 10; i += 2 {
		fmt.Printf("i: %d\n", i)
	}

	var numbs int
	fmt.Println("Enter the number of Stars:\t")
	fmt.Scan(&numbs)

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Println(PrintNStart(numbs))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Println(PrintSqrStar(numbs))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PrintHeightN(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PrintInvertedHeightN(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	YetAnotherPrintInvertedHeightN(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PatternEven(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PatternOdd(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	var (
		m int
		n int
	)
	fmt.Println("Enter the number of M::\t")
	fmt.Scan(&m)
	fmt.Println("Enter the number of M::\t")
	fmt.Scan(&n)
	fmt.Println(PrintMxNRectangle(m, n))
	fmt.Println(strings.Repeat("--", 20))

}

func PrintHeightN(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func PrintInvertedHeightN(n int) {
	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func YetAnotherPrintInvertedHeightN(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= (n - i + 1); j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func PrintMxNRectangle(m, n int) string {
	pattern := ""
	for row := 0; row < m; row++ {
		rowwise := ""
		for col := 0; col < n; col++ {
			rowwise += "*" //strconv.Itoa(col)
		}
		pattern += rowwise + "\n"
	}
	return pattern
}

func PrintSqrStar(n int) string {
	pattern := ""
	for row := 0; row < n; row++ {
		rowwise := ""
		for col := 0; col < n; col++ {
			rowwise += "*" //strconv.Itoa(col)
		}
		pattern += rowwise + "\n"
	}
	return pattern
}

func PrintNStart(n int) string {
	star := ""
	// for range n {
	// 	star += "*"
	// }
	for i := 0; i < n; i++ {
		star += "*"
	}
	return star
}

func PatternEven(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if j%2 == 0 {
				fmt.Printf("%d", j)
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}

func PatternOdd(n int) {
	for i := 1; i <= n; i++ {
		// init for every row
		c := 1
		for j := 1; j <= i; j++ {
			if j%2 != 0 {
				fmt.Printf("%d", c)
				c++
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}
