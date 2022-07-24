package zl

import (
	"fmt"
	"github.com-go/wzyx/js"
)

//指令包下有对应游戏的操作指令
//该结构体主要作用使用探查指令时查看附近生命或物体，
//并生成玩家附近怪兽
var (
	//四周变量为一个长度为21*21的一个数组，里面存放着键值对key为x，y
	sz [21 * 21]map[string]string
	//存放核心坐标点,当前版本地图大小为x=+-500，y=+-500,所以根据种族核心位置生成核心坐标
	hxzb = [...]map[string]int{
		//出生地坐标位置，怪兽会被削弱 type->1代表在人族
		{
			"x":    0,
			"y":    0,
			"type": 1,
		},
		//兽族有四个坐标位置
		{
			"x": -400,
			"y": 400,
		}, {
			"x": 400,
			"y": 400,
		}, {
			"x": -400,
			"y": -400,
		}, {
			"x": 400,
			"y": -400,
		},
		//龙族坐标位置 y的轴向为北面，即龙在北方
		{
			"x": 0,
			"y": 500,
		},
		//魔族坐标位置 魔族在西方
		{
			"x": -500,
			"y": 0,
		},
		//神族坐标位置  神族在东方
		{
			"x": 500,
			"y": 0,
		},
		//精灵族坐标位置  精灵在南方
		{
			"x": 0,
			"y": -500,
		},
	}
)

//活动四周对象
func GetSz() [21 * 21]map[string]string {
	return sz

}

//选择方向时将数组整个后移或前移21位，并生成21个map
func QjOrht(stats int) {
	switch stats {
	case Wstats:
		temp := [21 * 21]map[string]string{}
		for i, _ := range sz {
			temp[i] = sz[i]
			if i >= 21 {
				sz[i] = temp[i-21]
			} else {
				sz[i] = map[string]string{"xz": "空地"}
			}
		}
	case Sstats:
		temp := [21 * 21]map[string]string{}
		for i, _ := range sz {
			temp[i] = sz[i]
			if i >= 21*20 {
				sz[i] = map[string]string{"xz": "空地"}
			} else {
				sz[i] = temp[i]
			}
		}
	case Astats:
		temp := [21 * 21]map[string]string{}
		for i, _ := range sz {
			temp[i] = sz[i]
			if i%21 == 0 {
				sz[i] = map[string]string{"xz": "空地"}
			} else {
				sz[i] = temp[i-1]
			}
		}
	case Dstats:
		for i, _ := range sz {
			if (i+1)%21 == 0 {
				sz[i] = map[string]string{"xz": "空地"}
			} else {
				sz[i] = sz[i+1]
			}
		}
	default:
		fmt.Println("QjOrht-->未知")
		return
	}

}

//获取x,y坐标生成四周的环境，并每次移动，传送时更新
func InitSz(x int, y int) {
	for index, _ := range sz {
		if index%2 == 0 {
			sz[index] = map[string]string{"xz": "兽族"}
		} else {
			sz[index] = map[string]string{"xz": "人族"}
		}

	}
}

//打印四周（格式输出）
func DySz() {
	i := 21
	for index, m := range sz {
		if (index+1)%i == 0 {
			if index == 230 {
				fmt.Print("[", js.GetY()+10-index/i, ":", js.GetY(), "]\n")
			} else {
				fmt.Print("[", m["xz"], "]", "\n")
			}
		} else if (index+1)%i == 11 {
			fmt.Print("[", js.GetX(), ":", js.GetY()+10-index/i, "]\t")
		} else {
			if index >= 210 && index < 231 {
				fmt.Print("[", js.GetX()+index%i-10, ":", js.GetY(), "]\t")
			} else {
				fmt.Print("[", m["xz"], "]\t")
			}
		}
	}
}
