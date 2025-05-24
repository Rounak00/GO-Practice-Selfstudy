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


// Wait group for using instead of sleep

func task(id int, w *sync.WaitGroup){
 defer w.Done() // will remove that 1
 fmt.Println("Doint task ->", id);
}
func main(){
var wg sync.WaitGroup
  for i := 0; i< 10; i++ {
   wg.add(1)
    go task(i, &wg);
  }
wg.Wait()
}
