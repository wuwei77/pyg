package controllers

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

//主页展示
func (this *GoodsController) ShowIndex() {
	this.TplName="index.html"
}
