package main;

import "fmt"
import "time"
function main(){
  //simple switch
  switch i {
    case 1 :
          fmt.Println("One")
   case 2 :
         fmt.Println("Two)
    default :
         fmt.Println("Three")     
  }

  //Multiple condition switch
  switch time.Now().Weekday(){
    case time.Saturday, time.Sunday : 
           fmt.Println("This is a Weekend holiday");
  }


  //type switch
  whoami:= func(i interface{}){
    switch t := i.(type) {  //also do thi switch i.(type)
      case int:
           fmt.Pritnln("Its an integer")
      case string :
           fmt.Println("It's an string")
      
     default : 
           fmt.Println("It's an other", t)
   }}
  whoami("Hello") //It's an string
}
