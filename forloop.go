package main
import "fmt"
func main(){
  for i,j:=1,10;i<=10 && j>=1;i,j=i+1,j-1{
    fmt.Print(i, " ");
    fmt.Println(j);
  }
}
