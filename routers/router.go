package routers

import (
	"pyg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //注册
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //发送短信
    beego.Router("/sendMsg", &controllers.UserController{}, "post:SendMsg")
    //激活业务
    beego.Router("/active", &controllers.UserController{}, "get:ShowActive;post:HandleActive")
    //激活用户
    beego.Router("/activeUser", &controllers.UserController{}, "get:ActiveUser")
    //登录
    beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
    //主页
    beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")

}
