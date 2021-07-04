package controller

// import (
// 	"errors"
// 	"fmt"
// 	"reflect"
// 	"sixedu/util"
// 	"strconv"
// 	"strings"
// )

// var (
// 	autoController *AutoController
// 	next           string
// 	view           string
// )

// func init() {
// 	autoController = &AutoController{}
// }

// // 可以实现中间件
// func Run() {

// 	// sixedu.GetConfig()

// 	next = "index::Welcome"
// 	fmt.Println(util.GetConfig())
// 	for {
// 		flag := util.CReturn(util.Cfunc(dispatch))
// 		if flag {
// 			break
// 		}
// 	}
// 	fmt.Println("结束")
// }
// func dispatch() (bool, error) {

// 	// 1. 根据指定的控制器和方法执行
// 	args := strings.Split(next, "::")
// 	controller, ok := controllers[args[0]]
// 	if ok != true {
// 		// fmt.Println("读取不到控制器", args[0])
// 		return false, errors.New("读取不到控制器" + args[0])
// 	}
// 	// 反射执行方法
// 	// 1. 先传递执行的控制器对象 controlller
// 	// 2. 根据方法名，执行方法
// 	// fmt.Println("args[0]", args[0])
// 	// fmt.Println("args[1]", args[1])
// 	cr := reflect.ValueOf(controller) // reflect.Value
// 	cr.MethodByName(args[1]).Call([]reflect.Value{})

// 	// 2. 获取下一步执行的操作
// 	// opers = [][3]string{
// 	// 	0: {0: "auto", 1: "login", 2: "登入系统"},
// 	// 	1: {0: "auto", 1: "register", 2: "注册用户"},
// 	// }
// 	opers, ok := views[view]
// 	if ok != true {
// 		// return fal
// 		// fmt.Println("获取不到视图 ；", view)
// 		return false, errors.New("获取不到视图" + view)
// 	}

// 	// 3. 数据处理
// 	// methods{
// 	// 	0: "auto::login",
// 	// 	1:"auto::register",
// 	// }
// 	// desc:{
// 	// 	0:"登入系统",
// 	// 	1:"注册用户"
// 	// }
// 	methods, desc := toModelFormate(opers)
// 	// 输出下一步操作
// 	util.Coper(desc)

// 	// 4. 用户的界面展示
// 	for {
// 		input := util.CRead()
// 		if input == "x" {
// 			// fmt.Println("退出")
// 			return true, nil
// 		}
// 		flag, err := strconv.Atoi(input)
// 		if err == nil && flag < len(methods) {
// 			next = methods[flag]
// 			break
// 		}
// 		fmt.Println("输入有误重新输入")
// 	}
// 	return false, nil
// }
