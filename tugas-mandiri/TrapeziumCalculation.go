//The program is for measuring the area of trapesium.
package main
import "fmt"
var a,b,h int
func main () {
	fmt.Println("Enter three real numbers (top line, base line, and height)")
	fmt.Scanln(&a,&b,&h)
	fmt.Println("The area of trapesium is",(a+b)*h/2)
}