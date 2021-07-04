package controller

import "fmt"

type IndexController struct {
}

func (c *IndexController) Welcome() {
	fmt.Println("欢迎来到，此管理系统")
	fmt.Println("你要执行的操作")
	view = "login_view" // 这样在route的view中会使用这个view，然后展示login_view的视图
}

func (c *IndexController) Index() {
	fmt.Println("进入首页")
	view = "index_view"
}
