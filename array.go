//tensor:golang多维数组运算工具，支持索引，切片，维度变换，形状变换，广播功能的代数运算，线性变换，任意维的数据统计等功能。
package tensor
import "fmt"
import "sync"
//import "C"

//Array:核心结构体，记录了数组的维数，形状，数据。
type Array struct{
	level int
	lock *sync.RWMutex
	shape []int
	count []int
	data []float64
}
//NewArrayFromFloat64:用浮点数初始化0维数组。
func NewArrayFromFloat64(v float64)(a *Array){
	a=new(Array)
	a.init([]int{},[]float64{v})
	return
}
//NewArrayZeros:建立数据全为0的数组。shape:数组形状。
func NewArrayZeros(shape []int)(a *Array){
	a=new(Array)
	a.init(shape,nil)
	return
}
//NewArray:用原始数据的副本建立数组。shape:数组形状，data:原始数据
func NewArray(shape []int,data []float64)(a *Array){
	a=new(Array)
	a.init(shape,data)
	return
}
//Length:获得数组的直接子数组的数量。
func(a *Array)Length()(l int){
	if a.level==0{
		return 0
	}
	return a.shape[0]
}
//Count:获得数组全部数据长度。
func(a *Array)Count()(c int){
	return a.count[0]
}
//Shape:获得数组形状。
func(a *Array)Shape()(s []int){
	s=make([]int,a.level)
	copy(s,a.shape)
	return
}
//Copy:拷贝该数组，获得相同但互不关联的新数组。
func(a *Array)Copy()(a2 *Array){
	a.lock.RLock()
	defer a.lock.RUnlock()
	a2=new(Array)
	a2.init(a.shape,a.data)
	return
}
//IfChildShape:判断目标数组的形状是否与主数组任意深度的子数组形状相同。a2:目标数组
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
//ChildInner:索引，获取任意深度的子数组，获得的子数组与主数组共享底层数据。coord:索引
func(a *Array)ChildInner(coord []int)(a2 *Array){
	if !a.checkCoord(coord){
		panic("index out of range")
	}
	return a.childInner(coord)
}
//SetFloat64Inner:索引获得任意深度的一维底层数据，并全部替换为目标浮点数。coord:索引，f:目标浮点数。
func(a *Array)SetFloat64Inner(coord []int,f float64){
	if !a.checkCoord(coord){
		panic("index out of range")
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	data:=a.dataInner(coord)
	for i:=range data{
		data[i]=f
	}
	return
}
//SetChildInner:索引获得任意深度子数组，并替换为目标数组，可广播。coord:索引，a2:目标数组
func(a *Array)SetChildInner(coord []int,a2 *Array){
	if !a.IfChildShape(a2){
		panic("different shape and can not expand")
	}
	a.setChildInner(coord,a2)
}
//String:获得数组的字符串表示。
func(a *Array)String()(str string){
	a.lock.RLock()
	defer a.lock.RUnlock()
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
	end=end.Next
	return
}
type strNode struct{
	Data string
	Next *strNode
}
func(s *strNode)Comb()(str string){
	//s.checkLine(false)
	temp:=make([]byte,s.count())
	s.combNext(temp)
	return string(temp)
}
func(s *strNode)checkLine(check bool)(s2 *strNode){
	if check&&s.Data[0]=='['{
		s2=&strNode{"\n",s}
	}else{
		s2=s
	}
	if s.Next!=nil{
		s.Next=s.Next.checkLine(s.Data[len(s.Data)-1]==']')
	}
	return
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
func(a *Array)init(shape []int,data []float64){
	if shape==nil{
		if data==nil||len(data)==1{
			a.shape=[]int{}
		}else{
			a.shape=[]int{len(data)}
		}
	}else{
		a.shape=make([]int,len(shape))
		copy(a.shape,shape)
	}
	a.level=len(a.shape)
	a.lock=new(sync.RWMutex)
	a.count=make([]int,a.level+1)
	a.count[a.level]=1
	for i:=a.level-1;i>=0;i-=1{
		if a.shape[i]<0{
			panic("ele of shape can not smaller than 0")
		}
		a.count[i]=a.count[i+1]*a.shape[i]
	}
	a.data=make([]float64,a.count[0])
	if data!=nil{
		copy(a.data,data)
	}
}
func(a *Array)childInner(coord []int)(a2 *Array){
	data:=a.dataInner(coord)
	lgh:=len(coord)
	return &Array{a.level-lgh,a.lock,a.shape[lgh:],a.count[lgh:],data}
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
func(a *Array)dataInner(coord []int)(data []float64){
	a.lock.RLock()
	defer a.lock.RUnlock()
	start,end:=a.transCoord(coord)
	data=a.data[start:end]
	return
}
func(a *Array)setChildInner(coord []int,a2 *Array){
	a.lock.Lock()
	defer a.lock.Unlock()
	for i:=0;i<len(a.data);i+=len(a2.data){
		copy(a.data[i:],a2.data)
	}
}