package main

import (
	f "fmt"
)
var a, b, int

func main () {
	f.Print ("Input Value Desimal : ")
	f.Scanln (&a)
	f.Print ("Value Biner : ")
		for a > 0 {
			b = a % 2
			a = a / 2
		
			f.Print (b)
		}
}