package controller

import (
	"encoding/json"
	"fmt"
	"testing"
)

// type myFun func()

// func (f myFun) call() {
// 	fmt.Println("=== >>> start ===>>> ")
// 	f() // ?????????????????
// 	fmt.Println("=== >>> end ===>>> ")
// }

// /*
// [
//   "p" => 'k'
// ]

// route := {
//   "index" : {
//     "oper" : [
//         0:"login::登入",
//         1:"index::首页"
//     ],
//     "controller" : &controller{}
//   }
// }
// */

// func TestFu(t *testing.T) {
// 	// ffff := sum
// 	// map
// 	// fmt.Println(ffff())
// 	newSumcoo := myFun(dispatch) // string() int() int64()
// 	newSumcoo.call()
// }

// func dispatch() {
// 	fmt.Println("dispath lll")
// }

// //
// func sum() int {
// 	return 9
// }

type Config struct {
	DataPath string `json:"data_path"` //标签里第一个参数为类型，第二个参数为对应的键
	BasePath string `json:"root_path"`
}

func TestJson(t *testing.T) {
	c := Config{"datapath", "basepath"}
	c1, _ := json.Marshal(c)

	fmt.Println(string(c1))

	json_str := `
	{
		"data_path":"patha",
		"root_path":"pathb"
	}`
	tt := []byte(json_str)
	var cc Config
	json.Unmarshal(tt, &cc)
	fmt.Printf("%s %T", cc, cc)
}
