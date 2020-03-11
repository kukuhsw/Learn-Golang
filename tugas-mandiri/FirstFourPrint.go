package main

import  "fmt"

var n int
var data int
var tam int 
var first int

func main(){

fmt.Scanln(&data)
 

 n=4
 tam=0

 for data<9999{
 fmt.Scanln(data)
 
 if tam<=n{
 fmt.Println(data) 
 tam=tam+1 
 fmt.Scanln(data)

 
 }
 }
 }
