package main

import (
	"GinBlog/model"
	"GinBlog/routes"
)

func main() {
	model.InitDB()

	routes.InitRouter()

}
