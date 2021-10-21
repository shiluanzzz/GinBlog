package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"github.com/gin-gonic/gin"
	"log"
)
import "GinBlog/utils"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	PublicRouter := r.Group("api/v1")
	{
		//Public 分类
		PublicRouter.GET("/cate/list", v1.GetCates)
		//Public 文章
		PublicRouter.GET("/article/list", v1.GetArts)
		PublicRouter.POST("/article/cate", v1.GetALLArtByCate)
		PublicRouter.POST("/article/id", v1.GetAriById)
		// Login
		PublicRouter.POST("/login", v1.Login)
		PublicRouter.POST("/user/add", v1.AddUser)

	}
	adminRouter := r.Group("api/v1")
	adminRouter.Use(middleware.JwtToken())
	{
		// 用户相关
		adminRouter.DELETE("/user/delete/:id", v1.DeleteUser)
		adminRouter.PUT("/user/update/:id", v1.UpdateUser)
		adminRouter.GET("/user/list", v1.GetUsers)
		//分类相关
		adminRouter.POST("/cate/add", v1.AddCate)
		adminRouter.DELETE("/cate/delete/:id", v1.DeleteCate)
		adminRouter.PUT("/cate/update/:id", v1.UpdateCate)
		// 文章相关
		adminRouter.POST("/article/add", v1.AddArt)
		adminRouter.DELETE("/article/delete/:id", v1.DeleteArt)
		adminRouter.PUT("/article/update/:id", v1.UpdateArt)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		log.Fatalln("gin框架启动失败,", err)
	}
}
