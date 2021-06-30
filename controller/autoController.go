package controller

import (
	"fmt"
)

type AutoController struct {
}

func (c *AutoController) Login() {
	fmt.Println("登录")
}

func (c *AutoController) Register() {
	fmt.Println("注册用户")
}
