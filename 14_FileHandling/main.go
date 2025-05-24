
import "os"
func main(){
  f, err := os.Dopen("example.txt")
  if err != nil{ panic(err)}
 
  fileInfo,err:=f.Stat() //return file info or path error
   if err != nil{ panic(err)}
  fmt.Println(fileInfo.Name())
   fmt.Println(fileInfo.IsDir())
  fmt.Println(fileInfo.Size())
  fmt.Println(fileInfo.Mode())
  fmt.Println(fileInfo.ModTime())
  

//Read file
f, err := os.Open("example.txt")
  if err != nill{ panic(err)}
defer f.CLose()// we always have to close as well
buf := make([]byte, 12)
d,err := f.Read(buf)// d gives length
 if err != nil{ panic(err)}
//wrap it in a for loo of len(buf)
println(d,string(buff[i]))


//simplest wy to read for small files
data,err := os.ReadFile("example.txt")
if err!=nil {panic(err)}
fmt.Println(string(data))

//for big files we use streams

//read folders 
dir,err:=os.Open(".")//currewnt
if err!=nil{panic(err)}
defer dit.Close()
fileInfo,err := dir.ReadDir(1)//return slice of files name and only one bcz of 1
//-1 it will give all

//create file 
f,err:=os.Create("example2.txt")
if err !=nil {panic(err)}
defer f.close()
f.WriteString("Hi Go") //it will always append


-------------------------------------------------------
//add byte
bytes := []byte("Hello Rounak")
f.Write(bytes)

//using strea transfer data one file to another
source,err := os.Open("example.txt")
defer source.CLose()

dest,err := os.Create("example2.txt")
if err != nil {panic(err)}
reader :=bufio.NewReader(source)//bufio is a package
writer :=bufio.NewWriter(dest)

for {
 b,err:=reader.ReadByte()
 if err!=nil {
    if err.Error() != "EOF" {
           panic(err)
       }
  break
   }
   er:=writer.WriteByte(b) 
   if er!=nil {panic(er)}
  }
 writer.Flush()
fmt.Println("Written to new file")


//delete file 

err:=os.Remove("example2.txt")
if err!=nil {panic(err)}
}
