package main

import "fmt"

var num[100]     float64
var moverage     float64
var period       int
var i,j          int
var next         int
var tes  float64

func main(){
period = 5
tes = 0
j = 0
next = period

for j < next  {

	fmt.Scanln(&num[j])
	if  num[j] < 0{
		num[next] = -1
		j = next + 1
	}
	tes = tes + 1
	j = j + 1  
}

for num[next] > -1 {
	moverage = num[0]
	i = period - 1
	for i >= 1{
		moverage = moverage + num[i]
		i = i - 1
	}
	moverage = moverage / tes
	next = (next + period-1) % period
	fmt.Println("Average : " ,moverage)
	fmt.Scanln(&num[next])
}
}