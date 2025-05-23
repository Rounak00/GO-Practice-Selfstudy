
func changeNum(num *int){
 *num=5
}

func main(){
 num:=3
changeNum(&num)
fmt.Println(num)//3
}
