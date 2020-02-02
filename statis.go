package main
import "math"
func(a *Array)Std(axis int)(a2 *Array){
	f:=func(data []float64)(r float64){
		mean:=0.0
		for _,d:=range data{
			mean+=d
		}
		mean/=float64(len(data))
		sum:=0.0
		for _,d:=range data{
			sum+=math.Pow((d-mean),2)
		}
		sum/=float64(len(data))
		return math.Sqrt(sum)
	}
	return a.OptForStatis(axis,f)
}
func(a *Array)Min(axis int)(a2 *Array){
	f:=func(data []float64)(r float64){
		r=data[0]
		for _,d:=range data[1:]{
			if d<r{
				r=d
			}
		}
		return
	}
	return a.OptForStatis(axis,f)
}
func(a *Array)Max(axis int)(a2 *Array){
	f:=func(data []float64)(r float64){
		r=data[0]
		for _,d:=range data[1:]{
			if d>r{
				r=d
			}
		}
		return
	}
	return a.OptForStatis(axis,f)
}
func(a *Array)Mean(axis int)(a2 *Array){
	f:=func(data []float64)(r float64){
		for _,d:=range data{
			r+=d
		}
		r/=float64(len(data))
		return 
	}
	return a.OptForStatis(axis,f)
}
func(a *Array)Sum(axis int)(a2 *Array){
	f:=func(data []float64)(r float64){
		for _,d:=range data{
			r+=d
		}
		return 
	}
	return a.OptForStatis(axis,f)
}
func(a *Array)OptForStatis(axis int,f func(d []float64)float64)(a2 *Array){
	if axis<0||axis>=a.level{
		panic("axis out of range")
	}
	return a.optForStatis(axis,f)
}
func(a *Array)optForStatis(axis int,f func(d []float64)float64)(a2 *Array){
	shape:=make([]int,a.level-1)
	copy(shape[:axis],a.shape[:axis])
	copy(shape[axis:],a.shape[axis+1:])
	a2=NewArray(shape,nil)
	if a2.count[0]==0{
		return
	}
	a.optForStatisNext(axis,f,a2.data)
	return
}
func(a *Array)optForStatisNext(axis int,f func(d []float64)float64,data []float64)(data2 []float64){
	if axis>0{
		for i:=0;i<a.shape[0];i+=1{
			c:=a.childInner([]int{i})
			data=c.optForStatisNext(axis-1,f,data)
		}
		return
	}
	temp:=make([]float64,a.shape[0])
	for j:=0;j<a.count[1];j+=1{
		for i:=0;i<a.shape[0];i+=1{
			temp[i]=a.data[i*a.count[1]+j]
		}
		data[j]=f(temp)
	}
	return data[a.count[1]:]
}