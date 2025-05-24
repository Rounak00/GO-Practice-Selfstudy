//go routines are light wait threads

func task(id int){
 fmt.Println("Doint task ->", id);
}
func main(){
  for i := 0; i< 10; i++ {
    go task(i);//run it in a thread, go to schedular then call
  }
  time.Sleep(time.Second*2)// so main dont end before this task calls end
}


// we can also do this way
go func(i int){
  fmt.PRintln(i)
}(i)
