package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Concat("Hello!", " How are you?"))
	fmt.Println(piscine.Concat("Hello!", " "))
	fmt.Println(piscine.Concat("", ""))
}
