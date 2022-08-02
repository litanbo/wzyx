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
	fmt.Println("游戏加载中。。。")
	var name = ""
	//TODO
	//目前还没有做数据持久化，暂时无存档，默认出生点并创建四周视角
	fmt.Println("请输入玩家名：")
	fmt.Scan(&name)
	js.SetName(name)
	zl.InitSz(js.GetX(), js.GetY())
}
func main() {
	fmt.Println("请选择：1，开始新游戏。2，读取存档")
	var stauts int
	fmt.Scan(&stauts)
	if stauts == 1 {
		fmt.Println("创建角色,玩家：", js.GetName(), "加载完成")
		start()
	} else if stauts == 2 {
		fmt.Println("无存档")
	} else {
		fmt.Println("退出")
		return
	}
}

func start() {
	fmt.Println("正在加载中。。。")
	//time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("W,向前,S,向后，A,向左,D,向右,J,进攻,K,逃跑,Q,退出,I,菜单")
	//time.Sleep(time.Duration(1) * time.Second)
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
			sz := zl.GetSz()
			xz := sz[zl.Wdwz]
			fmt.Println("遇到", xz["xz"])
		case 2:
			time.Sleep(time.Duration(1) * time.Second)
			sz := zl.GetSz()
			xz := sz[zl.Wdwz]
			fmt.Println("遇到", xz["xz"])
		case 3:
			time.Sleep(time.Duration(1) * time.Second)
			sz := zl.GetSz()
			xz := sz[zl.Wdwz]
			fmt.Println("遇到", xz["xz"])
		case 4:
			time.Sleep(time.Duration(1) * time.Second)
			sz := zl.GetSz()
			xz := sz[zl.Wdwz]
			fmt.Println("遇到", xz["xz"])
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
