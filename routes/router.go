package routes

import (
	v1 "GinBlog/api/v1"
	"github.com/gin-gonic/gin"
)
import "GinBlog/utils"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routerUser := r.Group("api/v1/user")
	{
		routerUser.POST("/add", v1.AddUser)
		routerUser.GET("/list", v1.GetUsers)
		routerUser.DELETE("/delete/:id", v1.DeleteUser)
		routerUser.PUT("/update/:id", v1.UpdateUser)
		routerUser.POST("/check", v1.CheckUser)
		routerUser.GET("/get/:id", v1.GetUser)
	}
	routerCate := r.Group("api/v1/cate")
	{
		routerCate.POST("/add", v1.AddCate)
		routerCate.GET("/list", v1.GetCates)
		routerCate.DELETE("/delete/:id", v1.DeleteCate)
		routerCate.PUT("/update/:id", v1.UpdateCate)
	}
	routerArt := r.Group("api/v1/article")
	{
		routerArt.POST("/add", v1.AddArt)
		routerArt.GET("/list", v1.GetArts)
		routerArt.DELETE("/delete/:id", v1.DeleteArt)
		routerArt.PUT("/update/:id", v1.UpdateArt)
	}

	r.Run(utils.HttpPort)
}
