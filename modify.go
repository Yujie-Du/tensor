package tensor

//Reshape:修改数组形状，修改后总数据量大小必须与原来一致。shape:新的形状。
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
//Cut:切片，截取数据并重新排列形状。frontlimit:各维度索引前限(长度可小于原数组维度数)，backlimit:各维度索引后限(必须与前限长度一致)
func(a *Array)Cut(frontlimit,backlimit []int)(a2 *Array){
	if len(frontlimit)!=len(backlimit){
		panic("the length of frontlimit and backlimit should be same")
	}
	if len(frontlimit)==0{
		return a.Copy()
	}
	shape:=make([]int,a.level)
	for i:=range frontlimit{
		if frontlimit[i]<0||frontlimit[i]>backlimit[i]||backlimit[i]>a.shape[i]{
			panic("index out of range")
		}
		shape[i]=backlimit[i]-frontlimit[i]
	}
	for i:=len(frontlimit);i<a.level;i+=1{
		shape[i]=a.shape[i]
	}
	a2=new(Array)
	a2.init(shape,nil)
	if a2.count[0]==0{
		return
	}
	a.cutNext(frontlimit,backlimit,a2.data)
	return
}
//ExchangeAxis:交换两个维度。x1,x2:两个目标维度
func(a *Array)ExchangeAxis(x1,x2 int)(a2 *Array){
	if x1<0||x1>=a.level||x2<0||x2>=a.level{
		panic("axis out of range")
	}
	if x1==x2{
		return a.Copy()
	}
	return a.exchangeAxis(x1,x2)
}
//ReAxis:重新排列维度。axis:新的维度排序(长度可小于原数组维度数)
func(a *Array)ReAxis(axises []int)(a2 *Array){
	ch:=make([]bool,len(axises))
	for _,x:=range axises{
		ch[x]=true
	}
	for _,c:=range ch{
		if !c{
			panic("can not reaxis as request")
		}
	}
	return a.reAxis(axises)
}
func(a *Array)reAxis(axises []int)(a2 *Array){
	for i:=len(axises)-1;i>=0;i-=1{
		if axises[i]==i{
			axises=axises[:i]
		}else{
			break
		}
	}
	shape:=make([]int,a.level)
	for i,x:=range axises{
		shape[i]=a.shape[x]
	}
	for i:=len(axises);i<a.level;i+=1{
		shape[i]=a.shape[i]
	}
	a2=new(Array)
	a2.init(shape,nil)
	coord1,coord2:=make([]int,len(axises)),make([]int,len(axises))
	a.reAxisNext(coord1,coord2,axises,0,a2)
	return
}
func(a *Array)reAxisNext(last1,last2,axises []int,index int,a2 *Array){
	if index>=len(axises){
		data1:=a.dataInner(last1)
		data2:=a2.dataInner(last2)
		copy(data2,data1)
		return
	}else{
		for i:=0;i<a.shape[index];i+=1{
			last1[index]=i
			last2[axises[index]]=i
			a.reAxisNext(last1,last2,axises,index+1,a2)
		}
	}
}
func(a *Array)reShape(shape []int)(a2 *Array){
	data:=make([]float64,a.count[0])
	copy(data,a.data)
	a2=new(Array)
	a2.init(shape,data)
	return
}
func(a *Array)cutNext(frontlimit,backlimit []int,data []float64)(data2 []float64){
	if len(frontlimit)==0{
		copy(data,a.data)
		return data[len(a.data):]
	}
	for i:=frontlimit[0];i<backlimit[0];i+=1{
		data=a.ChildInner([]int{i}).cutNext(frontlimit[1:],backlimit[1:],data)
	}
	return data
}
func(a *Array)exchangeAxis(x1,x2 int)(a2 *Array){
	axises:=make([]int,a.level)
	for i:=range axises{
		axises[i]=i
	}
	axises[x1],axises[x2]=x2,x1
	return a.reAxis(axises)
}