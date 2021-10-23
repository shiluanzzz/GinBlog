package model

import (
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCateNotExist(name string) int {
	var data Category
	db.Select("id").Where("name=?", name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}
func CreateCate(cate *Category) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetCates(pageSize, pageNum int) ([]Category, int64) {
	var data []Category
	var count int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&data).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return data, count
}

func EditCate(id int, cate *Category) int {
	data := map[string]interface{}{}
	data["name"] = cate.Name
	err := db.Model(&Category{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		log.Println("更新分类信息出错,", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func DeleteCate(id int) int {
	err := db.Model(&Category{}).Where("id=?", id).Delete(&Category{}).Error
	if err != nil {
		log.Println("删除用分类错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
