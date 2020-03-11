package main
import "fmt"

var name,grade string
var finex,midterm,lab,finepro,raw float32

func main(){
	fmt.Println("Enter your name : ")
	fmt.Scan(&name)
	fmt.Println("Enter your final exam score : ")
	fmt.Scan(&finex)
	fmt.Println("Enter your midterm exam score : ")
	fmt.Scan(&midterm)
	fmt.Println("Enter your lab works score : ")
	fmt.Scan(&lab)
	fmt.Println("Enter your final project score : ")
	fmt.Scan(&finepro)
	
    raw=(finex*.40)+(midterm*.25)+(lab*.15)+(finepro*.20)
	
	if raw>80{
	grade="A"
	}else if raw>70{
	grade="B"
	}else if raw>60{
	grade="C"
	}else if raw>50{
	grade="D"
	}else{
	grade="E"}
	
	fmt.Println(name," your raw score is ",raw," and you receive grade ",grade,"!")
	
	}