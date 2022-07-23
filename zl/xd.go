package zl

//存放行动相关指令
import (
	"fmt"
	"github.com-go/wzyx/js"
	"strings"
	"time"
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
		return 0
	case "J":
		return 1
	case "K":
		return 2
	case "1":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("进攻进攻")
		return -1
	case "2":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("没啥说的快攻击")
		return -1
	case "3":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("您当前还不能收服")
		return -1
	case "4":
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("兽类妖兽，狗，血量：20，攻击：2，防御：2，")
		return -1
	case "I":
		fmt.Print("菜单：\n")
		Dy()
		return -1
	case "V":
		fmt.Println("玩家：", js.Name)
		js.Dy()
		return -1
	default:
		fmt.Println("请正确输入战斗指令（1，2，3，4,J,K）。")
		return -1
	}
}

//行动方法
func xd(str string) int {
	upper := strings.ToUpper(str)
	switch upper {
	case "Q":
		fmt.Println("退出。")
		return 0
	case "W":
		fmt.Println("您向前走了一步。")
		return 1
	case "S":
		fmt.Println("您向后退了一步。")
		return 2
	case "A":
		fmt.Println("您向前左转一步。")
		return 3
	case "D":
		fmt.Println("您向前右转一步。")
		return 4
	case "I":
		fmt.Print("菜单：\n")
		Dy()
		return -1
	case "V":
		fmt.Println("玩家：", js.Name)
		js.Dy()
		return -1
	case "B":
		fmt.Println("暂未开发")
		return -1
	case "F":
		fmt.Println("暂未开发")
		return -1
	case "E":
		fmt.Println("暂未开发")
		return -1
	default:
		fmt.Println("请正确输入方向指令（WSAD）。")
		return -1
	}
}
