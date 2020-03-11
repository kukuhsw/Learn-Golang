//checking bitwise operation
package main
import "fmt"
var a,b int
func main () {
	fmt.Println("Enter the a and b input")
	fmt.Scanln(&a,&b)
	fmt.Println(^a, -a, a&b, a|b, a^b)
var b uint
	fmt.Println(a<<b, a>>b)
	}
