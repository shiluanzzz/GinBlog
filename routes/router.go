package routes

import (
	v1 "GinBlog/api/v1"
	"github.com/gin-gonic/gin"
)
import "GinBlog/utils"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r:=gin.Default()
	router:=r.Group("api/v1")
	{
		router.POST("user/add",v1.AddUser)
		router.GET("user/list",v1.GetUsers)
		router.GET("user/get/:id",v1.GetUser)
		router.POST("user/delete",v1.DeleteUser)
		router.POST("user/update",v1.UpdateUser)
	}
	r.Run(utils.HttpPort)
}
