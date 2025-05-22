package main
import "ftm"


//numbered sequesnce of specific length
function main(){
   var nums [4]int 
   //we can also do it
    nums := [3]int{1,2,3}

   nums[0] =1
   fmt.Println(nums[0])//1
   fmt.Println(nums)//[1 0 0 0]
   fmt.Println(len(nums)) //4

   var vals [4]bool
   fmt.Println(vals) //false
  
   var sals []string
   fmt.Println(sals) //[    ] (empty char/string)

   // this emoty things called Zeroed values


   //2D array
   num:=[2][2]int{{1,2},{2,3}}

  // but we mostluy use slices in Go

/* use it when you know the size in prior
   helop in memory optimization
   give constant memory size */
} 
