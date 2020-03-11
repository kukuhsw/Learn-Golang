package main
import "fmt"

var a,b,c,d,p int

func main(){
    fmt.Scanln(&a,&b,&c,&d)
	
	p=a
	a=b
	b=c
	c=d
	d=p
 	
	fmt.Print(a,b,c,d)
}
	