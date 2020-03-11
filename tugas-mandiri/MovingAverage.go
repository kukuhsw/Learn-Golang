package main

import "fmt"

var num[100]     float64
var moverage     float64
var period       float64
var i,j          float64
var next         float64

func main(){
    
    i = 0
    period = 5
    next = period

    for i < next && num[i] >-1 {
        fmt.Scanln(&num[i])
        i = i + 1  
    }

    fmt.Scanln(&num[next])
    fmt.Println("This is saveper :",saveperiod)
    
    for num[next] > -1 {    
        moverage = num[0]
        j = period -1
        for j >= 1 {
            moverage = moverage + num[j]
            j = j - 1
        }
        moverage = moverage/period
        next = (next+period-1) % period
        fmt.Println("This is the average",moverage)
        fmt.Scanln(&num[next])
    }
}