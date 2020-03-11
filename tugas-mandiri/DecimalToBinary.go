package main
import "fmt"

var place,binary,decimal,rem,tempDecimal int

func main(){

    place = 1  
    binary = 0

    /* 
     * Reads decimal number from user 
     */  
	 
    fmt.Print("Enter any decimal number: ")  
    fmt.Scanln(&decimal)  
    tempDecimal = decimal  

    /* 
     * Converts the decimal number to binary number 
     */  
	 
    for (tempDecimal!=0){	
        rem = tempDecimal% 2;  
        binary = (rem * place) + binary;  
        tempDecimal = tempDecimal/2;  
        place =place* 10  
    }  

    fmt.Println("Decimal number = ", decimal);  
    fmt.Println("Binary number = ", binary);
}