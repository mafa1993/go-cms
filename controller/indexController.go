package controller

import "fmt"

type IndexController struct {
}

func (c *IndexController) Welcome() {
	fmt.Println("欢迎来到，此管理系统")
	fmt.Println("你要执行的操作")
}

func (c *IndexController) Index() {
	fmt.Println("进入首页")
}
