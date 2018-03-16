package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// 미들웨서 사용
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {				// 라우팅
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello/:data", func(c echo.Context) error {	// 쿼리파라미터 받기
		name := c.QueryParam("name")
		age := c.QueryParam("age")
		dataType := c.Param("data")

		if dataType == "string" {
			return c.String(http.StatusOK, fmt.Sprintf("your name is %s and you are %s\n", name, age))

		} else if dataType == "json" {
			return c.JSON(http.StatusOK, map[string]string {
				"name": name,
				"age": age,
			})
		}

		return c.String(http.StatusOK, "Error!!!")
	})

	e.Logger.Fatal(e.Start(":1323")) // 1323번 포트로 서버 시작
}
