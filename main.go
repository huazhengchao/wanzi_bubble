package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status int `json:"status"`
}

func initMysql() (err error) {
	dsn := "root:123456@(127.0.0.1:3305)/bubble"
	DB, err = gorm.Open(mysql.Open(dsn), nil)
	return
}

func main() {

	err := initMysql()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 配置模板及静态资源目录
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	V1Group := r.Group("/v1")
	{
		V1Group.GET("/todo", func(c *gin.Context) {
			var todo []Todo
			DB.Find(&todo)
			c.JSON(http.StatusOK, todo)
		})
		

	}

	r.Run(":85")
}