package zl

//存放行动相关指令
import (
	"fmt"
	"github.com-go/wzyx/js"
	"strings"
	"time"
)

const (
	//前进W方向状态代表
	Wstats = 1
	//后退S方向状态代表
	Sstats = 2
	//左转A方向状态代表
	Astats = 3
	//右转D方向状态代表
	Dstats = 4
	//Q退出代表
	Qstats = 0
	//失败代表
	Notstats = -1
	//战斗  进攻j状态代表
	Jstats = 1
	//战斗 逃跑k状态代表
	Kstats = 2
)

//战斗指令
func Xhzd() int {
ZD:
	fmt.Println("请输入：1.交流，2.提问，3.收服，4.查看，J.进攻，K.逃跑")
	var str string
	fmt.Scan(&str)
	i := zd(str)
	if i < 0 {
		goto ZD
	}
	return i
}

//方向指令
func Fx() int {
XD:
	fmt.Println("请操作，或I查看指令菜单")
	var str string
	fmt.Scan(&str)
	i := xd(str)
	if i < 0 {
		goto XD
	}
	return i
}

//战斗方法
func zd(str string) int {
	upper := strings.ToUpper(str)
	switch upper {
	case "Q":
		return Qstats
	case "J":
		return Jstats
	case "K":
		return Kstats
	case "1":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("进攻进攻")
		return Notstats
	case "2":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("没啥说的快攻击")
		return Notstats
	case "3":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("您当前还不能收服")
		return Notstats
	case "4":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("兽类妖兽，狗，血量：20，攻击：2，防御：2，")
		return Notstats
	case "I":
		fmt.Print("菜单：\n")
		Dy()
		return Notstats
	case "V":
		fmt.Println("玩家：", js.GetName())
		js.Dy()
		return Notstats
	case "E":
		DySz()
		return Notstats
	default:
		fmt.Println("请正确输入战斗指令（1，2，3，4,J,K）。")
		return Notstats
	}
}

//行动方法
func xd(str string) int {
	upper := strings.ToUpper(str)
	switch upper {
	case "Q":
		fmt.Println("退出。")
		return Qstats
	case "W":
		fmt.Println("您向前走了一步。")
		js.SetY(js.GetY() + 1)
		QjOrht(Wstats, js.GetX(), js.GetY())
		return Wstats
	case "S":
		fmt.Println("您向后退了一步。")
		js.SetY(js.GetY() - 1)
		QjOrht(Sstats, js.GetX(), js.GetY())
		return Sstats
	case "A":
		fmt.Println("您向前左转一步。")
		js.SetX(js.GetX() - 1)
		QjOrht(Astats, js.GetX(), js.GetY())
		return Astats
	case "D":
		fmt.Println("您向前右转一步。")
		js.SetX(js.GetX() + 1)
		QjOrht(Dstats, js.GetX(), js.GetY())
		return Dstats
	case "I":
		fmt.Print("菜单：\n")
		Dy()
		return Notstats
	case "V":
		fmt.Println("玩家：", js.GetName())
		js.Dy()
		return Notstats
	case "B":
		fmt.Println("暂未开发")
		return Notstats
	case "F":
		fmt.Println("暂未开发")
		return Notstats
	case "E":
		DySz()
		return Notstats
	default:
		fmt.Println("请正确输入方向指令（WSAD）。")
		return Notstats
	}
}
