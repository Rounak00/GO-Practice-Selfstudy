
function add(a int,b int)int{
  return a+b;
}

function main(){
  result:=add(3,5);
  fmt.Println(result);
}
//now iff arguments like same so we can write this 
add(a,b int)

//go return multiple things
retu() (string,string,bool) { //return type in same order
 return "golang","lawda", true
}
// recieve those
lang1,lang2,lang3 := retu() //if dont wanna use any one of them so we can do _ there

//same like javaScript in Go also Func are first class citizen
func proc(fn func(a int)int){
  fn(1)
}
fn:=func(a int)int{return 2}
proc(fn)

//return function
func proces() func (a int)int{
 return func(a int)int{
   return 4
 }
}
call:=process()
call(1)


//variadic function
//so println can take a lot or infinity parameters that is call vaiadic
func nums(nums ...int)int{
 //nums as slice
}

// : for any type in go we do this interface{}

//send data from slice to arguments
num:=[]int{1,2,3,4,5,6,6};
function(num...);



// closures
func a() func()int {
 var i int=0;
 return func()int{
    i++
    return i
  }
}

iinc := a()
iinc() //1
iinc()//2
