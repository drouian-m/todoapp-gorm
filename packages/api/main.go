package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type CreateTodoRequest struct {
	Text string `json:"text" binding:"required"`
}

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var input CreateTodoRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo := Todo{Text: input.Text, Checked: false}
		db.Create(&todo)

		c.JSON(http.StatusOK, todo)
	})

	r.POST("/todos/:id/check", func(c *gin.Context) {
		var todo Todo
		if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		todo.Checked = !todo.Checked

		db.Save(&todo)

		c.JSON(http.StatusOK, todo)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		db.Delete(&Todo{}, c.Param("id")) // WARN: do not use the autoincrement id in a real project

		c.Status(http.StatusNoContent)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
