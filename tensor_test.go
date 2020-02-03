package tensor
import "fmt"
import "testing"

func TestRandom(t *testing.T){
	//a:=NewArray([]int{2,3,4},nil)
	//a:=NewArrayRandom([]int{2,3},2,3)
	a:=NewArrayNorm([]int{20},0,1)
	fmt.Println(a)
}