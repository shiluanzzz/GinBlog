package model

import (
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	// category 是文章的子类 方便连接查询
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

// CreateArticle
//  @Description: 新增文章
//
func CreateArticle(cate *Article) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticleByCate
//  @Description: 通过分类id查询文章
//  @param cate 分类id
//  @param pageSize 页面数据量
//  @param pageNum 页码
//  @return []Article
//  @return int
//
func GetArticleByCate(cate int, pageSize, pageNum int) ([]Article, int) {
	var data []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", cate).Find(&data).Error
	if err != nil {
		return data, errmsg.ERROR_CATENAME_NOT_EXIST
	}
	return data, errmsg.SUCCESS
}

// GetArticleById 查询单个文章信息
func GetArticleById(id int) (Article, int) {
	var data Article
	err := db.Preload("Category").Where("id = ?", id).First(&data).Error
	if err != nil {
		return Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return data, errmsg.SUCCESS

}

// GetArticles  查询文章列表，带有分类标签
func GetArticles(pageSize, pageNum int) ([]Article, int) {
	var data []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
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
