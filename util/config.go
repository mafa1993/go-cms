package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	instance *Config
	conf     = flag.String("conf", "../data/config.json", "abc")
	// 获取命令行参数，以及设置默认值
)

type Config struct {
	DataPath string `json:"data_path"` //标签里第一个参数为类型，第二个参数为对应的键
	BasePath string `json:"root_path"`
}

func init() {
	flag.Parse() //分发
	//fmt.Println(*conf)
	getConfig(*conf)
}

//根据文件路径读取配置
func getConfig(name string) interface{} {
	fmt.Println("conff", *conf)
	if instance == nil {
		c := &Config{}
		bytes, err := ioutil.ReadFile(name) //读取整个文件
		fmt.Println(string(bytes))
		if err != nil {
			fmt.Println("读取失败")
			return nil
		}
		err = json.Unmarshal(bytes, c)
		if err != nil {
			fmt.Println("json 解析错误")
			return nil
		}
		instance = c
	}
	return instance
}
