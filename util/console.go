package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader

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
