package main

import (
	"fmt"
	"strings"
)

func main() {

	a1 := [10]int{}
	fmt.Printf("array:: %#v and type:%T; len:%d amd cap:%d\n", a1, a1, len(a1), cap(a1))

	a2 := make([]int, 2, 10) // slice
	fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", a2, a2, len(a2), cap(a2))

	a3 := make([]int, 10)
	fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", a3, a3, len(a3), cap(a3))

	var a4 []int
	// a4[2] = 10  `panic`  a nil slice
	fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", a4, a4, len(a4), cap(a4))

	// var n int
	// fmt.Print("Enter the size of Array:")
	// fmt.Scanln(&n)

	// a5 := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanln(&a5[i])
	// }
	// fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", a5, a5, len(a5), cap(a5))

	// Slicing
	s1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", s1, s1, len(s1), cap(s1))

	fmt.Println(s1[1:3])
	fmt.Println(s1[3:])
	fmt.Println(s1[:])

	// Iteration
	s := []string{"Foo", "Bar"}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Append
	a := []int{1, 2}
	a = append(a, 3, 4) // a==[1 2 3 4]

	b := []byte("ba")
	b1 := append(b, 'd')
	b2 := append(b, 'g')
	fmt.Printf("string:%s and len:%d and cap:%d\n", string(b1), len(b1), cap(b1))
	fmt.Printf("string:%s and len:%d and cap:%d\n", string(b2), len(b1), cap(b1))

	x1 := []int{1, 2, 3, 5, 6, 8, 32, 79}
	x2 := []int{11, 22, 11, 22}
	x1 = append(x1, x2...)
	fmt.Printf("array:: %#v and type:%T; len:%d and cap:%d\n", x1, x1, len(x1), cap(x1))

	var c1 = make([]int, 3)
	_ = copy(c1, []int{0, 1, 2, 3})

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Print("Biggest Element in the Array,x1: ")
	fmt.Println(getMaxMinInArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("check Array in Inc Order [1,2,34,56,7,8]: %t\n", CheckPairIncOrder([]int{1, 2, 34, 56, 7, 8}))
	fmt.Printf("check Array in Inc Order [1,2,3,4,5,6,7,8]: %t\n", CheckPairIncOrder([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("frequeny if array x1:%v = %v\n", x1, FrequencyOfArray(x1))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	// var mm int
	// fmt.Scan(&mm)
	fmt.Printf("check %d in array x1:%v = %t\n", 32, x1, checkKInArray(x1, 32))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("count %d in array x1:%v = %d\n", 55, x1, CountElementInArray(x1, 55))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("check is array sorted: %t\n", iSArraySorted([]int{1, 2, 3, 4, 8, 6, 7}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("pair count [1,1,1]: %d\n", PairCount([]int{1, 1, 1}, 2))
	fmt.Printf("pair count [1, 2,3,2,1]: %d\n", PairCount([]int{1, 2, 3, 2, 1}, 5))
	fmt.Println(strings.Repeat("--", 20))

}
