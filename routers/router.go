package routers

import (
	"github.com/astaxie/beego/context"
	"pyg/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//路由过滤器,参数过滤的正则,过滤器放的位置,过滤器函数名)(在一些需要登录的操作前面加过滤器)
	beego.InsertFilter("/user/*", beego.BeforeExec,filterFunc)

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
    //退出登录
    beego.Router("/logout", &controllers.UserController{}, "get:Logout")
    //用户中心信息页
    beego.Router("/user/userCenterInfo",&controllers.UserController{}, "get:ShowUserCenterInfo")
    //用户中心地址页
	beego.Router("/user/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite;post:HandleUserCenterSite")
	//生鲜模块
	beego.Router("/indexSx",&controllers.GoodsController{}, "get:ShowIndexSx")
}
//过滤器函数--参数必须是ctx *centex.Contex
//过滤一些需要登录的界面操作
func filterFunc(ctx*context.Context)  {
	//校验session
	userName := ctx.Input.Session("pyg_userName")
	if userName ==nil{
		ctx.Redirect(302,"/login")
		return
	}

}

