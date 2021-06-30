package sixedu_test

import (
	"fmt"
	"testing"
)

type myFun func()

func (f myFun) call() {
	fmt.Println("=== >>> start ===>>> ")
	f() // ?????????????????
	fmt.Println("=== >>> end ===>>> ")
}

/*
[
  "p" => 'k'
]

route := {
  "index" : {
    "oper" : [
        0:"login::登入",
        1:"index::首页"
    ],
    "controller" : &controller{}
  }
}
*/

func TestFu(t *testing.T) {
	// ffff := sum
	// map
	// fmt.Println(ffff())
	newSumcoo := myFun(dispatch) // string() int() int64()
	newSumcoo.call()
}

func dispatch() {
	fmt.Println("dispath lll")
}

//
func sum() int {
	return 9
}
