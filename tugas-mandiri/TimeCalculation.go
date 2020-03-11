package main

import "fmt"

type TimeType struct {
	hour int
	minute int
	second int
} 

var timer TimeType

func main(){
	var h,m,s int
	var comd  string
	var i int

	i = 1

	for i > 0 {
	fmt.Println("Enter the command")
	fmt.Scan(&comd)
	if comd == "C"{
		fmt.Scan(&h,&m,&s)
		construct(h,m,s)
	}else if comd == "P"{
		stringp()
	}else if comd == "A"{
		fmt.Scan(&s)
		addSeconds(s)
	}else if comd == "S"{
		seconds()
	}
}

}

func seconds(){
	var secondscalc int
	secondscalc = (timer.hour*3600) + (timer.minute*60) + timer.second
	fmt.Println("Total Seconds : ", secondscalc)
}

func addSeconds(sec int)int{
	timer.second = timer.second + sec
	return timer.second
}

func stringp(){
	fmt.Println("This is your stored time : ",timer.hour,"hour,",timer.minute,"minute,",timer.second,"second.")
}

func construct(h,m,s int){
	timer.hour = h
	timer.minute = m
	timer.second = s
}
