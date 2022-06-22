package main

import (
	"fmt"
)

func main() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a 变量的地址是：%x\n", &a)
	fmt.Printf("ip 存储的地址是： %x\n", ip)
	fmt.Printf("*ip 变量的值：%d\n", *ip)

	var ip2 *int
	fmt.Printf("prt的值为：%x\n", ip2)
	if ip2 == nil {
		fmt.Print("nil")
	}
}
