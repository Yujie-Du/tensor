package tensor
import "math/rand"
import "time"

func init(){
	rand.Seed(time.Now().UnixNano())
}
//NewArrayRandom:随机生成start至end之间的随机数组成的数组。shape:数组形状，start,end:上下限
func NewArrayRandom(shape []int,start,end float64)(a *Array){
	delta:=end-start
	f:=func()float64{
		return rand.Float64()*delta+start
	}
	return newArrayRandomFunc(shape,f)
}
//NewArrayNorm:随机生成期望值mean标准差std的随机数组成的数组。shape:数组形状，mean:期望，std:标准差
func NewArrayNorm(shape []int,mean,std float64)(a *Array){
	f:=func()float64{
		return rand.NormFloat64()*std+mean
	}
	return newArrayRandomFunc(shape,f)
}
func newArrayRandomFunc(shape []int,f func()float64)(a *Array){
	a=NewArray(shape,nil)
	for i:=range a.data{
		a.data[i]=f()
	}
	return
}