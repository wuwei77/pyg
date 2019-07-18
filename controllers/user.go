package controllers

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"pyg/models"
	"regexp"
	"time"
)

type UserController struct {
	beego.Controller
}

//错误处理返回,参数2-错误信息,参数3-传递到哪个页面
func ErrResp(this *UserController,errmsg string,fileName string)  {
	this.Data["errmsg"] = errmsg
	this.TplName = fileName
}
//显示注册页面
func (this *UserController) ShowRegister() {
	this.TplName="register.html"
}

//处理注册页面
func (this *UserController) HandleRegister() {
	//获取数据
	phone := this.GetString("phone")
	code := this.GetString("code")
	pwd := this.GetString("password")
	rpwd := this.GetString("repassword")
	//校验数据
	if phone == "" || code == "" || pwd == "" || rpwd == ""{
		this.Data["errmsg"] = "输入数据不能为空"
		this.TplName = "register.html"
		return
	}

	if pwd != rpwd {
		this.Data["errmsg"] = "两次密码输入不一致"
		this.TplName = "register.html"
		return
	}

	//验证码校验   1.首先要短信发送

	//从redis取出验证码
	conn,err := redis.Dial("tcp", "192.168.31.21:6379")
	if err!=nil{
		ErrResp(this, "redis连接失败", "register.html")
		return
	}
	//从redis中获取数据
	result, err:=redis.String(conn.Do("get", phone+"_code"))
	if result != code{
		ErrResp(this, "验证码错误", "register.html")
		return
	}

	//3.处理数据;把数据存储到数据库中
	o := orm.NewOrm()
	var user models.User

	user.Name = phone
	user.PassWord = pwd
	//插入
	o.Insert(&user)

	//4.返回数据,返回到激活页面
	this.Redirect("/active", 302)

}

//短信发送
func (this *UserController) ShowSendMsg() {
	//接收电话号码
	phone := this.GetString("phone")
	if phone !=""{
		ErrResp(this, "电话号码不能为空", "register.html")
		return
	}
	//电话格式检验
	reg, _:=regexp.Compile(`^1[3-9][0-9]{9}$`)
	//找到符合的返回找到的字符串,没有找到为空
	result:=reg.FindString(phone)
	if result == ""{
		ErrResp(this, "电话号码不能为空", "register.html")
		return
	}

	//随机生成6位验证码
	//添加随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成6位数随机数
	//fmt.Printf("", )打印到控制台,fmt.Sprintf,打印到返回值,凑成字符串
	//%06d  按六位输出,前面不足补0
	vscode :=fmt.Sprintf("%06d",rand.Int31n(1000000))


	//后台要拿出验证码和输入的作比较,这里验证要存到redis中去
	//另一种方法.发送到前端,用隐藏域传值

	//实现把验证码写入redis中去
	conn, err :=redis.Dial("tcp", "192.168.31.21:6379")
	if err!=nil{
		//这里要返回错误给前端,这里返回的是一个json
		resp := make(map[string]interface{})
		resp["statusCode"] = 401  //自己定义
		resp["msg"] = "redis连接失败"
		this.Data["json"] = resp
		this.ServeJSON()
	}
	//存入redis
	//设置唯一性,key值,value值vscode;设置验证码超时时间
	conn.Do("setex", phone+"_code",60*5,vscode)


	//阿里短信发送
	//1.初始化客户端
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "LTAI49yQmf3Tbhdi", "dDNrUp9tKQK4kOORDXMNIkWV23dl4R")
	if err != nil {
		panic(err)
	}
	//基本配置
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	//定义内容
	//发送给谁的电话号码
	request.QueryParams["PhoneNumbers"] = phone
	//签名
	request.QueryParams["SignName"] = "品优购"
	//模板的code
	request.QueryParams["TemplateCode"] = "SMS_164275022"
	//验证码
	request.QueryParams["TemplateParam"] = `{"code":`+vscode+`}`

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())


	//回复ajax给前端
	//beego发送json数据
	//1.第一步要有一个容器
	//go中key:value容器--map和struct
	//使用map做容器
	resp := make(map[string]interface{})
	//2.给容器赋值
	//返回的状态吗
	resp["statusCode"] =200
	//返回的容器
	resp["msg"] = "OK"
	//本来是要序列化的,beego有一个方法可以不序列化
	//3.指定返回方式-json
	this.Data["json"] = resp
	//4.返回数据--beego发送json给前端的方法
	this.ServeJSON()



}
