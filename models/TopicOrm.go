package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddTopic(topicName string, category string, topicContent string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:     topicName,
		Content:   topicContent,
		Category:  category,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	category1 := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(category1)
	if err != nil {
		return err
	}
	_, err = qs.Filter("title", category).Update(orm.Params{
		"TopicCount":      category1.TopicCount + 1,
		"TopicLastUserId": topic.Id,
	})
	if err != nil {
		return err
	}

	return err
}

func QueryTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	var err error

	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	if isDesc {
		_, err = qs.OrderBy("-updated").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func QueryTopic(topicId string, isModify bool) (*Topic, error) {
	o := orm.NewOrm()
	var err error
	tid, err1 := strconv.ParseInt(topicId, 10, 64)
	if err1 != nil {
		return nil, err
	}

	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	if !isModify {
		topic.Views++
	}
	_, err = o.Update(topic)

	return topic, err
}

func ModifyTopic(topicId, title, category, content string) error {
	tid, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	_, err = o.QueryTable("topic").Filter("id", tid).Update(orm.Params{
		"Title":    title,
		"Content":  content,
		"Category": category,
		"Updated":  time.Now(),
	})
	if err != nil {
		return err
	}

	return err
}

func DeleteTopic(topicId string) error {
	tid, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic1 := new(Topic)
	err = o.QueryTable("topic").Filter("id", tid).One(topic1)
	if err != nil {
		return err
	}

	_, err = o.QueryTable("topic").Filter("id", tid).Delete()
	if err != nil {
		return err
	}

	category1 := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", topic1.Category).One(category1)
	if err != nil {
		return err
	}
	_, err = qs.Filter("title", topic1.Category).Update(orm.Params{
		"TopicCount": category1.TopicCount - 1,
	})
	if err != nil {
		return err
	}

	return err
}
