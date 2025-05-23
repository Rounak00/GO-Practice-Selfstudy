----------Slice-------
// slice are dynamic and have usefullmethods
function map(){

	// in go null is nil so in slice uninitilised nill
	var num[]int
    fmt.Println(num)//[] empty slice
    fmt.Println(len(num))
    num:=[]int{} //another way intialization 
	var num = make([]int,2); //[0,0] 
	//here the third parameter is initial capacity
    // after intial capacity also if i push elements it will increase



    cap(num)//2 max size of number can fit but it can store more also
	num=append(num,2);//push back
    

	//copy slice
    car nums2=make([]int,len(num))
	copy(nums2,num)


	//slice operator
	var nums =[]int{1,3,4,5}
	fmt.Println(nums[0:1])//1 [from:to]
    fmt.Println(nums[:1]) //we can also do this
    fmt.Println(nums[1:])


	//slice package
	var nums1=[]int{1,2}
	var nums2=[]int{2,3}
	fmt.Println(slices.Equal(nums1,nums2)) //return fslse


	//2d slices
	var num=[][]int{{1,2},{2,3}}
}

--------------Map-----------------
--------------Range---------------
