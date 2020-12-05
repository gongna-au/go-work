package main

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "magic.chen"
	c.Data["Email"] = "cfqcsunng@gmail.com"
	c.TplName = "index.tpl"
}

// 基本的GET请求获取用户名
type UserController struct {
	beego.Controller
}

type Date struct {
	name string
	age  int
	sex  string
}

func (c *UserController) Get() {
	data := Date{"magic", 24, "男"}
	content := strings.Join([]string{"姓名:" + data.name, "性别:" + data.sex, "年龄:" + strconv.Itoa(data.age)}, " ")
	c.Ctx.WriteString(content)
}

// 基本的Post请求
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	content := strings.Join([]string{"用户名:" + username, "密码:" + password}, " ")
	c.Ctx.WriteString(content)
}

func main() {
	func (this *MainController) Get() {
		this.Ctx.WriteString("hello")
		func (this *UserController) Get() {
			this.Ctx.WriteString("hello")
		}
		func (this *Date()) Get() {
			this.Ctx.WriteString("hello")
		}
	}
}