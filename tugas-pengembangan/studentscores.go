package main
import "fmt"

type StudentType struct{
    name string
	sid  int
	gpa  float64
}

const N=53 

var stuarr[N] StudentType
var countave int

func main(){
    var average float64
    var reqgpa float64
    var deviation float64

    countave = 0
    ReadStudents()
    average = Average(countave)
    deviation = Deviation(countave,average)
    reqgpa = average + (deviation/2)
    WriteStudents(reqgpa)
}

func ReadStudents(){

    var i int

    
    fmt.Scanln(&stuarr[i].sid,&stuarr[i].name,&stuarr[i].gpa)
    i = 1
    for (i < N) && (stuarr[i-1].name != "NOMORE"){
      fmt.Scanln(&stuarr[i].sid,&stuarr[i].name,&stuarr[i].gpa)
      i = i + 1
      countave = countave + 1
    }    
}

func WriteStudents(reqgpa float64){

    var i int
    fmt.Println("These are people who passed the gpa : ")
    i = 0
    for i < N {
      if stuarr[i].gpa > reqgpa{
          fmt.Println("Student #",i+1,stuarr[i].name,",","GPA =",stuarr[i].gpa)
      }
      i = i + 1
    }
}

func Average(num int)float64{

    var divisor float64
    var avestu  float64
    var i       int

    divisor = 0 
    avestu = 0
    i = 0 
    for (i < num){
      avestu = avestu + stuarr[i].gpa
      divisor = divisor + 1
      i = i + 1
    }	
    return (avestu/divisor)
}

func Deviation(num int, ave float64)float64{

    var totdev float64
    var s      float64
    var i      int

    totdev = 0
    i = 0
    for i < num{
      s = stuarr[i].gpa - ave
      if s < 0 {
        s = ave - stuarr[num].gpa
      }
      s = s / (N-1)
      totdev = totdev + s
      i = i + 1
    }
    return totdev
}


