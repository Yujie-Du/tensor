package main

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
func(a *Array)Dot(a2 *Array)(a3 *Array){
	if a2.level==0||a.level==0{
		return a.Mul(a2)
	}
	if a2.level>2{
		panic("no function to do Dot opt for array that have more than 2 level")
	}
	if a.shape[a.level-1]!=a2.shape[0]{
		panic("can not do dot opt with wrong shape")
	}
	if a2.level==1{
		return a.dot1(a2)
	}else{
		return a.dot2(a2)
	}
}
func(a *Array)dot1(a2 *Array)(a3 *Array){
	a3=new(Array)
	shape:=make([]int,a.level-1)
	copy(shape,a.shape[:a.level-1])
	a3.init(shape,nil)
	lgh:=a2.shape[0]
	for i:=0;i<a.count[0]/lgh;i+=1{
		a3.data[i]=float64dot1(a.data[i*lgh:(i+1)*lgh],a2.data)
	}
	return
}
func float64dot1(l1,l2 []float64)(r float64){
	for i:=0;i<len(l1);i+=1{
		r+=l1[i]*l2[i]
	}
	return
}
func(a *Array)dot2(a2 *Array)(a3 *Array){
	a3=new(Array)
	shape:=make([]int,a.level)
	copy(shape,a.shape[:a.level-1])
	shape[a.level-1]=a2.shape[1]
	a3.init(shape,nil)
	l1:=a2.shape[0]
	l2:=a2.shape[1]
	for i:=0;i<a.count[0]/l1;i+=1{
		copy(a3.data[i*l2:(i+1)*l2],float64dot2(a.data[i*l1:(i+1)*l1],a2.data))
	}
	return
}
func float64dot2(l1 []float64,l2 []float64)(l3 []float64){
	nl:=len(l2)/len(l1)
	l3=make([]float64,nl)
	for i:=0;i<len(l1);i+=1{
		for j:=0;j<nl;j+=1{
			l3[j]+=l1[i]*l2[i*nl+j]
		}
	}
	return
}