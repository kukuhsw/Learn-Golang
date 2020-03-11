package main
import "fmt"

var a,b,savemin,savemax int

func main(){

fmt.Scanln(&a)

savemin=a
savemax=a

	   for a > -9999{	   
	   if savemin>a{
	   savemin=a
	   }else if savemax<a{
	   savemax=a}
	   fmt.Scanln(&a)
	   }
	   
fmt.Print(savemin,savemax)

}