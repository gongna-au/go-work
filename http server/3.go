//定义一个函数不接收参数的函数getCounter它的返回值是这个func () int 
func getCounter () func () int {
//初始化一个整型变量num
    num :=0
//在这个函数里面定义一个有返回值的匿名函数，如果我要用这个函数，我找个值来接收它
	count=func ()int{num++;return num }//利用分号可以把两个语句放一行
//用我找的这个变量把我们在里面设的函数返回

return count 
}
a :=getCounter()//计数器函数a
b :=getCounter()//计数器函数b 
println (a())
println (a())
println (b())
println (b())