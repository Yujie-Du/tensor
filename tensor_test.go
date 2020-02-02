package main
import "fmt"
import "testing"

func TestTensor(t *testing.T){
	a:=NewArray([]int{2,3},[]float64{1,2,3,4,5,6})
	fmt.Println(a.Std(0))
}