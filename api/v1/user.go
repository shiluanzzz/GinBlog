package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code int
	code = model.CheckUserNotExist(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func UpdateUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckUserNotExist(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// CheckUser 检查用户输入的密码是否正确
func CheckUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	var code int
	code = model.CheckPassword(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
func GetUser(c *gin.Context) {
	//todo
}
