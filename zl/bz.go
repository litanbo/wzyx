package zl

import "fmt"

//指令集列表 类似帮助

var (
	//菜单指令列表
	cd = map[string]string{"W": "向前", "S": "向后", "A": "向左",
		"D": "向右", "J": "进攻", "K": "逃跑", "Q": "退出", "I": "菜单", "F": "传送", "E": "探查",
		"B": "回家", "V": "属性"}
)

//菜单的打印方法
func Dy() {
	i := 0
	for k, v := range cd {
		i++
		fmt.Printf("%v : %v ", k, v)
		if i%5 == 0 || i == len(cd) {
			fmt.Println()
		}
	}
}
