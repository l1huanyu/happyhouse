package main

import (
	"happyhouse/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	h := new(handler.Handler)
	e.GET("/wechat", h.CheckSignature)
	e.POST("/wechat", h.ReceiveMessage)
	e.Start(":8823")
}
