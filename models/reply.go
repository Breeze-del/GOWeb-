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
// 添加评论信息,更新文章回复数和最近回复时间
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
	if err1 != nil {
		return err1
	}
	// 更新文章的回复时间和回复数
	topic := &Topic{
		Id: tid,
	}
	if o.Read(topic) == nil {
		topic.ReplyCount++
		topic.ReplyTime = reply.Created
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}
	return nil
}

// 删除评论,更新文章回复数和最近回复时间
// id 为回复的ID
func DeleteReply(id, tid string) error {
	o := orm.NewOrm()
	idNum, err := S2int64(id)
	topicId, err := S2int64(tid)
	if err != nil {
		return err
	}
	reply := &Reply{
		Id:  idNum,
		Tid: topicId,
	}
	_, err1 := o.Delete(reply)
	if err1 != nil {
		return err1
	}
	topic := new(Topic)
	topic.Id = topicId
	// 精确化最后评论时间和评论数
	replies := make([]*Reply, 0)
	qs := o.QueryTable("reply")
	_, err = qs.Filter("tid", topicId).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	if o.Read(topic) == nil {
		topic.ReplyCount = int64(len(replies))
		if len(replies) == 0 {
			topic.ReplyTime = time.Now()
		} else {
			topic.ReplyTime = replies[0].Created
		}
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
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

// 删除相关文章所有评论
func DeleteReplies(tid string) error {
	id, err := S2int64(tid)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	replies := make([]*Reply, 0)
	_, err = o.QueryTable("reply").Filter("tid", id).Delete()
	if err != nil {
		return err
	}
	if len(replies) > 0 {
		_, err := o.Delete(replies)
		if err != nil {
			return err
		}
	}
	return nil
}
