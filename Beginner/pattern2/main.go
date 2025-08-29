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
	PatternWithSpace(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PatternWithLessSpace(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	LowerTrianglePattern(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	HollowInvertedPyramidPattern(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	InvertedHeight(numbs)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	DiamondPattern(numbs)
	fmt.Println(strings.Repeat("--", 20))

}

func PatternWithSpace(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("*")
		for range n - 1 {
			fmt.Printf("*")
		}
		fmt.Printf(" ")
		fmt.Println()
	}
}

func PatternWithLessSpace(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("*")
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		fmt.Println()
	}
}

func LowerTrianglePattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (n - i) < j {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ") // strconv.Itoa(j)
			}
		}
		fmt.Println()
	}
}

func HollowInvertedPyramidPattern(n int) {
	// Loop N times
	for i := 1; i <= n; i++ {

		// print i stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		// print 2*(n-i) spaces
		for j := 1; j <= 2*(n-i); j++ {
			fmt.Print(" ")
		}

		// print i stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println("")

	}
}

func InvertedHeight(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= (i - 1); j++ {
			print(" ")
		}
		for k := 1; k <= n-i+1; k++ {
			print("*")
		}
		print("\n")
	}
}

func DiamondPattern(n int) {
	// First 3 line
	for i := 1; i <= n; i++ {
		// Empty box
		for j := 1; j <= n; j++ {
			print(".")
		}
		// Triangle Space
		for k := 1; k <= n-i; k++ {
			print(".")
		}
		// Triangle Star
		for k := 1; k <= (2*(i-1) + 1); k++ {
			print("*")
		}
		// Triangle Space
		for l := 1; l <= n-i; l++ {
			print(".")
		}
		// Empty box
		for m := 1; m <= n; m++ {
			print(".")
		}
		print("\n")
	}

	// Second 3 lines
	for i := 1; i <= n; i++ {

		// Side triangle Space
		for j := 1; j <= (i - 1); j++ {
			print(".")
		}
		// Side triangle Star
		for k := 1; k <= (n - i + 1); k++ {
			print("*")
		}
		// Star Box
		for l := 1; l <= (2*n - 1); l++ {
			print("*")
		}
		// Side triangle star
		for m := 1; m <= (n - i + 1); m++ {
			print("*")
		}
		// Side triangle Space
		for j := 1; j <= (i - 1); j++ {
			print(".")
		}
		print("\n")
	}

	// Next 3 lines
	for i := 1; i <= n; i++ {
		// Side triangle Space
		for j := 1; j <= (n - i); j++ {
			print(".")
		}
		// Side triangle Star
		for k := 1; k <= i; k++ {
			print("*")
		}
		// Star box
		for l := 1; l <= (2*n - 1); l++ {
			print("*")
		}
		// Side triangle Star
		for m := 1; m <= i; m++ {
			print("*")
		}
		// Side triangle space
		for o := 1; o <= (n - i); o++ {
			print(".")
		}
		print("\n")
	}
	//  Next 3 lines
	for i := 1; i <= n; i++ {
		// Empty Box
		for j := 1; j <= n; j++ {
			print(".")
		}
		// Triangle Space
		for k := 1; k <= (i - 1); k++ {
			print(".")
		}
		// Triangle star
		for k := 1; k <= (2*(n-i) + 1); k++ {
			print("*")
		}
		// Triangle star
		for k := 1; k <= (i - 1); k++ {
			print(".")
		}
		// Triangle Empty box
		for m := 1; m <= n; m++ {
			print(".")
		}
		print("\n")
	}
	print("\n")

}
