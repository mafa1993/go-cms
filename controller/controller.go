package controller

var (
	views       map[string][][3]string // [a=>[[//长度3，内容为字符串],不一定几个，切面]]
	controllers map[string]interface{}
)

func init() {
	// 初始化视图和控制器
	views = make(map[string][][3]string)
	controllers = make(map[string]interface{})

	initViews()
	initController()
}

func initViews() {
	views["login_view"] = [][3]string{
		0: {"auto", "Login", "登录系统"},
		1: {"auto", "Register", "注册"},
	}

	views["index_view"] = [][3]string{
		0: {"index", "Index", "首页"},
		1: {"user", "List", "展示信息"},
	}
}

func initController() {
	controllers["index"] = &IndexController{}
	controllers["user"] = &UserController{}
	controllers["auto"] = &AutoController{}
}

func formatView(opers [][3]string) ([]string, []string) {
	var method []string = make([]string, len(opers))
	var desc []string = make([]string, len(opers))

	for k, v := range opers {
		method[k] = v[0] + "::" + v[1]
		desc[k] = v[2]
	}

	return method, desc
}
