package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	
	"github.com/gin-gonic/gin"
	"onlinetools/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", listFiles)
	r.POST("/upload", uploadFile)
	r.GET("/download/:filename", downloadFile)
	r.POST("/delete/:filename", deleteFile)
	r.PUT("/tools/:id", handlers.UpdateHandler)  // 修改为PUT方法并指向新处理器
}

func listFiles(c *gin.Context) {
	files, _ := filepath.Glob("./filename/*")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Files": files,
	})
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "请选择要上传的文件"})
    return
}
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "请选择要上传的文件"})
    return
}
	dst := filepath.Join("./filename", file.Filename)
	c.SaveUploadedFile(file, dst)
	c.Redirect(http.StatusFound, "/")
}

func downloadFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("./filename", filename)
	
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.File(filePath)
}

func deleteFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("./filename", filename)
	os.Remove(filePath)
	c.Redirect(http.StatusFound, "/")
}

func updateFile(c *gin.Context) {
	file, err := c.FormFile("file")
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "请选择要上传的文件"})
    return
}
filename := file.Filename
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "未找到上传文件"})
		return
	}

	dst := filepath.Join("./filename", filename)
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "无法创建目录"})
		return
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}