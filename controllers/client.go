package controllers

import (
	"net/http"
	"github.com/marco210/stack-web/database"
	"github.com/marco210/stack-web/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context){
	db:= database.DBConn()

	// json
	var json models.Book

	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO books(id,title, author) VALUES(?,?,?)",)
		if err != nil {
			c.JSON(500, gin.H{
				"messages" : err,
			})
		}

		insPost.Exec(json.ID,json.Title, json.Author) 
		c.JSON(200, gin.H{
			"messages": "inserted",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}

func CreateBookForm(c *gin.Context){
	var form models.Book

	if err:=c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if  form.Title != "" || form.Author !="" {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
		return
	} 
}

func GetBookById(c *gin.Context){

}

func GetAllBook(c *gin.Context){
	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM books ")
	if err != nil{
		c.JSON(500, gin.H{
			"messages" : "Book not found",
		});
	}

	book := models.Book{}
	books := make([]models.Book, 0)

	for rows.Next(){
		var id int
		var title, author string

		err = rows.Scan(&id, &title, &author)
		if err != nil {
			panic(err.Error())
		}

		book.ID = id
		book.Title = title
		book.Author = author
		books = append(books,book)
	}

	c.IndentedJSON(http.StatusOK, books)
	defer db.Close()

}

