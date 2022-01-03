package main
import "fmt"
func factorial(num int) int{
  if (num==1){
    return num;
  } 
  return num*(factorial(num-1))
}
func main(){
  var num int 
  fmt.Scanln(&num)
  fmt.Println(factorial(num))
}
