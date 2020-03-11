package main
import "fmt"

type StudentType struct{
    name string
	sid  int
	gpa  float64
}

const N=53 

var stuarr[N] StudentType
var sortGPA[N] int
var sortName[N] int
var sortSid[N] int

func main(){
    var commandstu string
    var sortcmd    string
    var searchstu  string
    var searchgpa  float64
    var searchdex  int
    var shiftname  string
    var shiftsid    int
    var shiftgpa   float64
    var i          int

    fmt.Println("Type the command you want to do : ")
    fmt.Scanln(&commandstu)
    for commandstu != "exit"{
      if commandstu == "data"{
        ReadMoreData()
      }
      if commandstu == "write"{
        WriteData()
      }
      if commandstu == "sort"{
        fmt.Println("Type what you want to sort : ")      
        fmt.Scanln(&sortcmd)
          if sortcmd == "gpa"{
            SortGPA()
            i = 0
            for i < N{
              shiftgpa = stuarr[i].gpa
              stuarr[i].gpa = stuarr[sortGPA[i]].gpa
              stuarr[sortGPA[i]].gpa = shiftgpa
              i = i + 1
            }
          }
          if sortcmd == "name"{
            SortName()
            i = 0
            for i < N{
              shiftname = stuarr[i].name
              stuarr[i].name = stuarr[sortName[i]].name
              stuarr[sortName[i]].name = shiftname
              i = i + 1
            }
          }
          if sortcmd == "sid"{
            SortId()
            i = 0
            for i < N{
              shiftsid = stuarr[i].sid
              stuarr[i].sid = stuarr[sortSid[i]].sid
              stuarr[sortSid[i]].sid = shiftsid
              i = i + 1
            }
          }
          if (sortcmd != "gpa") || (sortcmd != "name") || (sortcmd != "sid"){
            fmt.Println("Cannot sort. Invalid command.")
          }
      }
      if commandstu == "search"{
        fmt.Println("Type the student name you want to search")
        fmt.Scan(&searchstu)
        searchdex = SearchStudent(searchstu)
        if searchdex == -1 {
          fmt.Println("There is no student in the list.")
        }else{
          fmt.Println("We found something in the list ! ")
          WriteInfo(searchdex)
        }
      }
      if commandstu == "find"{
        fmt.Println("Type the student gpa you want to search")
        fmt.Scan(&searchgpa)
        searchdex = FindGPA(searchgpa)
        if searchdex == -1 {
          fmt.Println("There is no student in the list.")
        }else{
          fmt.Println("We found something in the list ! ")
          WriteInfo(searchdex)
        }
      }
      fmt.Println("Type the command you want to do : ")
      fmt.Scanln(&commandstu)
      }   
}


func ReadMoreData(){
    var  i int
    
    fmt.Scanln(&stuarr[i].sid,&stuarr[i].name,&stuarr[i].gpa)
    i = 1
    for (i < N) && (stuarr[i-1].name != "NOMORE"){
      fmt.Scanln(&stuarr[i].sid,&stuarr[i].name,&stuarr[i].gpa)
      i = i + 1
    }    
}

func WriteData(){
    var i int

    i = 0
    for i < N {
      fmt.Println("#",stuarr[i].sid,stuarr[i].name,",","GPA =",stuarr[i].gpa)
      i = i + 1
    }
}

func WriteInfo(index int){

    if (index < 0) || (index>N){
      fmt.Println("Student record not found.")
    }else{
      fmt.Println("Student Identification : ", stuarr[index].sid)
      fmt.Println("Name : ", stuarr[index].name)
      fmt.Println("GPA : ", stuarr[index].gpa)
    }
}

func SearchStudent(name string)int{
    var index int
    var found bool
    var i     int

    index = -1
    found = false
    i = 0
    for (i<N) && (found == false){
      if stuarr[i].name == name{
        found = true
        index = i
      }else{
        i = i + 1    
      }
    }
    return index	
}


func FindGPA(gpa float64)int{
    var index int
    var found bool
    var i     int

    index = -1
    found = false
    i = 0
    for (i<N) && (found == false){
      if stuarr[i].gpa >= gpa {
        found = true
        index = i
      }else{
        i = i + 1    
      }
    }
    return index
}

func SortId(){
    var i,j,max int
    var shiftid int

    i = 0
    for i < N{
      sortSid[i] = i
      i = i + 1
    }
    i = N - 1
    for i > 0{
      max = 0
      j = 2
      for j <= i{
        if stuarr[sortSid[j]].sid > stuarr[sortSid[max]].sid{
          max = j
        }
        j = j + 1
      }
      shiftid = sortSid[max]
      sortSid[max] = sortSid[j]
      sortSid[j] = shiftid
      i = i - 1
      }
}

func SortName(){
    var i,j,max int
    var shiftname int

    i = 0
    for i < N{
      sortName[i] = i
      i = i + 1
    }
    i = N - 1
    for i > 0{
      max = 0
      j = 2
      for j <= i{
        if stuarr[sortName[j]].sid > stuarr[sortName[max]].sid{
          max = j
        }
        j = j + 1
      }
      shiftname = sortName[max]
      sortName[max] = sortName[j]
      sortName[j] = shiftname
      i = i - 1
      }
}

func SortGPA(){
    var i,j,max int
    var shiftGPA int

    i = 0
    for i < N{
      sortGPA[i] = i
      i = i + 1
    }
    i = N - 1
    for i > 0{
      max = 0
      j = 2
      for j <= i{
        if stuarr[sortGPA[j]].sid > stuarr[sortGPA[max]].sid{
          max = j
        }
        j = j + 1
      }
      shiftGPA = sortGPA[max]
      sortGPA[max] = sortGPA[j]
      sortGPA[j] = shiftGPA
      i = i - 1
      }
}

