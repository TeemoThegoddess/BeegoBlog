package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Category        string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id       int64
	Tid      int64
	Nickname string
	Content  string    `orm:"size(5000)"`
	Created  time.Time `orm:"index"`
}

func init() {
	//从配置文件获取数据库相关信息
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dbinfo := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8&loc=Local"
	//注册model
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	//注册驱动，对于mysql、sqlite3、dp可以不用注册
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", dbinfo, 10)
	//orm.RegisterDataBase("default", "mysql", "root:xzzhz515218638@/blog?charset=utf8", 10)
	//打开调试
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//建表
	orm.RunSyncdb("default", false, true)
}
