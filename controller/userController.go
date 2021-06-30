package controller

import "fmt"

type UserController struct {
}

func (c *UserController) Welcome() {
	fmt.Println("欢迎来到，六星教育系统")
	fmt.Println("你要执行的操作")
}
