package main

import (
	"fmt"
	"github.com-go/wzyx/gs"
)

//初始化方法
func init() {

}

func main() {
	fmt.Println("请输入1开始游戏")
	var stats int
	fmt.Scan(&stats)
	if stats == 1 {
		fmt.Println("创建角色")
	} else {
		fmt.Println("退出")
		return
	}
	fmt.Println("创建怪物。。。")
	cjgs(1)

}

//创建怪兽
func cjgs(t int) gs.Gs {
	g := new(gs.Gs)
	g.Lx = "狗"
	g.Xz = "兽类"
	g.Gj = 2
	g.Xl = 30
	g.Fy = 2
	g.Zl = 0
	switch t {
	case 1:
		fmt.Printf("当前生成的怪兽：%v\n", *g)
		return *g
	default:
		fmt.Println("未知错误")
		return *g
	}
}
