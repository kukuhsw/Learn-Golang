package main
import "fmt"
type StudentType struct {
	name  string
	sid  int
	gpa  float64
}
const N = 53

var stuarr[N] StudentType

func main(){
	var reqgpa float64
	var numstu int
	var avetot float64
	var totalstu int
	var dev float64

	fmt.Println("Insert the gpa Requirement :")
	fmt.Scan(&reqgpa)
	totalstu = ReadStudents()
	
	fmt.Println("Insert the number of students you want to calculate the ave : ")
	fmt.Scan(&numstu)
	avetot = average(numstu)
	dev = Deviation(numstu,avetot)
	fmt.Println("This is your average :",avetot)
	fmt.Println("This is your Deviation :",dev)
	fmt.Println("Total Students : ",totalstu)
}

func ReadStudents()int{
	var i int
	var totalstudents int
	totalstudents = 0
	i = 0
	fmt.Println("")
		for (stuarr[i].sid != 99999) || (stuarr[i].name != "NOMORE") ||(stuarr[i].gpa != 0.00){
			fmt.Scan(&stuarr[i].sid,&stuarr[i].name,&stuarr[i].gpa)
			i = i + 1
			totalstudents = totalstudents + 1
		}
	return totalstudents
}


func average(num int)float64{
	var save float64
	var ma float64
	save = 0
	for num < -1 {
		ma = ma + stuarr[num].gpa
		num = num - 1
		save = save + 1
	}
	ma = ma /save
	return ma
}

func Deviation(num int, ave float64)float64{
	var absolute float64
	var s		 float64
	var sum float64
	var i int
	var save float64

	sum = 0
	save = 0
	for num < -1{
		absolute = stuarr[i].gpa - ave
		if absolute < 0{
			absolute = ave - stuarr[i].gpa
		}
		save = save + 1
		s = absolute/save
		sum = sum + s 
		num = num - 1
	}
	return sum
}