package tensor
import "fmt"
import "testing"

func TestTensor(t *testing.T){
	a:=NewArrayZeros([]int{2,3})
	fmt.Println(a)
}