package main
import "fmt"
import "io"
import "os"

var kernel[100] string
var flesh[1012] string
var err         error
var file        *os.File
var n, i, j        int
var fname       string

func main(){
    fname = "mod9dat1.dat"
	file, err = os.Open(fname)
	if err == nil{
	    i = 0
		j = 0
		fmt.Fscan(file,&n)
        for n > 0{
		    fmt.Fscan(file,&kernel[i])
			i = i + 1
			n = n - 1
		}
		_,err = fmt.Fscan(file,&flesh[j])
		for err != io.EOF{
		    _,err = fmt.Fscan(file,&flesh[j])
			j = j + 1
		}
	}
	fmt.Println("Kernel is",i,"and flash is",j)
	file.Close()
}