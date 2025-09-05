package main

import (
	"fmt"
	"html"
	"net/url"
	"strings"
	"unicode/utf8"
)

func main() {

	var s1 string
	fmt.Printf("string, s:%s\n", s1)

	//
	var s2 string = `\xe6`
	fmt.Printf("string, s:%s\n", s2)

	//
	fmt.Println(url.PathEscape("A B"))
	fmt.Println(html.EscapeString("<>"))

	// concetenate
	fmt.Println("In" + "dia")

	//
	country := "india"
	fmt.Printf("india==india: %t\n", "india" == country)
	fmt.Printf("india==INDIA: %t\n", strings.EqualFold("india", "INDIA"))
	fmt.Printf("India<india : %t\n", "India" < "india")

	//
	fmt.Printf("length of string:%d\n", len("日"))
	fmt.Printf("length of string:%d\n", utf8.RuneCountInString("日"))
	fmt.Printf("is Valid utf?%t:\n", utf8.ValidString("日"))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("check anangram of goD and doG:: %t\n", Anagram("goD", "doG"))
	fmt.Println(strings.Repeat("--", 20))

}
