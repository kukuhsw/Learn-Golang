package main

import "fmt"

var num[100]     float64
var moverage     float64
var period       int
var i,j          int
var next         int
var tes  float64
var weight float64
var saveper float64

func main(){
period = 5
tes = 1
j = 0
next = period
weight = 1
saveper = 0

for j < next && num[j] > -1  {

    fmt.Scanln(&num[j])
    num[j] = num[j] * weight
    saveper = saveper + 1
    j = j + 1  
}

for num[next] > -1 {
    
    moverage = num[0]
    weight = saveper
    i = period - 1 

    for i >= 1{
        tes = tes + weight
       
        i = i - 1
        moverage = moverage + (num[i] * weight)
        weight = weight - 1
        fmt.Println("I :",i)
        fmt.Println("Weight : ",weight)
        fmt.Println("Test :",tes)
        fmt.Println("Nilai : ",moverage)
    }	
    
	moverage = moverage / tes
	next = (next + period-1) % period
	fmt.Println("Average : " ,moverage)
	fmt.Scanln(&num[next])
}
}