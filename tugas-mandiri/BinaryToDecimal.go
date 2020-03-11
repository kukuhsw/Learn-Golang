package main
import "fmt"

var binary,digit,hit float64

func main(){
fmt.Scanln(&binary)
digit=binary
hit=1

for digit>1.2{
digit=digit/10
hit=hit+1}



fmt.Println(hit)

}