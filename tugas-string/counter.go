package main

import "fmt"

var t       int
var dl,hl   int
var v       byte 
var d[17]   byte
var h[100000]      byte
var counter int
var j       int

func main(){
    fmt.Println("Type the number for searching process.")
    fmt.Scanln(&t)
    for ( t > 0 ){
       
        dl = 0
        fmt.Scanf("%c", &v)
        for (v != '\r') &&(v != '\n'){
	        d[dl] = v
	        fmt.Scanf("%c",&v)
			dl = dl + 1
        }
        hl = 0
		fmt.Scanf("%c",&v)
	    for (v != '\r') && (v != '\n'){
	        hl = hl + 1
	        h[hl] = v
	        fmt.Scanf("%c",&v)
		}
   	    counter = 0
		j = 1
		for j <= hl {
		    if h[j] == d[0] {
			    counter = counter + 1
			}
			j = j + 1
		}
		fmt.Println(counter)
	    t = t - 1
    }
}