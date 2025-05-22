package main

import "fmt"

//for i s onlu loop in go
function main(){
  //while loop
  i:=1
  for i < 3 {
    fmt.Println(i)
    i=i+1
    //i++ also correct
  }


  //infinte loop
  for {
    fmt.Println("Infinity")
   }

  // classic for loop
  for i:=0; i<3; i++ {
    fmt.Println(i);
  }

  // we have same break and continue here

  //using range
  for i:= range 3 {...} //it will run from 0 till 3 that is >3 not >=3


  age:=18;
  if age>18{
     fmt.Println("Person is an adult")
  } else if age==18{
     fmt.Println("Person is exact 18")
  } else {
     fmt.Println("Peron in child)
  }


 // we can use like this,  if condition 1 || condition 2 {} / condition 1 && condition 2 {}

// we can also do this
if age:=15; age >18 { fmt.Println("He is a boy") }
 

// go doesnot have any ternary operator
}
