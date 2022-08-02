package js

import "fmt"

//角色相关信息

var (
	name = "姜子星"
	x    = 0
	y    = 0
	//玩家数据列表
	Jzx = map[string]int{"体质": 8, "力量": 6, "智力": 3, "等级": 1, "经验": 0,
		"攻击": 12, "防御": 0, "生命": 64, "X": 0, "Y": 0}
)

func GetY() int {
	return y
}
func SetY(i int) {
	Jzx["Y"] = i
	y = i
}
func GetX() int {
	return x

}
func SetX(i int) {
	Jzx["X"] = i
	x = i
}
func GetName() string {
	return name

}
func SetName(username string) {
	name = username
}

//打印玩家数据
func Dy() {
	i := 0
	for k, v := range Jzx {
		i++
		fmt.Printf("%v : %v ", k, v)
		if i%5 == 0 || i == len(Jzx) {
			fmt.Println()
		}
	}
}
