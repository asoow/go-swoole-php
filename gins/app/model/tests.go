package model

import (
	"errors"
	"gins/config/db"
)

type Tests struct {
	//Model
	ID    uint   `gorm:"primary_key,column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
}

//创建用户
func (w *Tests) CreateTests(title string) (err error, id uint) {
	Db := db.DB
	ts := Tests{Title: title}
	res := Db.Create(&ts)
	err = res.Error

	if _, ok := res.Value.(*Tests); ok {
		id = ts.ID
	} else {
		err = errors.New("write-error")
	}
	return
}

func (w *Tests) GetIdByTests(id uint) (wd Tests, err error) {
	Db := db.DB
	Db.Where("id = ?", id).First(&wd)
	return
}

func (w *Tests) DeleteTests(Tests Tests) (err error) {
	Db := db.DB
	res := Db.Delete(&Tests)
	err = res.Error
	return
}
