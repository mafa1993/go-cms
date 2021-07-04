package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 统一的模型接口
type Model interface {
	ToString() string // 格式化输出数据信息
}

var (
	path   string                 = "F:\\learn\\golang\\src\\sixedu\\data\\" // 数据路径
	suffix string                 = ".sql"
	models map[string]interface{} // 记录标识 user =》 结构体
)

// "User" => User{} 不支持
func init() {
	fmt.Println("init")
	// 标识绑定注册
	models = make(map[string]interface{})
	models["user"] = NewUser
	fmt.Println(models)

	userDatas = make(map[string]Model)
	rfdata("user", "username", userDatas)
}

// 数据库文件 -> 通过配置设置
// name 数据库名称   user,admin
// pirmay 查询主键
// datas 存放数据
func rfdata(name, pirmay string, datas map[string]Model) error {
	// 1. 读取数据库文件 -》读取那个文件？
	f, ferr := os.Open(path + name + suffix) // 根据路径读取文件信息
	if ferr != nil {
		fmt.Println("文件读取异常,", ferr)
		return errors.New("文件查询不到 error")
	}
	// 关闭文件的资源流
	defer f.Close()
	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)
	// 2. 遍历数据  每一行的数据 字段根据 , 分割；数据通过 \n 分割
	field := make([]string, 0)
	for {
		row, rerr := buf.ReadBytes('\n') // 根据换行读取文件信息 , 返回的是byte[]
		if rerr != nil {
			if rerr == io.EOF { // 是否文件读取结束
				break
			}
			fmt.Println("抛出缓存读取文件异常", rerr)
		}
		// 去掉字符串，并分割数据
		data := strings.Split(string(row), ",")
		// fmt.Println("读取到的数据", data)

		// 2.1. 是否为字段
		// 根据数据判断操作 ； 是记录字段还是设置数据
		if len(field) == 0 {
			field = data // 存储字段
			for k, v := range data {
				field[k] = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
			}
		} else {
			//    2.2. 存储数据到 models
			//        2.2.1. 根据name得到模型
			//        2.2.2. 利用反射-》对模型赋值
			//        2.2.3. 再把模型存储在datas
			// 把data 的数据存放在datas
			/*
				datas := {
						"pirmay":model{data},
						"pirmay1":model{data}
						"pirmay2":model{data}
				}
			*/
			toModel(name, pirmay, datas, data, field)
		}
	}
	// fmt.Println("存储的字段", field)
	return nil
}

//    2.2. 存储数据到 models
//        2.2.1. 根据name得到模型
//        2.2.2. 利用反射-》对模型赋值
//        2.2.3. 再把模型存储在datas
func toModel(name, pirmay string, datas map[string]Model, data, field []string) error {
	// 2.2.1. 根据name得到模型
	if models[name] == nil {
		return errors.New("不存在模型 : " + name)
	}
	// 2.2.2. 利用反射-》对模型赋值 -- 如果是采用构造函数的方式则需要利用反射获取
	modelV := reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]
	// fmt.Printf("modelV type %T \n", modelV)
	// fmt.Println("data 数据", data)
	// fmt.Println("field 数据", field)

	var primayValue string // 记录当前主键的数值
	/*
		datas := {
				"pirmay":model{data},
				"pirmay1":model{data}
				"pirmay2":model{data}
		}
		field : password age sex username
		data  : 123456   18   男   shineyork
	*/
	for k, v := range data {
		if field[k] == pirmay { // 判断是否为主键字段的值 -》
			primayValue = v
			// fmt.Println("查询的主键值 : ", primayValue)
		}
		// 等到model中对应字段的set方法
		fset := modelV.MethodByName("Set" + strings.Title(field[k]))
		// fmt.Printf("fset type %T \n", fset)
		// fmt.Println("fset ", fset)

		// 调用方法
		// 你怎么知道这个值是int
		// 1、根据模型获取到属性字段
		fset.Call([]reflect.Value{
			reflect.ValueOf(toTypeValue(modelV, field[k], v)),
		})
		// fmt.Println("field[k] ", field[k])
		// mtype := modelV.Elem().FieldByName(field[k]).Type().Name()
		// fmt.Println("mtype : ", mtype)

		// reflect.Value.Elem() 得到类型
		// reflect.Zero ： &User{} => User{} 根据指针得到原有对象
		// mtype := reflect.Zero(modelV.Type().Elem()).FieldByName(field[k]).Type().Name()
		// fmt.Println("mtype : ", mtype)
	}

	// fmt.Println("model", modelV)

	datas[primayValue] = modelV.Interface().(Model)

	return nil
}

func toTypeValue(modelV reflect.Value, field, value string) interface{} {
	mtype := modelV.Elem().FieldByName(field).Type().Name()
	switch mtype {
	case "int":
		b, _ := strconv.Atoi(value)
		return b
	}
	return string(value)
}

func rwdata() {

}
