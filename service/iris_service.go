/*
@Time : 2020/11/15 18:32
@Author : Firbath
@File : iris_service.go
@Software: GoLand
*/
package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func IrisRun(host string, port int) {
	app := iris.New()
	m := mvc.New(app.Party("/books"))
	m.Handle(new(BookController))
	app.Listen(fmt.Sprintf("%s:%d", host, port))
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

// Book example.
type Book struct {
	Title string `json:"title"`
}
