package main

import "fmt"

func main () {
	var (a = 1
		 b = 1)
	for a <= 1000 {
	fmt.Print(a," ")	
	fmt.Print(b," ")
	a = a+b
	b = b+a
}

}