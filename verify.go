package main

import "fmt"

func main() {
	var(
		ac,psw,ac1,psw1 string
	)
	m:=make(map[string]string)

	fmt.Println("注册账号\n请输入账号:")
	fmt.Scan(&ac)
	fmt.Println("请输入密码:")
	fmt.Scan(&psw)
	m[ac]=psw
	fmt.Println("注册成功，请重新登陆")
	fmt.Println("登陆\n请输入账号:")
	fmt.Scan(&ac1)
	fmt.Println("请输入密码:")
	fmt.Scan(&psw1)
	for k,v:=range m{
		if k==ac1&&v!=psw1{
			fmt.Println("密码错误")
		}else if k==ac1&&v==psw1{
			fmt.Println("登陆成功")
		}else{
			fmt.Println("该账号未注册")
		}
	}
}
