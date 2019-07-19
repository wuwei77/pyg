package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pyg/models"
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

	//不同类型组合存储需要通过interface{}存储,我们有需要标识-标识有两种方式map和切片
	//这里我们使用map来标识.这样一个一级标题key其value就是存储二级标题就是一个map[string]interface{}
	//这样多个一级标题就是map切片  []map[string]interface{}

	//定义一个大容器
	var goodsTypes []map[string]interface{}

	//获取一级菜单
	o := orm.NewOrm()
	//获取一级标题对象
	var tpshops []models.TpshopCategory
	//通过Pid来查询所有一级标题,因为一级标题父Pid为0固定==得到所有一级标题
	o.QueryTable("TpshopCategory").Filter("Pid", 0).All(&tpshops)

	//根据一级菜单获取二级菜单
	//先把每一个一级切片取出来--通过循环
	for _, yiji := range tpshops{
		//定义一个一级容器来存储一个一级标题和其对应的二级标题切片,是用map来存储
		tempContainer :=make(map[string]interface{})
		//获取二级标题对象
		var erji []models.TpshopCategory
		//查询二级标题,通过一级标题来获取二级标题的Pid,来查询出所有二级标题
		o.QueryTable("TpshopCategory").Filter("Pid", yiji.Id).All(&erji)
		//把一级和二级标题都存入这个一级map容器中
		tempContainer["yiji"] = yiji
		tempContainer["erji"] = erji
		//把这个map加入到大容器,map切片中存储
		goodsTypes = append(goodsTypes, tempContainer)

	}
	//获取三级菜单--通过二级菜单的Pid来查询
	//这个v存储了(一级标题和其对应的二级标题)的map,就是一级容器
	for _, v :=range goodsTypes{
		//获取二级容器,用来存储二级标题和其对应的三级标题
		var erjiContainer []map[string]interface{}
		//循环二级标题的切片--把每一个二级标题取出来
		for _, erClass :=range v["erji"].([]models.TpshopCategory){
			//获取一个三级标题对象
			var sanji []models.TpshopCategory
			//定义一个三级容器--存储二级标题和三级标题切片
			tempContainer :=make(map[string]interface{})
			//通过二级标题的Pid来查询其对应的三级标题
			o.QueryTable("TpshopCategory").Filter("Pid", erClass.Id).All(&sanji)
			//把二级和三级标题都存入这个三级map容器中
			tempContainer["erji"] = erClass
			tempContainer["sanji"] = sanji

			//把这个三级map容器放到二级map切片容器中
			erjiContainer = append(erjiContainer,tempContainer)
		}
		//把二级容器map切片放到一级容器中
		v["sanji"] = erjiContainer
	}
	this.Data["goodsTypes"]=goodsTypes
	this.TplName = "index.html"
}

//生鲜模块首页展示
func (this *GoodsController) ShowIndexSx()  {
	this.TplName="index_sx.html"

}



