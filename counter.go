package main

import "fmt"

//计算器
type counter struct {
	a,b,res float64
	oper string

}

//运算
func (c counter)operate()float64{
	if c.oper=="+"{
		c.res=c.a+c.b
	}else if c.oper=="-"{
		c.res=c.a-c.b
	}else if c.oper=="*"{
		c.res=c.a*c.b
	}else if c.oper=="/" {
		c.res=c.a/c.b
	}else {
		fmt.Println("输入的运算符有错，请重新运行")
	}
	return c.res
}


func main() {
var(
	a,b float64
	oper string
)
fmt.Println("input:")
fmt.Scan(&a,&oper,&b)
c:=counter{a:a,oper:oper,b:b}
fmt.Println("output:\n",c.operate())

}
