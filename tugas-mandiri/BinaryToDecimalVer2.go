	package main
	import "fmt"

	var num int
	var pow int
	var pang int

	func main(){

	fmt.Scanln(&num)

	pow=1
	pang=0

	for pow<=num{
	pow=pow*2
	pang=pang+1
	num=num/2}

	fmt.Println(pang)

	}