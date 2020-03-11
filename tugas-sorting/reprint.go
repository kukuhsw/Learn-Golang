package main

import "fmt"

var t       int
var dl,hl   int
var v       byte 
var d[17]   byte
var h[10000] byte

func main(){
    fmt.Println("Type the number for searching process.")
    fmt.Scanln(&t)
    for ( t > 0 ){
	    fmt.Println()
        dl = 0
        fmt.Scanf("%c", &v)
        for (v != '\r') &&(v != '\n'){
	        dl = dl + 1
			d[dl] = v
	        fmt.Scanf("%c",&v)
        }
		hl = 0
        fmt.Scanf("%c",&v)
	    for (v != '\r') && (v != '\n'){
	        hl = hl + 1
	        h[hl] = v
	        fmt.Scanf("%c",&v)
			fmt.Printf("%c",h[hl])
	    }
	    t = t - 1
    }
}