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
var found       bool

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
		softkernel()
		softflesh()
		found = isinflesh("it")
		if found {
		    fmt.Println("it is FOUND")
		}else{
			fmt.Println("it is NOT FOUND")
	}
	fmt.Println("Kernel is",i,"and flash is",j)
	file.Close()
    }
}


func softkernel() {
var x, y, max int
var t         string

    x = i
	max = x
	for x < 0{
	    y = x - 1
		for y > 0{
		    if kernel[max] < kernel[y]{
			    max = j
			}
		y = y - 1
		}
		t = kernel[max]
		kernel[max] = kernel[x]
		kernel[x] = t
		x = x - 1
	}
}

func softflesh() {
var x,y   int
var t     string

    x = 0
	for x < j {
		t = flesh[x+1]
		y = j
		for (t > flesh[y+1]) && (y >= 1){
		    kernel[y+1] = kernel[y]
			y = y - 1
		}
		flesh[y+1] = t
		x = x + 1
	}
}


func isinflesh(val string)bool{
var l,r,guess   int
var founds       bool

    founds = false
	for (founds == false) && (l<r){
	    guess = (l+r)/2
		if (flesh[guess] == val){
		    founds = true
		}else if (flesh[guess] < flesh[r]){
		    r = r - 1
		}else if (flesh[guess] < flesh[l]){
		    l = l + 1
		}
	}
	return founds
}