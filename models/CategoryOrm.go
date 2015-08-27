package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddCategory(name string) error {
	o := orm.NewOrm()

	category := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}

	qs := o.QueryTable("category")
	err1 := qs.Filter("title", name).One(category)
	if err1 == nil {
		return err1
	}
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func QueryCategories(isDesc bool) ([]*Category, error) {
	o := orm.NewOrm()
	var err error

	categories := make([]*Category, 0)
	qs := o.QueryTable("category")
	if isDesc {
		_, err = qs.OrderBy("-created").All(&categories)
	} else {
		_, err = qs.All(&categories)
	}

	return categories, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	category := new(Category)
	err = o.QueryTable("category").Filter("id", cid).One(category)
	if err != nil {
		return err
	}
	_, err = o.QueryTable("topic").Filter("category", category.Title).Delete()
	if err != nil {
		return err
	}
	_, err = o.QueryTable("category").Filter("id", cid).Delete()
	if err != nil {
		return err
	}

	return err
}
