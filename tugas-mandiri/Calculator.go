package main
import "fmt"

var op1,op2 float64
var i,result float64

func main(){
	
	var calc string
	
	i=100
	for i>1{
		fmt.Print("Enter next expression: ")
		fmt.Scanln(&calc,&op1,&op2)
		if calc == "+"{
			fmt.Println(add())
		}else if calc == "-"{
			fmt.Println(substract())
		}
	}

}



func add()float64{
	result=op1+op2
	return result
}

func substract()float64{
	result=op1-op2
	return result
}

func multiply()float64{
	result=op1*op2
	return result
}

func divide()float64{
	result=op1/op2
	return result
}
