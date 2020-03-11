package main
import "fmt"



var a,b,c int
var x1,x2 int
var root1 int
var formula1,v1,i,v0 int 
func main(){

fmt.Print("Please input the a,b, and c: ")
fmt.Scanln(&a,&b,&c)
formula1=(b*b)-(4*(a*c))
root1 = squareroot(v1)

x1=(-b+root1)/(2*a)
x2=(-b-root1)/(2*a)

fmt.Println(" Your root square is : ",x1,x2)

}

func squareroot(v1 int)int{
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

//Kamu nakal ih gabisa jalan error mulu. Kzl. 