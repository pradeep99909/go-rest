package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	db := DB{}
	r := gin.Default()
	api := r.Group("/api")
	version := api.Group("/v1")
	version.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Main Page",
		})
	})
	version.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
			"status": 200,
		})
	})

	users := version.Group("/users")

	// User Routes
	users.GET("/", func(c *gin.Context) {
		users := db.getAll()
		c.JSON(200, gin.H{
			"message": "User List",
			"users": users,
		})
	})

	users.GET("/:id", func(c *gin.Context) {
		var id  int = c.Param("id")
		user := db.get(id)
		c.JSON(200, gin.H{
			"message": "User",
			"user": user,
		})
	})

	users.POST("/create", func(c *gin.Context){
		username := c.PostForm("username")
		
		c.JSON(200, gin.H{
			"message":"User created Sucessfully!",
			"user":db.create(username),
		})
	})

	users.PUT("/update", func(c *gin.Context){
		username := c.PostForm("username")
		new_username := c.PostForm("new_username")
		c.JSON(200, gin.H{
			"message":"return created User",
			"username": db.update(username, new_username).Name,
		})
	})

	users.DELETE("/delete", func(c *gin.Context){
		username := c.PostForm("username")
		db.delete(username)
		c.JSON(200, gin.H{
			"message":"User Deleted!",
		})	
	})


	r.Run("127.0.0.1:8080") // listen and serve on
}