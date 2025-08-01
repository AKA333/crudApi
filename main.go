package main

import (
	"crudApi/controllers"
	"crudApi/internals"
	"crudApi/models"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	internals.LoadEnvVariables()
	internals.ConnectDB()
}

func main() {
	if !internals.DB.Migrator().HasTable(&models.Post{}) {
		if err := internals.DB.AutoMigrate(&models.Post{}); err != nil {
			log.Printf("Error migrating database: %v", err)
		}
	}

	r := gin.Default()
	r.POST("/createPost", controllers.CreatePost)
	r.GET("/getPosts", controllers.GetPost)
	r.GET("/getPost/:id", controllers.GetPostWithId)
	r.PUT("/updatePost/:id", controllers.UpdatePost)
	r.DELETE("/deletePost/:id", controllers.DeletePost)
	r.Run()
}
