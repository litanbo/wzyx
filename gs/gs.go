package gs

import "fmt"

//怪兽包下的一个操作怪兽的结构体以及方法
//怪兽生成器，怪兽的加载方式是在人物移动时或创建时加载附近半径为10的正方形范围
//目前开放地图+-500 中心为出生地，X:0,Y:0,人族在中心位置 ，神族在X:500,Y:0
//魔族在X:-500,Y:0， 精灵在0,-500，龙族在0,500，兽族在四角
//根据加载位置随机生成怪兽，属性也不同，越距离中心，生成的怪兽属性越高
type Gs struct {
	//形状Xz代表了是兽，精灵，人，龙，神，魔
	Xz string
	//类型Lx代表所属势力的种类，比如狗，狼，熊为兽族，张一，张二，李一，李二人族，
	//*.*精灵族，（金木水火土）*龙族，（天，地，人）*神族，（血煞，屠戮，魔尊）*魔族
	Lx string
	//攻击
	Gj int
	//防御
	Fy int
	//血量
	Xl int
	//智力
	Zl int
}

//用于获取一个怪兽的工厂方法
func (gs Gs) getGs(x int, y int) Gs {
	//怪兽目前根据坐标位置来生成属性，距离该类型怪物栖息地越近怪物越厉害

	g := new(Gs)
	return *g
}

//打印当前怪兽
func dy(gs Gs) {
	fmt.Println("gs:", gs)
}

//创建怪兽
func Cjgs(t int) Gs {
	g := new(Gs)
	g.Lx = "狗"
	g.Xz = "兽类"
	g.Gj = 2
	g.Xl = 30
	g.Fy = 2
	g.Zl = 0
	switch t {
	case 1:
		return *g
	default:
		fmt.Println("未知错误")
		return *g
	}
}
