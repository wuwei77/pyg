package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户表
type User struct {
	Id       int
	Name     string     `orm:"size(20);unique"` //用户名
	PassWord string     `orm:"size(20)"`        //登陆密码
	Email    string     `orm:"size(50)"`        //邮箱
	Active   bool       `orm:"default(false)"`  //是否激活
	Power    int        `orm:"default(0)"`      //权限设置  0 表示未激活  1表示激活
	Address  []*Address `orm:"reverse(many)"`
}

//地址表
type Address struct {
	Id        int
	Receiver  string `orm:"size(20)"`       //收件人
	Addr      string `orm:"size(50)"`       //收件地址
	Zipcode   string `orm:"size(20)"`       //邮编
	Phone     string `orm:"size(20)"`       //联系方式
	Isdefault bool   `orm:"default(false)"` //是否默认 0 为非默认  1为默认
	User      *User  `orm:"rel(fk)"`        //用户ID
}

func init() {
	//注册数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.31.47:3306)/pyg")

	//注册表
	orm.RegisterModel(new(User), new(Address))

	//运行
	orm.RunSyncdb("default", false, true)
}
