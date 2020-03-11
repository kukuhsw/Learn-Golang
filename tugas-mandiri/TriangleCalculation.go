	package main

	import "fmt"

	var a,b,c,add,add1,side1,side2,long float32

	func main(){

		fmt.Scanln(&a,&b,&c)
		
		
		if (a>b) && (a>c){
		  long=a;side1=b;side2=c;
		  }else if (b>a) && (b>c){
		  long=b;side1=a;side2=c;
		  }else if (c>a) && (c>b){
		  long=c;side1=a;side2=b;}

			add=(side1*side1)+(side2*side2)
			add1=long*long
		
		if add1>add{
		  fmt.Print("Obtuse")		
		}else if add1<add{
		  fmt.Print("Acute")
		
		  }else if add==add1{
		 fmt.Print("Ngentot")
		 }}