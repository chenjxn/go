package One

import (
	"fmt"
	"os"
)

// os.Args 获取命令行运行的传入的参数
func Command() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
}
