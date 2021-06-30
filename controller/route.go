package controller

import (
	"fmt"
	"reflect"
	"sixedu/util"
	"strconv"
	"strings"
)

var (
	autoController *AutoController
	next           string
	view           string
)

func init() {
	autoController = &AutoController{}
}

// 可以实现中间件
func Run() {
	fmt.Println("123")
	next = "index::Welcome"
	for {
		dispatch()
	}

}
func dispatch() {
	// fmt.Println("======= >>> start ======>>> ")
	// fmt.Println("欢迎来到，六星教育系统")
	// fmt.Println("你要执行的操作")
	// oper := []string{"登入", "注册"}
	// for k, v := range oper {
	// 	fmt.Printf("(%d) : %s\n", k, v)
	// }
	// // 查看用户信息
	// // 修改用户

	// // 进入修改用户界面 -》 用户的修改属性是那一项；还是说返回
	// flag, _ := strconv.Atoi(util.CRead())
	// fmt.Println("======= >>> end ======>>> ")

	// fmt.Println("======= >>> start ======>>> ")

	// // if 是否为登入 =>
	// switch flag {
	// case 0:
	// 	al := autoController.login()
	// 	if al {
	// 		sessionid = true
	// 	}
	// case 1:
	// }
	// fmt.Println("======= >>> end ======>>> ")
	// fmt.Println("======= >>> start ======>>> ")
	// if sessionid { // 登入成功
	// 	oper = []string{"用户信息展示", "修改用户", "添加用户"}
	// 	for k, v := range oper {
	// 		fmt.Printf("(%d) : %s\n", k, v)
	// 	}
	// 	flag, _ = strconv.Atoi(util.CRead())
	// 	switch flag {
	// 	case 0:
	// 		// 用户信息展示

	// 		// ??
	// 		oper = []string{"返回", "修改用户", "添加用户"}
	// 		fmt.Println("======= >>> end ======>>> ")
	// 	case 1:

	// 	}

	// 1. 根据指定的控制器方法进行执行
	args := strings.Split(next, "::") //0为控制器 1为方法
	controller := controllers[args[0]]

	//反射执行,
	cr := reflect.ValueOf(controller) // valueof 返回value结构体
	c := cr.MethodByName(args[1])     // 根据
	c.Call([]reflect.Value{})         //call 调用

	// 获取下一步操作
	opers, ok := views[view]
	if !ok {
		fmt.Println("没有获取到视图", view)
	}

	// 数据处理
	method, desc := formatView(opers)
	// 输出下一步操作
	util.Comper(desc)

	for {
		input := util.CRead()
		if input == "x" {
			fmt.Println("退出")
		}

		flag, err := strconv.Atoi(input)
		// string 转换成 int 相当于parseInt(s,10,0)

		if err == nil && flag < len(method) {
			next = method[flag]
			break
		}

		fmt.Println("输入错误")
	}
}
