package main

import (
	"fmt"
	"github.com-go/wzyx/gs"
	"github.com-go/wzyx/js"
	"github.com-go/wzyx/zl"
	"time"
)

//初始化方法
func init() {

}

func main() {
	fmt.Println("请输入1开始游戏")

	var stats int
	fmt.Scan(&stats)
	if stats == 1 {
		fmt.Println("创建角色,玩家：", js.Name, "加载完成")
	} else {
		fmt.Println("退出")
		return
	}
	fmt.Println("正在加载中。。。")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("W,向前,S,向后，A,向左,D,向右,J,进攻,K,逃跑,Q,退出,I,菜单")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("加载完成！！！\n当前坐标:", js.Jzx["X"], ",", js.Jzx["Y"])
	b := true
	for b {
		fx := zl.Fx()
		switch fx {
		case 0:
			time.Sleep(time.Duration(1) * time.Second)
			b = false
			break
		case 1:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("遇到妖兽")
		case 2:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("遇到妖兽")
		case 3:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("遇到妖兽")
		case 4:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("遇到妖兽")
		default:
			fmt.Println("未知错误")
		}
		if !b {
			break
		}
		zd := zl.Xhzd()
		switch zd {
		case 0:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("正在退出")
			b = false
			break
		case 1:
			time.Sleep(time.Duration(1) * time.Second)
			zdjs()
			if js.Jzx["生命"] <= 0 {
				fmt.Println("您已经死亡")
				b = false
				break
			}
			fmt.Println("您战胜了,生命:", js.Jzx["生命"])
		case 2:
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("您逃了")
		default:
			fmt.Println("未知错误")
		}
	}
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("bye...")
	time.Sleep(time.Duration(2) * time.Second)
}

//初步结算（测试用）
func zdjs() {
	g := gs.Cjgs(1)
	//j := new(js)
	js.Jzx["生命"] -= g.Gj
	js.Jzx["经验"] += js.Jzx["智力"]

}
