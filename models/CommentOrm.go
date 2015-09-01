package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddComment(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	comment := &Comment{
		Tid:      tidNum,
		Nickname: nickname,
		Content:  content,
		Created:  time.Now(),
	}
	_, err = o.Insert(comment)
	if err != nil {
		return err
	}

	qs := o.QueryTable("topic")
	_, err = qs.Filter("id", tidNum).Update(orm.Params{
		"ReplyLastUserId": comment.Id,
		"ReplyTime":       comment.Created,
	})
	if err != nil {
		return err
	}

	return err
}

func GetAllComments(tid string) ([]*Comment, error) {
	o := orm.NewOrm()

	tidNum, err := strconv.ParseInt(tid, 10, 64)
	comments := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&comments)
	if err != nil {
		return nil, err
	}
	qs = o.QueryTable("topic")
	_, err = qs.Filter("id", tidNum).Update(orm.Params{
		"ReplyCount": len(comments),
	})

	return comments, err
}

func DeleteComment(id, tid string) error {
	o := orm.NewOrm()
	topic := new(Topic)
	/*	var maps []orm.Params
		var num int64*/
	comments := make([]*Comment, 0)

	idNum, err := strconv.ParseInt(id, 10, 64)
	tidNum, _ := strconv.ParseInt(tid, 10, 64)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("id", idNum).Delete()
	if err != nil {
		return err
	}
	qs = o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}
	comments, err = GetAllComments(tid)
	if len(comments) == 0 {
		_, err = qs.Filter("id", tidNum).Update(orm.Params{
			"ReplyLastUserId": 0,
			"ReplyTime":       topic.Created,
		})
	} else {
		topic.ReplyLastUserId = comments[0].Id
		topic.ReplyTime = comments[0].Created
		_, err = o.Update(topic)
	}

	return err
	/*if (topic.ReplyCount - 1) == 0 {
		_, err = qs.Filter("id", tidNum).Update(orm.Params{
			"ReplyLastUserId": 0,
			"ReplyTime":       topic.Created,
		})
	} else {
		num, err = o.Raw("SELECT created,id FROM comment WHERE created=(SELECT max(created) FROM comment)").Values(&maps)
		if err == nil && num > 0 {
			_, err = qs.Filter("id", tidNum).Update(orm.Params{
				"ReplyLastUserId": maps[0]["id"],
				"ReplyTime":       maps[0]["created"],
			})
		}
	}
	return err*/
}

func DeleteCommentWithTopic(id string) error {
	o := orm.NewOrm()

	tid, err := strconv.ParseInt(id, 10, 64)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tid).Delete()
	if err != nil {
		return err
	}

	return err
}
