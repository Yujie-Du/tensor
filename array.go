package main
//import "fmt"
//import "C"

func main(){}
type Array struct{
	level int
	shape []int
	count []int
	data []float64
}
func NewArrayFromFloat64(v float64)(a *Array){
	a=new(Array)
	a.init([]int{},[]float64{v})
	return
}
func NewArrayZeros(shape []int)(a *Array){
	a=new(Array)
	s2:=make([]int,len(shape))
	copy(s2,shape)
	a.init(s2,nil)
	return
}
func NewArray(shape []int,data []float64)(a *Array){
	a=new(Array)
	s2:=make([]int,len(shape))
	copy(s2,shape)
	a.init(s2,data)
	return
}
func(a *Array)init(shape []int,data []float64){
	if shape==nil{
		if data==nil{
			data=make([]float64,1)
		}
		shape=[]int{len(data)}
	}
	a.level=len(shape)
	a.shape=shape
	a.count=make([]int,a.level+1)
	a.count[a.level]=1
	for i:=a.level-1;i>=0;i-=1{
		if a.shape[i]<0{
			panic("ele of shape can not smaller than 0")
		}
		a.count[i]=a.count[i+1]*a.shape[i]
	}
	if data==nil{
		data=make([]float64,a.count[0])
	}else if len(data)!=a.count[0]{
		panic("num of data can not match the shape")
	}
	a.data=data
}
func(a *Array)Shape()(s []int){
	s=make([]int,a.level)
	copy(s,a.shape)
	return
}
func(a *Array)ReShape(shape []int)(a2 *Array){
	c:=1
	for _,s:=range shape{
		c*=s
	}
	if c!=a.count[0]{
		panic("can not change to this shape")
	}
	return a.reShape(shape)
}
func(a *Array)reShape(shape []int)(a2 *Array){
	data:=make([]float64,a.count[0])
	copy(data,a.data)
	a2=new(Array)
	a2.init(shape,data)
	return
}
func(a *Array)Level()(l int){
	return a.level
}
func(a *Array)Copy()(a2 *Array){
	a2=new(Array)
	shape:=make([]int,a.level)
	copy(shape,a.shape)
	data:=make([]float64,len(a.data))
	copy(data,a.data)
	a2.init(shape,data)
	return
}
func(a *Array)IfChildShape(a2 *Array)(b bool){
	if a2.level>a.level{
		return false
	}
	delta:=a.level-a2.level
	for i:=0;i<a2.level;i+=1{
		if a2.shape[i]!=a.shape[i+delta]{
			return false
		}
	}
	return true
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
func(a *Array)ChildInner(coord []int)(a2 *Array){
	if !a.checkCoord(coord){
		panic("index out of range")
	}
	return a.childInner(coord)
}
func(a *Array)childInner(coord []int)(a2 *Array){
	data:=a.dataInner(coord)
	lgh:=len(coord)
	return &Array{a.level-lgh,a.shape[lgh:],a.count[lgh:],data}
}
func(a *Array)checkCoord(coord []int)(b bool){
	for i,v:=range coord{
		if v<0||v>=a.shape[i]{
			return false
		}
	}
	return true
}
func(a *Array)transCoord(coord []int)(start,end int){
	for i,v:=range coord{
		start+=v*a.count[i+1]
	}
	return start,start+a.count[len(coord)]
}
func(a *Array)DataInner(coord []int)(data []float64){
	if !a.checkCoord(coord){
		panic("index out of range")
	}
	return a.dataInner(coord)
}
func(a *Array)dataInner(coord []int)(data []float64){
	start,end:=a.transCoord(coord)
	data=a.data[start:end]
	return
}
func(a *Array)SetInnerFloat64(coord []int,f float64){
	if !a.checkCoord(coord){
		panic("index out of range")
	}
	data:=a.dataInner(coord)
	for i:=range data{
		data[i]=f
	}
}
func(a *Array)SetChildInner(coord []int,a2 *Array){
	if !a.IfChildShape(a2){
		panic("different shape and can not expand")
	}
	a.setChildInner(coord,a2)
}
func(a *Array)setChildInner(coord []int,a2 *Array){
	for i:=0;i<len(a.data);i+=len(a2.data){
		copy(a.data[i:],a2.data)
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
func(a *Array)Cut(uplimit,downlimit []int)(a2 *Array){
	if len(uplimit)!=len(downlimit){
		panic("the length of uplimit and downlimit should be same")
	}
	if len(uplimit)==0{
		return a.Copy()
	}
	shape:=make([]int,a.level)
	for i:=range uplimit{
		if uplimit[i]<0||uplimit[i]>downlimit[i]||downlimit[i]>a.shape[i]{
			panic("index out of range")
		}
		shape[i]=downlimit[i]-uplimit[i]
	}
	for i:=len(uplimit);i<a.level;i+=1{
		shape[i]=a.shape[i]
	}
	a2=new(Array)
	a2.init(shape,nil)
	if a2.count[0]==0{
		return
	}
	a.cutNext(uplimit,downlimit,a2.data)
	return
}
func(a *Array)cutNext(uplimit,downlimit []int,data []float64)(data2 []float64){
	if len(uplimit)==0{
		copy(data,a.data)
		return data[len(a.data):]
	}
	for i:=uplimit[0];i<downlimit[0];i+=1{
		data=a.ChildInner([]int{i}).cutNext(uplimit[1:],downlimit[1:],data)
	}
	return data
}
func(a *Array)ExchangeAxis(x1,x2 int)(a2 *Array){
	if x1<0||x1>=a.level||x2<0||x2>=a.level{
		panic("axis out of range")
	}
	if x1==x2{
		return a.Copy()
	}
	return a.exchangeAxis(x1,x2)
}
func(a *Array)exchangeAxis(x1,x2 int)(a2 *Array){
	axis:=make([]int,a.level)
	for i:=range axis{
		axis[i]=i
	}
	axis[x1],axis[x2]=x2,x1
	return a.reAxis(axis)
}
func(a *Array)ReAxis(axis []int)(a2 *Array){
	ch:=make([]bool,len(axis))
	for _,x:=range axis{
		ch[x]=true
	}
	for _,c:=range ch{
		if !c{
			panic("can not reaxis as request")
		}
	}
	return a.reAxis(axis)
}
func(a *Array)reAxis(axis []int)(a2 *Array){
	for i:=len(axis)-1;i>=0;i-=1{
		if axis[i]==i{
			axis=axis[:i]
		}else{
			break
		}
	}
	shape:=make([]int,a.level)
	for i,x:=range axis{
		shape[i]=a.shape[x]
	}
	for i:=len(axis);i<a.level;i+=1{
		shape[i]=a.shape[i]
	}
	a2=new(Array)
	a2.init(shape,nil)
	coord1,coord2:=make([]int,len(axis)),make([]int,len(axis))
	a.reAxisNext(coord1,coord2,axis,0,a2)
	return
}
func(a *Array)reAxisNext(last1,last2,axis []int,index int,a2 *Array){
	if index>=len(axis){
		data1:=a.dataInner(last1)
		data2:=a2.dataInner(last2)
		copy(data2,data1)
		return
	}else{
		for i:=0;i<a.shape[index];i+=1{
			last1[index]=i
			last2[axis[index]]=i
			a.reAxisNext(last1,last2,axis,index+1,a2)
		}
	}
}