package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader

type Cfun func()

//定义一个函数类型，让这个函数在另一个函数里执行，类似于依赖注入
func CReturn(c Cfun) {
	fmt.Println("========开始=========")

	c()
	fmt.Println("========结束=========")
}

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

func CRead() string {
	input, _ := inputReader.ReadString('\n')
	return strings.TrimSpace(strings.TrimSuffix(input, "\n"))
}

//输出内容
func Comper(operate []string) {
	for key, value := range operate {
		fmt.Printf("{%d}:%s \n", key, value)
	}

	fmt.Println("退出请输入x")
}
