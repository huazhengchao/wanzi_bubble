package main

import (
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
	Status bool `json:"status"`
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

	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	// 配置模板及静态资源目录
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	V1Group := r.Group("/v1")
	{
		// 获取
		V1Group.GET("/todo", func(c *gin.Context) {
			var todo []Todo
			DB.Find(&todo)
			c.JSON(http.StatusOK, todo)
		})
		// 添加
		V1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			DB.Select("title").Create(&todo)
			c.JSON(http.StatusOK, gin.H{
				"code" : 0,
				"msg"  : "成功",
				"data" : todo,
			})
		})
		// 更新
		V1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")

			var todo Todo
			DB.Where("id = ?", id).First(&todo)

			DB.Save(&todo)
			c.JSON(http.StatusOK, gin.H{
				"code" : 0,
				"msg"  : "成功",
			})
		})
		// 删除
		V1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			DB.Where("id = ?", id).Delete(&todo)

			c.JSON(http.StatusOK, gin.H{
				"code" : 0,
				"msg"  : "成功",
			})
		})
		

	}

	r.Run(":85")
}