package main
import "fmt"

var h1,m1,s1 int
var h2,m2,s2 int
var cost,roundup int


func main(){
	fmt.Scanln(&h1,&m1,&s1)
	fmt.Scanln(&h2,&m2,&s2)
	fmt.Println(Compute(roundup,cost))
	}
	
func Compute(cost,roundup int)(int){
	
	h2=h2-h1
	m2=m2-m1
	s2=s2-s1
	s2=s2/60
	m2=(m2+s2)/60
	h2=h2+m2
	
	if h2>16{
	cost=cost+10000}

	if h2>=2{
	h2=h2-2
	cost=cost+5000}
	
	roundup=h2*3000
	cost=cost+roundup
	return roundup
	return cost
	
	}
	