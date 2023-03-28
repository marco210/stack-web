package main
import (
	"web/controllers"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test GET func",
		})
	})

	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id",controllers.GetBookById)
	router.GET("/books", controllers.GetAllBook)

	router.Run("localhost:9000")
}