package gs

import (
	"fmt"
	"math/rand"
	"time"
)

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

//游戏世界的基本规则

const (
	//智力等级 数越大智力越高
	zldj1, zldj2, zldj3, zldj4, zldj5 = 1, 2, 3, 4, 5
	// 郊外核心,人族核心,兽族核心,龙族核心,精灵核心,神族核心
	jwhx, rzhx1, shouzhx2, lzhx3, jlhx4, mzhx5, shenzhx6 = 0, 1, 2, 3, 4, 5, 6
)

//用于获取一个怪兽的工厂方法
func GetGs(x int, y int) Gs {
	ishx := ishx(x, y)
	gj, xl, fy := sxdz(x + y)
	g := new(Gs)
	g.Fy = fy
	g.Gj = gj
	g.Xl = xl
	g.Zl = 1
	if ishx == jwhx {
		g.Xz = "兽族"
		g.Lx = "无"
	}
	return *g
}

//根据坐标来获取是否需要创建人物的概率
func IsNewGs(x int, y int) int {
	ishx := ishx(x, y)
	sum := x + y
	rand.Seed(time.Now().UnixNano())
	switch ishx {
	case jwhx:
		if sum <= 100 && sum >= -100 {
			//如果在正负50以内的人物概率 人族概率百分之40 兽族概率百分之29
			intn := rand.Intn(10000)
			if intn <= 4000 {
				return 1
			} else if intn >= 7100 {
				return 2
			} else {
				return -1
			}
		} else if sum <= 200 && sum >= -200 {
			//如果在正负50以外100以内 人族概率百分之35 兽族百分之35
			intn := rand.Intn(10000)
			if intn <= 3500 {
				return 1
			} else if intn >= 6500 {
				return 2
			} else {
				return -1
			}
		} else if x+y <= 300 {
			//如果在正负100以外，150以内人族概率百分之30 兽族 百分之40
			intn := rand.Intn(10000)
			if intn <= 3000 {
				return 1
			} else if intn >= 6000 {
				return 2
			} else {
				return -1
			}
		} else if x+y <= 400 {
			//如果在正负150以外，200以内 人族百分之25  兽族百分之45
			intn := rand.Intn(10000)
			if intn <= 2500 {
				return 1
			} else if intn >= 5500 {
				return 2
			} else {
				return -1
			}
		} else {
			//如果在200以外 人族百分之 10 兽族 百分之 60
			intn := rand.Intn(10000)
			if intn <= 1000 {
				return 1
			} else if intn >= 4000 {
				return 2
			} else {
				return -1
			}
		}
	case rzhx1:
		//人族核心  70 概率遇到
		intn := rand.Intn(10000)
		if intn <= 7000 {
			return 1
		} else {
			return -1
		}
	case shouzhx2:
		//兽族核心  90概率遇到
		intn := rand.Intn(10000)
		if intn <= 9000 {
			return 2
		} else {
			return -1
		}
	case lzhx3:
		//龙族核心 20概率遇到
		intn := rand.Intn(10000)
		if intn <= 2000 {
			return 3
		} else {
			return -1
		}
	case jlhx4:
		//精灵核心 40概率遇到
		intn := rand.Intn(10000)
		if intn <= 4000 {
			return 4
		} else {
			return -1
		}
	case mzhx5:
		//魔族核心 70概率遇到
		intn := rand.Intn(10000)
		if intn <= 7000 {
			return 5
		} else {
			return -1
		}
	case shenzhx6:
		//神族核心 70概率遇到
		intn := rand.Intn(10000)
		if intn <= 7000 {
			return 6
		} else {
			return -1
		}
	default:
		return -1
	}
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

//属性点值，
func sxdz(sumxy int) (gj int, xl int, fy int) {
	rand.Seed(time.Now().UnixNano())
	//rand.Int()
	if sumxy <= 100 {
		return rand.Intn(8) + 3, rand.Intn(3), rand.Intn(300) + 80
	} else {
		return rand.Intn(8) + 3, rand.Intn(3), rand.Intn(300) + 80
	}

}

//判断是否为核心
func ishx(x int, y int) int {
	if x >= 450 && y >= 300 {
		return jwhx
	}
	if x >= 450 && y >= -300 {
		return shenzhx6
	}
	if x >= 450 && y >= -500 {
		return jwhx
	}
	if x >= 350 && y >= 450 {
		return jwhx
	}
	if x >= 350 && y >= 350 {
		return shouzhx2
	}
	if x >= 350 && y >= 250 {
		return jwhx
	}
	if x >= 350 && y >= 150 {
		return shouzhx2
	}
	if x >= 350 && y >= 50 {
		return jwhx
	}
	if x >= 350 && y >= -50 {
		return rzhx1
	}
	if x >= 350 && y >= -150 {
		return jwhx
	}
	if x >= 350 && y >= -250 {
		return shouzhx2
	}
	if x >= 350 && y >= -350 {
		return jwhx
	}
	if x >= 350 && y >= -450 {
		return shouzhx2
	}
	if x >= 350 && y >= -500 {
		return jwhx
	}
	if x >= 250 && y >= 450 {
		return lzhx3
	}
	if x >= 250 && y >= -450 {
		return jwhx
	}
	if x >= 250 && y >= -500 {
		return jlhx4
	}
	if x >= 150 && y >= 450 {
		return lzhx3
	}
	if x >= 150 && y >= 350 {
		return jwhx
	}
	if x >= 150 && y >= 250 {
		return shouzhx2
	}
	if x >= 150 && y >= 150 {
		return jwhx
	}
	if x >= 150 && y >= 50 {
		return rzhx1
	}
	if x >= 150 && y >= -50 {
		return jwhx
	}
	if x >= 150 && y >= -150 {
		return rzhx1
	}
	if x >= 150 && y >= -250 {
		return jwhx
	}
	if x >= 150 && y >= -350 {
		return shouzhx2
	}
	if x >= 150 && y >= -450 {
		return jwhx
	}
	if x >= 150 && y >= -500 {
		return jlhx4
	}
	if x >= 50 && y >= 450 {
		return lzhx3
	}
	if x >= 50 && y >= 350 {
		return jwhx
	}
	if x >= 50 && y >= 250 {
		return shouzhx2
	}
	if x >= 50 && y >= 150 {
		return jwhx
	}
	if x >= 50 && y >= 50 {
		return rzhx1
	}
	if x >= 50 && y >= -50 {
		return jwhx
	}
	if x >= 50 && y >= -150 {
		return rzhx1
	}
	if x >= 50 && y >= -250 {
		return jwhx
	}
	if x >= 50 && y >= -350 {
		return shouzhx2
	}
	if x >= 50 && y >= -450 {
		return jwhx
	}
	if x >= 50 && y >= -500 {
		return jlhx4
	}
	if x >= -50 && y >= 450 {
		return lzhx3
	}
	if x >= -50 && y >= 350 {
		return jwhx
	}
	if x >= -50 && y >= 250 {
		return shouzhx2
	}
	if x >= -50 && y >= 150 {
		return jwhx
	}
	if x >= -50 && y >= 50 {
		return rzhx1
	}
	if x >= -50 && y >= -50 {
		return jwhx
	}
	if x >= -50 && y >= -150 {
		return rzhx1
	}
	if x >= -50 && y >= -250 {
		return jwhx
	}
	if x >= -50 && y >= -350 {
		return shouzhx2
	}
	if x >= -50 && y >= -450 {
		return jwhx
	}
	if x >= -50 && y >= -500 {
		return jlhx4
	}
	if x >= -150 && y >= 450 {
		return lzhx3
	}
	if x >= -150 && y >= 350 {
		return jwhx
	}
	if x >= -150 && y >= 250 {
		return shouzhx2
	}
	if x >= -150 && y >= 150 {
		return jwhx
	}
	if x >= -150 && y >= 50 {
		return rzhx1
	}
	if x >= -150 && y >= -50 {
		return jwhx
	}
	if x >= -150 && y >= -150 {
		return rzhx1
	}
	if x >= -150 && y >= -250 {
		return jwhx
	}
	if x >= -150 && y >= -350 {
		return shouzhx2
	}
	if x >= -150 && y >= -450 {
		return jwhx
	}
	if x >= -150 && y >= -500 {
		return jlhx4
	}
	if x >= -250 && y >= 450 {
		return lzhx3
	}
	if x >= -250 && y >= 350 {
		return jwhx
	}
	if x >= -250 && y >= 250 {
		return shouzhx2
	}
	if x >= -250 && y >= 150 {
		return jwhx
	}
	if x >= -250 && y >= 50 {
		return rzhx1
	}
	if x >= -250 && y >= -50 {
		return jwhx
	}
	if x >= -250 && y >= -150 {
		return rzhx1
	}
	if x >= -250 && y >= -250 {
		return jwhx
	}
	if x >= -250 && y >= -350 {
		return shouzhx2
	}
	if x >= -250 && y >= -450 {
		return jwhx
	}
	if x >= -250 && y >= -500 {
		return jlhx4
	}
	if x >= -350 && y >= 450 {
		return lzhx3
	}
	if x >= -350 && y >= -450 {
		return jwhx
	}
	if x >= -350 && y >= -500 {
		return jlhx4
	}
	if x >= -450 && y >= 450 {
		return jwhx
	}
	if x >= -450 && y >= 350 {
		return shouzhx2
	}
	if x >= -450 && y >= 250 {
		return jwhx
	}
	if x >= -450 && y >= 150 {
		return shouzhx2
	}
	if x >= -450 && y >= 50 {
		return jwhx
	}
	if x >= -450 && y >= -50 {
		return rzhx1
	}
	if x >= -450 && y >= -150 {
		return jwhx
	}
	if x >= -450 && y >= -250 {
		return shouzhx2
	}
	if x >= -450 && y >= -350 {
		return jwhx
	}
	if x >= -450 && y >= -450 {
		return shouzhx2
	}
	if x >= -450 && y >= -500 {
		return jwhx
	}
	if x >= -500 && y >= 300 {
		return jwhx
	}
	if x >= -500 && y >= -300 {
		return mzhx5
	}
	if x >= -500 && y >= -500 {
		return jwhx
	}
	return -1
}
