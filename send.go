package main

import "fmt"

/*
你要发金币了，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后发出去了多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)

}


func dispatchCoin()int  {
	money:=make([]int, len(users))
	sum:=0

	for i := 0; i < len(users); i++ {

		distribution[users[i]]=money[i]
	}
	for k,v:=range distribution{
		c:=[]byte(k)
		for _,ch:=range c{
			if ch=='e'||ch=='E'{
				v+=1
			}else if ch=='i'||ch=='I'{
				v+=2
			}else if ch=='o'||ch=='O'{
				v+=3
			}else if ch=='u'||ch=='U'{
				v+=4
			}
		}
		//发现这样才可以保存v的值
		distribution[k]=v
		fmt.Println(k,v,"个金币")
		sum+=v
	}
	fmt.Println(distribution)
	return sum
}
