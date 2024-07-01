package main

import ("fmt"
		"github.com/google/go-cmp/cmp")

func main() {
	a := 1
	b := 2
	fmt.Println(cmp.Equal(a,b))
}