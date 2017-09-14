package main

import "fmt"
import "time"

func main(){

    var i = 1

    switch i {
        case 1: fmt.Println( "One is the case")
        case 2 : fmt.Println( " Two is the case")
        case 3 : fmt.Println(" Three is the case")
    }
    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday :
                fmt.Println("its a weekend")
        default:
                fmt.Println("its a weekday")
    }
    
   /* switch time.Now().Hour() {
        case time.Hour() > 12 : 
            fmt.Println("Its after noon")
        default : 
                fmt.Println("its before noon")
    }*/
}