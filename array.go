package tensor
import "fmt"
//import "C"


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
func(a *Array)Length()(l int){
	if a.level==0{
		return 0
	}
	return a.shape[0]
}
func(a *Array)Count()(c int){
	return a.count[0]
}
func(a *Array)Shape()(s []int){
	s=make([]int,a.level)
	copy(s,a.shape)
	return
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
func(a *Array)SetFloat64Inner(coord []int,f float64){
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
func(a *Array)String()(str string){
	head,_:=a.stringNext()
	return head.Comb()
}
func(a *Array)stringNext()(head,end *strNode){
	if a.level==0{
		head=&strNode{fmt.Sprint(a.data[0]),nil}
		end=head
		return
	}
	if a.level==1{
		head=&strNode{fmt.Sprint(a.data),nil}
		end=head
		return
	}
	head=&strNode{"[",nil}
	end=head
	for i:=0;i<a.shape[0];i+=1{
		c:=a.ChildInner([]int{i})
		h,e:=c.stringNext()
		end.Next=h
		end=e
	}
	end.Next=&strNode{"]",nil}
	return
}
type strNode struct{
	Data string
	Next *strNode
}
func(s *strNode)Comb()(str string){
	temp:=make([]byte,s.count())
	s.combNext(temp)
	return string(temp)
}
func(s *strNode)count()(c int){
	if s==nil{
		return 0
	}
	return len(s.Data)+s.Next.count()
}
func(s *strNode)combNext(str []byte){
	if s==nil{
		return
	}
	copy(str,[]byte(s.Data))
	s.Next.combNext(str[len(s.Data):])
}