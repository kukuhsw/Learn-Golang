// The program is for reading your name.
package main
import "fmt"
var name string
func main () {
	fmt.Println("Hello, Enter your name:")
	fmt.Scanln(&name)
	fmt.Println(name, "Welcome to Algorithm")
}