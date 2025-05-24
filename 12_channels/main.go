//channel is like a pipe
// send data between go routines

main(){
 message := make(chan string)

 message <- "Ping" //channels are blocking until second side ready to recieve
msg:= <- message
gmt.Println(msg)
}
//now it will give a fatal error of deadlock


func processNum(num chan int){
fmt.Println(<-num)
   
}
main(){
 numchan:= make(chan int)
 go processNum(numchan)
numchan <- 5

time.Sleep(time.Second *2)
}
// ____________________________________________________________________________________
//also do it like this
for{
  numchan<-rand.Intn(100)
}
// access like this 
for num := range numchan {
  fmt.Println(<-numchan)
  time.Sleep(time.Scond)
}

// ______________________________________________________________
func sum(result chan int, num1 int, num2 int){
  numResult := num1 +num2
  result <- numResult
}
main(){
   result := make(chan int)
   go sum(result, 4, 5)
  res:= <-result
  fmt.Println(res)
}

//------------------------------------------------------------------------
//use channel for wait group
func task(done chan bool){
 defer func () {done <- true}
 fmr.Println("PRocessing ....")
}
main(){
    done :=make(chan bool)
   go task(done)
   <- done //it will block until something send
}
//all these are un bufferd channel - one time one value will send we can send only next when previous vaue revieved
// /---------------------------------------------------------------------------
//buffered channels - we can send limited data without blocking

emailChan := make(chan String,100)
emailChan <- "one"
emailChan<- "two"

fmt.println(<- emailChan)
fmt.println(<- emailChan)
//no deadlock here


//Implement email sender queue system
func emailSender(email chan string,done chan bool){
 defer func() {done <- true}
  for e := range email{
    fmt .Println("Done email => ",e)
    time.Sleep(time.Second)
  }
}
func main(){
 emailChan := make(chan String,100)
  done := make(chan bool)
go emailSender(emailChan,done)
for i :=0; i<10; i++ {
  emailChan <- fmtSprintf("%d@gmail.com", i)
}
fmt.Println("Done sending")
close(emailChan)//buffered channel may cause the deadlock and done will stay waiting
<-done
}
//----------------------------------------------------------------------
//Listen to multiple channel
func main(){
  chan1 := make(chan int)
  chan2 :=make(chan string)

 go func(){
  chan1 <- 10
 }()

  go func(){
   chan2 <- "pong" 
 }()

 for i:=0; i<2; i++ {
  select{
     case chan1val := <-chan1
               fmt.Println("Channel1 value ->",chan1val)
     case chan2val := <-chan2
               fmt.Println("Channel2 value ->",chan2val)
   }
 }
}
//scope a channel -> like make it recieve only

emailSender(emailChan <-chan string) // means only can send

