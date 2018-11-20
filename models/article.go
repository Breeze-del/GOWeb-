package models

import (
	"github.com/CommonFuc/com"
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 建立数据库 sqlite3
type Category struct {
	Id    int64
	Title string
	// tag 建立表时候建索引
	Created time.Time `orm:"index"`
	// 浏览次数
	Views      int64     `orm:"index"`
	TopicTime  time.Time `orm:"index"`
	TopicCount int64
	// 最后一次操作人id
	TopicLastUserId int64
}

type Topic struct {
	Id    int64
	Uid   int64
	Title string
	// 建表 大小5000
	Content string `orm:"size(5000)"`
	// 附件
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

// 创建目录
func RegisterDB() {
	// 判断文件是否存在
	if !com.IsExist(_DB_NAME) {
		// 创建目录  path.Dir 取出目录的路径
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 使用beego中的orm中数据库驱动引擎
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 创建数据库名字 强制有一个数据库必须是default
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
