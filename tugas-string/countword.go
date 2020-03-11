package main
import "fmt"
 
var t              int
var v              byte 
var j,k            int
var d[17]          byte
var h[100000]      byte
var dl,hl          int
var counter        int
var wordstart      bool

func main(){
    fmt.Scanln(&t)
	j = - 1
    for ( t > 0 ){
       
        dl = 0
        fmt.Scanf("%c", &v)
        for (v != '\r') &&(v != '\n'){
	        d[dl] = v
	        fmt.Scanf("%c",&v)
			dl = dl + 1
        }
        hl = 0
		fmt.Scanf("%c",&v)
	    for (v != '\r') && (v != '\n'){
	        hl = hl + 1
	        h[hl] = v
	        fmt.Scanf("%c",&v)
		}
   	    counter = 0
		wordstart = true
		j = j + 1
		wordstart = h[j] == '\r'
		for (j <= hl) && (wordstart){
		    if j + dl + 1 <= hl{
			    k = 1
				for (k <= dl) && (h[j+k-1] == d[k-1]) {
				    k = k + 1
				}
			    if (k > dl) && (h[j+dl+1] == '\r') {
			        counter = counter + 1
			    }
			}
			j = j + 1
		}
		fmt.Println(counter)
	    t = t - 1
    }
}