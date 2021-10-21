package model

import (
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

func CreateArticle(cate *Article) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticles TODO 查询文章
func GetArticles(pageSize, pageNum int) []Article {
	var data []Article
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}

func EditArticle(id int, article *Article) int {
	data := map[string]interface{}{}
	data["title"] = article.Title
	data["cid"] = article.Cid
	data["desc"] = article.Desc
	data["content"] = article.Content
	data["img"] = article.Img
	data["comment_count"] = article.CommentCount
	data["read_count"] = article.ReadCount
	err := db.Model(&Article{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		log.Println("更新文章信息出错,", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func DeleteArticle(id int) int {
	err := db.Model(&Article{}).Where("id=?", id).Delete(&Article{}).Error
	if err != nil {
		log.Println("删除用文章错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
