package tensor
import "math"

//MapForUnit:广播自定义的单变量运算，运算结果组成新数组。f:自定义的运算函数
func(a *Array)MapForUnit(f func(v float64)float64)(a2 *Array){
	a2=a.Copy()
	a2.MapForUnitInner(f)
	return
}
//MapForUnitInner:广播自定义的单变量运算，运算结果替代原数组数据。f:自定义的运算函数
func(a *Array)MapForUnitInner(f func(v float64)float64){
	for i,v:=range a.data{
		a.data[i]=f(v)
	}
}
//Abs:广播绝对值运算，运算结果组成新数组。
func (a *Array)Abs()(a2 *Array){
	f:=func(v float64)float64{return math.Abs(v)}
	return a.MapForUnit(f)
}
//Log:广播对数运算，运算结果组成新数组。
func (a *Array)Log()(a2 *Array){
	f:=func(v float64)float64{return math.Log(v)}
	return a.MapForUnit(f)
}
//Sin:广播正弦运算，运算结果组成新数组。
func (a *Array)Sin()(a2 *Array){
	f:=func(v float64)float64{return math.Sin(v)}
	return a.MapForUnit(f)
}
//Sinh:广播反正弦运算，运算结果组成新数组。
func (a *Array)Sinh()(a2 *Array){
	f:=func(v float64)float64{return math.Sinh(v)}
	return a.MapForUnit(f)
}
//Cos:广播余弦运算，运算结果组成新数组。
func (a *Array)Cos()(a2 *Array){
	f:=func(v float64)float64{return math.Cos(v)}
	return a.MapForUnit(f)
}
//Cosh:广播反余弦运算，运算结果组成新数组。
func (a *Array)Cosh()(a2 *Array){
	f:=func(v float64)float64{return math.Cosh(v)}
	return a.MapForUnit(f)
}
//Tan:广播正切运算，运算结果组成新数组。
func (a *Array)Tan()(a2 *Array){
	f:=func(v float64)float64{return math.Tan(v)}
	return a.MapForUnit(f)
}
//Tanh:广播反正切运算，运算结果组成新数组。
func (a *Array)Tanh()(a2 *Array){
	f:=func(v float64)float64{return math.Tanh(v)}
	return a.MapForUnit(f)
}
//MapForUnit:广播自定义的二元运算，运算结果组成新数组。a2:第二操作数组，f:自定义的运算函数
func(a *Array)OptForUnit(a2 *Array,f func(v1,v2 float64)float64)(a3 *Array){
	if a2.level>a.level{
		f2:=func(v1,v2 float64)float64{return f(v2,v1)}
		return a2.OptForUnit(a,f2)
	}
	if !a.IfChildShape(a2){
		panic("different shape and can not expand")
	}
	a3=a.Copy()
	a3.optForUnitInner(a2,f)
	return
}
//MapForUnitInner:广播自定义的二元运算，运算结果替代原数组数据。a2:第二操作数组，f:自定义的运算函数
func(a *Array)OptForUnitInner(a2 *Array,f func(v1,v2 float64)float64){
	if !a.IfChildShape(a2){
		panic("different shape and can not expand")
	}
	a.optForUnitInner(a2,f)
}
//Add:广播数组对数组的加法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)Add(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1+v2}
	return a.OptForUnit(a2,f)
}
//AddFloat64:广播数组对浮点数加法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)AddFloat64(f float64)(a2 *Array){
	return a.Add(NewArrayFromFloat64(f))
}
//Rdt:广播数组对数组的减法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)Rdt(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1-v2}
	return a.OptForUnit(a2,f)
}
//RdtFloat64:广播数组对浮点数减法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)RdtFloat64(f float64)(a2 *Array){
	return a.Rdt(NewArrayFromFloat64(f))
}
//Mul:广播数组对数组的乘法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)Mul(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1*v2}
	return a.OptForUnit(a2,f)
}
//MulFloat64:广播数组对浮点数乘法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)MulFloat64(f float64)(a2 *Array){
	return a.Mul(NewArrayFromFloat64(f))
}
//Div:广播数组对数组的除法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)Div(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1/v2}
	return a.OptForUnit(a2,f)
}
//DivFloat64:广播数组对浮点数除法运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)DivFloat64(f float64)(a2 *Array){
	return a.Div(NewArrayFromFloat64(f))
}
//Pow:广播数组对数组的乘方运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)Pow(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return math.Pow(v1,v2)}
	return a.OptForUnit(a2,f)
}
//PowFloat64:广播数组对浮点数乘方运算，运算结果组成新数组。a2:第二操作数组
func(a *Array)PowFloat64(f float64)(a2 *Array){
	return a.Pow(NewArrayFromFloat64(f))
}
func(a *Array)optForUnitInner(a2 *Array,f func(v1,v2 float64)float64){
	c:=a2.count[0]
	for i:=range a.data{
		a.data[i]=f(a.data[i],a2.data[i%c])
	}
}