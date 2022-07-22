package main

import "fmt"

func main()  {
	fmt.Println("请输入1开始游戏")
	var stats int
	fmt.Scan(&stats)
	if stats == 1 {
		fmt.Println("创建角色")
	} else {
		fmt.Println("退出")
		return
	}
}
