/*
@Time : 2020/11/15 20:42
@Author : Firbath
@File : book_controller.go
@Software: GoLand
*/
package service

import "github.com/kataras/iris/v12"

// Book example.
type Book struct {
	Title string `json:"title"`
}
type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
	return []Book{
		{"HelloWorld"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b Book) int {
	println("Received Book: " + b.Title)

	return iris.StatusCreated
}
