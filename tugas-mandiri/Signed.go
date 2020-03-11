//Program signedroot.go

package main
import "fmt"

//input Var
var a float64

//Squareroot Var
var x,v0,v1 float64
var i int

func main(){
fmt.Print("Input number for sign and squareroot : ")
fmt.Scanln(&a)
fmt.Print("This is sign of the number : ")
fmt.Println(signed(a))
fmt.Print("This is the squareroot of the number : ")
if a>-1{
fmt.Println(squareroot(v1)) 
}else{
fmt.Println("Your input is a negative number. Cannot Process.")}

}


func squareroot(v1 float64)float64{
i=1
v1=100.0
if a>0{
for i<=10{
    v0=v1
	v1=v0-(v0*v0-a)/(2*v0)
	i=i+1

}
}else{
v1=-1}

return v1
}

func signed(a float64)float64{
   	    if a>0{
		a=1
		}else if a==0{
		a=0
		}else if a<0{
		a=-1}
		return a
	}	
