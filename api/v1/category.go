package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	var code int
	code = model.CheckCateNotExist(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}
func GetCates(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCates(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func UpdateCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckCateNotExist(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
