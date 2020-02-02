package tensor

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