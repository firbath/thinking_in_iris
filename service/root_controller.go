/*
@Time : 2020/11/15 20:42
@Author : Firbath
@File : root_controller.go
@Software: GoLand
*/
package service

import "github.com/kataras/iris/v12"

type RootController struct {
	/* dependencies */
}

// Book example.
type Root struct {
	Msg string `json:"msg"`
}

// GET: http://localhost:9090/
func (c *RootController) Get() Root {
	return Root{
		Msg: "Hello IRIS",
	}
}

// POST: http://localhost:9090/
func (c *RootController) Post(b Root) int {
	println("Received root: " + b.Msg)

	return iris.StatusCreated
}
