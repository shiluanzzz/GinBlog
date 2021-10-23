package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddArt(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	var code int = 200
	model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}
func GetArts(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, totalPage := model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"data":  data,
		"total": totalPage,
		"msg":   errmsg.GetErrMsg(code),
	})
}
func UpdateArt(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	model.EditArticle(id, &data)
	var code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

func GetAriById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.GetArticleById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

func GetALLArtByCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cateId, _ := strconv.Atoi(c.Query("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticleByCate(cateId, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}
