----------------------------------Sructs-----------------------------
import "time"
//like classes 

type order struct{
 createdAt time.time
  name string
}
//struct related function - we use first char of struct
func (o *order) chagename(name string){ //only one change we need pointer, in get we dont need
o.name=name;
}

//make a constructor function 
func newOrder(name string) *order{
  o:= order { name : "Rounak", }
  return &o;
}



func main(){
  var o order
  o:= order { name : "Rounak", } //not strict for all

  fmt.Println(o) //{rounak, {0 0 <nil>}}

  o.createdAt = time.Now()
  o.changeName("Rounsk 2")

  // taking from constructor
  newOrder := newOrder("Bablu") //struct dereference for us 

  lang:=struct{ //made a struct tha willuse only once
      isGood bool,
      name string
  }{true,"JavaScript"}
}

//struct embading : if a struct rfrerence another struct
type customer struct{
 id number
 order// order methods and fields will come
}
//assign it
 customer := customer{id:2}
 order := order{name:"RM", customer:customer}
  order := order{name:"RM", customer:customer{id:2}}
fmt.Println(order.customer.id)//2


_________________________Interface_________________________


type payment struct{}
func (p payment) makePayment(amount int){
  razorpayGateway:=razorpay{}
  razorpayGateway.pay(amount)
}

type razorpay struct{}
func (r razorpay) pay(amount int){
  //Pay logic
  fmt.Println("Making payment using Razorpay",amount)
}
function main(){
   newPayment:=payment{}
   newPayment.makePayment(1000)
/// now aassume some other paymnt method also need like stripe
// now for that we can make a stripe struct, function , object and acall its function but need to do in optimise way
// we have can make the instance in main and then pass the instance in payment struct
// like this
/* 
type payment struct{ gateway string}
main(){
 stripePaymentGateway:= stripe{}
 newPayment := payment{gateway stripePaymentGateway}
}*/
}

//Still need modification/
//interface is like a contract, end with "er"
type paymenter interface{
 pay(amount int)//for return pay(amount int)bool
}
type payment struct{
 gateway paymneter
}
type razorpay struct {}
func (r razorpay) pay(amount int){
 fmt.Println("AMount".amount)
}
type fake struct {}
func (f fake) pay(amount int){
 fmt.Println("AMount".amount)
}

func main(){
   razorpayPayment := razorpay{}
   // fakePayment:= fake{}
   newPayment := payment{
       gateway: razorpayPayment,//we can send here fakePayment also // use for future modification
    }
}

// here it tells in Open close principle we are not modifiing anything in the funcions
// we also use dependency inversion
_______________Enums______________________
//Enumarated types
// there is no built in enums, we use cont in go
type OrderStatus int
const(
 Recieved OrderStatus = iota //auto matically start from 0 and others will be increase
 //iota only for integer we can use string here also
 Confirm 
 Delivered 
)
func changeOrderStatus(status OrderStatus){
   fmt.Println("New Status",status);//0 
}
func main(){
   changeOrderStatus(Recieved)
}
________________Generics_________________


// func printSlice(items []int){
//   for _, item := range items{
//     fmt.Println(items)
//   }
// }
func printSlice[T any](items []T){ // by doing this anything it can take
  for _, item := range items{
    fmt.Println(items)
  }
}
//we can also write interface{} in place of any

/* we can also do this 
func printSlice[T bool | string | string](items []T){ // by doing this anything it can take
  for _, item := range items{
    fmt.Println(items)
  }
}
*/
funcion main(){
 printSLice([]int {1,2,3,4,5})
//assume same for string we have so have to make another function
}


//generics on struct
type stack[T any] struct{
  elements []T
}

myStack := stack[string]{
  elements: []string{"GO","stop","ENter"}
}


//in any we can also write comparable, build int types





