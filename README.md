# tensor
该包用于做多维数组的运算，并尽可能模仿python的numpy包的功能。
具体函数及方法暂请查看各文件注释。
## 已完成：
### array.go
该文件定义了Array结构体类型，几种初始化方法，索引获得子数组，查看修改数据等功能。
### modify.go
该文件定义了修改Array结构的方法，包括切片，维度变换，形状变换等功能
### algebra.go
该文件定义了有广播功能的代数运算的方法，包括加减乘除幂运算等，可以传入自定义函数进行带广播功能的运算。
### linear.go
该文件定义了做线性代数运算的工具，目前仅完成了线性变换。
### statis.go
该文件定义了几种统计工具，包括在任意维度上求均值最值标准差等
## 待完成：
### concurrency.go
并发支持工具
### random.go
随机初始化工具
### probability.go
概率分布工具
### learning/...
机器学习工具
### learing/gradient.go
反向梯度传播工具
### learning/network/...
神经网络工具