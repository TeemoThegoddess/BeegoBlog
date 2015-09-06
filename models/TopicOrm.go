package models

import (
	//"blog/controllers"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

func AddTopic(topicName, labels, category, topicContent string) error {
	//operate labels
	labels = "$" + strings.Join(strings.Split(labels, " "), "#$") + "#"

	o := orm.NewOrm()
	var sfContent string
	var topicCount []*Topic

	if len(topicContent) > 30 {
		sfContent = topicContent[0:30] + "..."
	} else {
		sfContent = topicContent
	}

	topic := &Topic{
		Title:           topicName,
		ShortForContent: sfContent,
		Content:         topicContent,
		Labels:          labels,
		Category:        category,
		Created:         time.Now(),
		Updated:         time.Now(),
		ReplyTime:       time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	topicCount, err = QueryTopicsByCategory(false, category)
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
		"TopicCount":      int64(len(topicCount)),
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

	if err != nil {
		return nil, err
	}

	return topics, err
}

func QueryTopicsByCategory(isDesc bool, category string) ([]*Topic, error) {
	o := orm.NewOrm()
	var err error

	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	if isDesc {
		_, err = qs.Filter("category", category).OrderBy("-updated").All(&topics)
	} else {
		_, err = qs.Filter("category", category).All(&topics)
	}

	if err != nil {
		return nil, err
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

	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)

	return topic, err
}

func ModifyTopic(topicId, title, labels, category, content string) error {
	//operate labels
	labels = "$" + strings.Join(strings.Split(labels, " "), "#$") + "#"
	topic := new(Topic)
	category1 := new(Category)
	category2 := new(Category)
	oldTopics := make([]*Topic, 0)
	newTopics := make([]*Topic, 0)

	tid, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	err = o.QueryTable("topic").Filter("id", tid).One(topic)
	if err != nil {
		return nil
	}
	oldCategory := topic.Category
	_, err = o.QueryTable("topic").Filter("id", tid).Update(orm.Params{
		"Title":    title,
		"Content":  content,
		"Labels":   labels,
		"Category": category,
		"Updated":  time.Now(),
	})
	if err != nil {
		return err
	}

	qs := o.QueryTable("category")
	err = qs.Filter("title", oldCategory).One(category1)
	if err == nil {
		oldTopics, _ = QueryTopicsByCategory(false, oldCategory)
		category1.TopicCount = int64(len(oldTopics))
		o.Update(category1)
	}
	err = qs.Filter("title", category).One(category2)
	if err == nil {
		newTopics, _ = QueryTopicsByCategory(false, category)
		category2.TopicCount = int64(len(newTopics))
		o.Update(category2)
	}

	return err
}

func DeleteTopic(topicId string) error {
	topicCount := make([]*Topic, 0)
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

	topicCount, err = QueryTopicsByCategory(false, topic1.Category)

	category1 := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", topic1.Category).One(category1)
	if err != nil {
		return err
	}
	_, err = qs.Filter("title", topic1.Category).Update(orm.Params{
		"TopicCount": int64(len(topicCount)),
	})
	if err != nil {
		return err
	}

	err = DeleteCommentWithTopic(topicId)
	if err != nil {
		return nil
	}

	return err
}

func QueryTopicsByLabel(isDesc bool, label string) ([]*Topic, error) {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")

	_, err := qs.Filter("labels__contains", label).OrderBy("-created").All(&topics)
	if err != nil {
		return nil, err
	}

	return topics, err
}

func GetAllTopics(isDesc bool, label, category string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")

	if len(category) > 0 {
		qs = qs.Filter("category", category)
	}
	if len(label) > 0 {
		qs = qs.Filter("labels__contains", label)
	}
	if isDesc {
		_, err = qs.OrderBy("-updated").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err

}
