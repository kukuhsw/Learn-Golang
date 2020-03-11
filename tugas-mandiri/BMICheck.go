package main

import "fmt"

var bb,tb,bmi,x float32

func main () {

	
	fmt.Print("masukkan berat badan : ")
	fmt.Scanln(&bb)
	fmt.Print("masukkan tinggi badan : ")
	fmt.Scanln(&tb)
	bmi = bb/(x*x)
	fmt.Print("BMI : ")
	fmt.Println(bmi)
	if bmi < 18.5 {
  		fmt.Print("Too low")
 	}else if bmi < 24.9{
  		fmt.Print("Normal")
 	}else if bmi < 29.9{
		fmt.Print("Overweight")
 	}else{
  		fmt.Print("Obesity")
 	}
 }