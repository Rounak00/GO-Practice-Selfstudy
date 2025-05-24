//Mutual execution - for race condition
// multiple processes change same resource
import "sync"
import "fmt"
type post struct{
 view int
}
func (p *post)inc (wg *sync.WaitGroup){
 defer wg.Done()
  p.view+=1
}
func main(){
 var wg sync.WaitGroup
 mypost:=post{view:0}
for i:=0; i<100;i++{
  wg.Add(1)
  go mypost.inc(&wg)
}
wg.Wait()
fmt.Println(mypost.view)
}
// it gives wrong count______________________________________
type post struct{
 view int
  mu sync.Mutex
}
func (p *post)inc (wg *sync.WaitGroup){
 defer func() {wg.Done() p.mu.Unlock()}
  p.mu.Lock()
  p.view+=1
  //why unlock is not here bcz if error came then it will never unlock
  // also this previous two lines are bottle neck as those are not give power of concurrency
// also best practice give lock to only to this that needed
}
func main(){
 var wg sync.WaitGroup
 mypost:=post{view:0}
for i:=0; i<100;i++{
  wg.Add(1)
  go mypost.inc(&wg)
}
wg.Wait()
fmt.Println(mypost.view)
}
