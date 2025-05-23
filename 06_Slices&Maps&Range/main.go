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
//creating
m:= map[string]int{"price":40}
 m :=make(map[string]string)
 m["name"]="Rounak"
fmt.Println(m["name"])//rounak
fmt.Println(m["something"])//if key doesnot exist then return zero value
len(m)//1
delete(m,"name")
clear(m)
_, ok := m["price"] //check if the element there based on ok, any name we can give
k, ok :=m["price"] // here k returns value for the key 

//maps package also here 
maps.Equal(m1,m2)

--------------Range---------------

// iterating over DS
nums := []int {6,7,8,9,8}

for _,num := range nums{ // the first argument is index so for ind,num := range nums{}
   fmt.Println(num);
 }

//for map
for key,value := range map{}


//string 
for ind, char := range "golang" {} // print like this 0 103 it is unicode of character
//as go use unicode so posiible to see a byte will take two index as unicode bigger than ascii

