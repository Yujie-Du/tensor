package tensor
import "fmt"
import "testing"

func TestTensor(t *testing.T){
	a,_:=A(3)
	fmt.Println(a)
}

func A(b int)(a int,err error){
	if b>0{
		return b,fmt.Errorf("b>0")
	}
	return b,nil
}