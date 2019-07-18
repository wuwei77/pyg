package routers

import (
	"pyg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister")
    //发送短信
    beego.Router("/sendMsg", &controllers.UserController{}, "get:ShowSendMsg")
}
