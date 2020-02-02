package tensor
import "math"

func(a *Array)MapForUnit(f func(v float64)float64)(a2 *Array){
	a2=a.Copy()
	a2.MapForUnitInner(f)
	return
}
func(a *Array)MapForUnitInner(f func(v float64)float64){
	for i,v:=range a.data{
		a.data[i]=f(v)
	}
}
func Abs(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Abs(v)}
	return a.MapForUnit(f)
}
func Log(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Log(v)}
	return a.MapForUnit(f)
}
func Sin(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Sin(v)}
	return a.MapForUnit(f)
}
func Sinh(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Sinh(v)}
	return a.MapForUnit(f)
}
func Cos(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Cos(v)}
	return a.MapForUnit(f)
}
func Cosh(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Cosh(v)}
	return a.MapForUnit(f)
}
func Tan(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Tan(v)}
	return a.MapForUnit(f)
}
func Tanh(a *Array)(a2 *Array){
	f:=func(v float64)float64{return math.Tanh(v)}
	return a.MapForUnit(f)
}
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
func(a *Array)OptForUnitInner(a2 *Array,f func(v1,v2 float64)float64){
	if !a.IfChildShape(a2){
		panic("different shape and can not expand")
	}
	a.optForUnitInner(a2,f)
}
func(a *Array)optForUnitInner(a2 *Array,f func(v1,v2 float64)float64){
	c:=a2.count[0]
	for i:=range a.data{
		a.data[i]=f(a.data[i],a2.data[i%c])
	}
}
func(a *Array)Add(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1+v2}
	return a.OptForUnit(a2,f)
}
func(a *Array)AddFloat64(f float64)(a2 *Array){
	return a.Add(NewArrayFromFloat64(f))
}
func(a *Array)Rdt(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1-v2}
	return a.OptForUnit(a2,f)
}
func(a *Array)RdtFloat64(f float64)(a2 *Array){
	return a.Rdt(NewArrayFromFloat64(f))
}
func(a *Array)Mul(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1*v2}
	return a.OptForUnit(a2,f)
}
func(a *Array)MulFloat64(f float64)(a2 *Array){
	return a.Mul(NewArrayFromFloat64(f))
}
func(a *Array)Div(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return v1/v2}
	return a.OptForUnit(a2,f)
}
func(a *Array)DivFloat64(f float64)(a2 *Array){
	return a.Div(NewArrayFromFloat64(f))
}
func(a *Array)Pow(a2 *Array)(a3 *Array){
	f:=func(v1,v2 float64)float64{return math.Pow(v1,v2)}
	return a.OptForUnit(a2,f)
}
func(a *Array)PowFloat64(f float64)(a2 *Array){
	return a.Pow(NewArrayFromFloat64(f))
}