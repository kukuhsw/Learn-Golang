// This program is for checking the boolean value
package main
import "fmt"
var p,q bool
func main () {
	fmt.Scanln(&p,&q)
	fmt.Println(p&&q, p||q, !p, p==q)
}
