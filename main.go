package main

import (
	"html/template"
	"path/filepath"
	"onlinetools/auth"
	"onlinetools/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(auth.AuthMiddleware())

	r.SetFuncMap(template.FuncMap{
		"base": func(path string) string {
			return filepath.Base(path)
		},
	})

	// 创建文件存储目录
	if _, err := os.Stat("./filename"); os.IsNotExist(err) {
		os.Mkdir("./filename", 0755)
	}

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routes.RegisterRoutes(r)
	r.Run(":8080")
}