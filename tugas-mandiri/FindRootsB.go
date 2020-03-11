package main
import "fmt"

var a,b,c float64
var x1,x2 float64
var root1 float64
var formula1,v1,i,v0 float64 
var z float64
func main(){

fmt.Print("Please input the a,b, and c: ")
	fmt.Scanln(&a,&b,&c)
z=b
formula1=(b*b)-(4*(a*c))
b=signed(b)
root1 = squareroot(v1)

x1=(-z+(b*root1))/(2*z)
x2=c/(a*x1)

fmt.Println(" Your root square is : ",x1,x2)

}

func squareroot(v1 float64)float64{
i=1
v1=100.0
if formula1>0{
for i<=10{
    v0=v1
	v1=v0-(v0*v0-formula1)/(2*v0)
	i=i+1

}
}else{
v1=-1}

return v1
}

func signed(b float64)float64{
   	    if b>0{
		    b=1
		}else if b==0{
		    b=0
		}else if b<0{
		    b=-1}
		return b
	}	