package main
import "fmt"

var v,n,m float64
var exp float64
var add float64
var fact,plus float64

func exper(exp float64)float64{

exp=1
n=2
m=0
fact=1
add=1
plus=1
v=2

for m<n{
add=add*v
fact=fact*plus
exp=exp+(add/fact)
plus=plus+1
m=m+1

}
return exp

	}
	
func main(){
fmt.Println(exper(exp))
}
