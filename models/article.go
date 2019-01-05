package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

const (
	_SQLITE3_DRIVER = "mysql"
)

// 建立数据库 mysql
// orm方法 insert增 delete删 update改 查read 查多个queryTable（）Filter（）All（）
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
	Id       int64
	Uid      int64
	Title    string
	Category string
	Lables   string
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

// 创建表
func RegisterDB() {
	//// 判断文件是否存在
	//if !com.IsExist(_DB_NAME) {
	//	// 创建目录  path.Dir 取出目录的路径
	//	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	// 以下是beego通用的注册数据库的方法
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic), new(Reply))
	// 使用beego中的orm中数据库驱动引擎
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRMySQL)
	// 创建数据库名字 强制有一个数据库必须是default
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, "root:ainiyu@/default?charset=utf8")
	// 账号为root 密码为ainiyu 数据库名字为 default
}

// ***********************category 操作**************************
// 数据库操作--添加category
func AddCategory(name string) error {
	// 初始化orm
	orm := orm.NewOrm()
	// 必须要是指针的形式
	cate := &Category{
		Title:   name,
		Created: time.Now(),
	}
	// 查询name是不是已经存在
	qs := orm.QueryTable("category")
	// 传入category的指针，去查询title为name的记录并返回回来
	err := qs.Filter("title", name).One(cate)
	// 当err为空表示再数据库中找到了title为name的记录，那么返回出去
	// 当err不为空表示，没有找到，那么进行后面的插入操作
	if err == nil {
		return err
	}
	// 发生错误表示插入失败
	_, err1 := orm.Insert(cate)
	if err1 != nil {
		return err
	}
	return nil
}

// 数据库操作--删除记录[ 根据ID 删除数据库中的记录]
func DeleteCategory(id string) error {
	// 转化为 10进制 int64类型
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	orm := orm.NewOrm()
	// beego orm操作方法 申请一个实体传入 然后读取实体内容来操作数据库
	cate := &Category{
		Id: Id,
	}
	_, err1 := orm.Delete(cate)
	return err1
}

// 获取数据库中所有category
func GetAllCategories() ([]*Category, error) {
	orm := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := orm.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

// ***********************topic 操作**************************
// 添加topic到数据库
func AddTopic(title, content, lable, category string) error {
	// 处理标签 空格作为多个标签的分隔符
	lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"

	orm := orm.NewOrm()
	topic := &Topic{
		Title:    title,
		Content:  content,
		Created:  time.Now(),
		Updated:  time.Now(),
		Category: category,
		Lables:   lable,
	}
	_, err := orm.Insert(topic)
	if err != nil {
		return err
	}
	// 更新分类统计
	cate := new(Category)
	qs := orm.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		// 如果存在，修改文章个数，然后更新记录
		cate.TopicCount++
		_, err = orm.Update(cate)
		if err != nil {
			return err
		}
	}
	return err
}

// 获得所有topics
// 参数：true表示按创建时间反序排列 false表示按照id正序排列
func GetAllTopics(cate, lable string, isDesc bool) ([]*Topic, error) {
	orm := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := orm.QueryTable("topic")
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(lable) > 0 {
			qs = qs.Filter("lables__contains", "$"+lable+"#")
		}
		_, err := qs.OrderBy("-created").All(&topics)
		return topics, err
	} else {
		_, err := qs.All(&topics)
		return topics, err
	}
}

// 通过id获取到文章
func GetTopic(id string) (*Topic, error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	orm := orm.NewOrm()
	topic := new(Topic)
	qs := orm.QueryTable("topic")
	err = qs.Filter("id", Id).One(topic)
	if err != nil {
		return nil, err
	}
	// 浏览数加1
	topic.Views++
	_, err = orm.Update(topic)

	topic.Lables = strings.Replace(
		strings.Replace(
			topic.Lables, "#", " ", -1),
		"$", "", -1)
	return topic, err
}

// 修改更新文章
func ModifyTopic(tid, title, content, lable, category string) error {
	IdNum, err := strconv.ParseInt(tid, 10, 64)
	lable = "$" + strings.Join(strings.Split(lable, ""), "#$") + "#"
	if err != nil {
		return err
	}
	// 表示未修改之前的分类
	var oldCate string
	orm := orm.NewOrm()
	topic := &Topic{
		Id: IdNum,
	}
	err1 := orm.Read(topic)
	// err1==nil 说明找到了id为idnum的记录 然后进行修改
	if err1 == nil {
		oldCate = topic.Category

		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		topic.Lables = lable
		_, err2 := orm.Update(topic)
		if err2 != nil {
			return err2
		}
	}
	// 更新旧的分类统计
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := orm.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = orm.Update(cate)
		}
	}
	// 更新新的分类统计
	cag := new(Category)
	qs := orm.QueryTable("category")
	err = qs.Filter("title", category).One(cag)
	if err == nil {
		cag.TopicCount++
		_, err = orm.Update(cag)
	}
	return nil
}

// 删除文章
func DeleteTopic(tid string) error {
	IdNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	orm := orm.NewOrm()
	topic := &Topic{
		Id: IdNum,
	}
	// 查询出文章
	if orm.Read(topic) == nil {
		oldCate := topic.Category
		_, err1 := orm.Delete(topic)
		if err1 != nil {
			return err1
		}
		cate := new(Category)
		qs := orm.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		cate.TopicCount--
		_, err = orm.Update(cate)
		if err != nil {
			return err
		}
	}
	return nil
}

// 删除分类的所有文章 -- 删除分类，那么把分了所有文章删除
func DeleteTopics(category string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("topic").Filter("category", category).Delete()
	if err != nil {
		return err
	}
	return nil
}
