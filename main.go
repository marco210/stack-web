package main
import (
	"github.com/marco210/stack-web/controllers"
	"github.com/gin-gonic/gin"
)
	
func main(){	
	

	router := gin.Default()
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test GET func",
		})
	})
	// router.GET("/books/new", controllers.CreateBookHandler)
	router.POST("/books", controllers.CreateBook)
	router.POST("/books/new", controllers.CreateBookForm)
	router.GET("/books/:id",controllers.GetBookById)
	router.GET("/books", controllers.GetAllBook)

	router.Run(":9000")
}