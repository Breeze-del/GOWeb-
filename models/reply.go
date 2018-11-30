package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// 评论表单
type Reply struct {
	Id       int64
	Tid      int64 // 文章的id
	Nickname string
	Content  string    `orm:"size(1000)"`
	Created  time.Time `orm:"index"`
}

// 将string转为int64
func S2int64(id string) (int64, error) {
	IdNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return IdNum, nil
}

// ***********************topic 操作**************************
// 添加评论信息
func AddReply(id, nickname, content string) error {
	o := orm.NewOrm()
	tid, err := S2int64(id)
	if err != nil {
		return err
	}
	reply := &Reply{
		Tid:      tid,
		Nickname: nickname,
		Content:  content,
		Created:  time.Now(),
	}
	_, err1 := o.Insert(reply)
	return err1
}

// 删除评论
func DeleteReply(id string) error {
	o := orm.NewOrm()
	idNum, err := S2int64(id)
	if err != nil {
		return err
	}
	reply := &Reply{
		Id: idNum,
	}
	_, err1 := o.Delete(reply)
	if err1 != nil {
		return err1
	}
	return nil
}

// 获取所有的评论
func GetAllReplies(tid string) ([]*Reply, error) {
	id, err := S2int64(tid)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("reply")
	replies := make([]*Reply, 0)
	_, err1 := qs.Filter("tid", id).All(&replies)
	return replies, err1
}
