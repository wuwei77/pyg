package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

//主页展示
func (this *GoodsController) ShowIndex() {
	//二:1.获取session值
	userName := this.GetSession("pyg_userName")
	if userName != nil { //如果没有用户名,给前端传空
		this.Data["userName"] = userName.(string)
	}

	this.TplName = "index.html"
}



