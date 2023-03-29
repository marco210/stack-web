package models

// json
type Book struct{
	ID int `json:"id" gorm:"primary_key" form:"id"`
	Title string `json:"title" form:"title"`
	Author string `json:"author form:"author"`
}


// form
// type Book struct{
// 	ID string `form:"id" gorm:"primary_key"`
// 	Title string `form:"title"`
// 	Author string `form:"author"`
// }