package main
import "fmt"
var k,m,r int

func main(){
    fmt.Scanln(&k,&m)
    fmt.Println(GCD(k))
	}
	
func GCD(k int)(int){
      for (m != 0){
         r = k % m	
         k = m
         m = r
      }
        return k
	}