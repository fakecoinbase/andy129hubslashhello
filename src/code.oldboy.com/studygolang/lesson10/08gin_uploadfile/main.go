package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)


// 处理单个文件上传
func uploadHandler(c *gin.Context){

	fileObj, err := c.FormFile("filename")   // "filename"  对应的是 .html 里面的 <input type="file" name="filename" >
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err, 
		})
		return 
	}

	saveFilePath := fmt.Sprintf("./%s", fileObj.Filename)
	err = c.SaveUploadedFile(fileObj, saveFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err, 
		})
		return 
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": "上传文件成功", 
	})

}	

// 处理多文件上传
func uploadMultiHandler(c *gin.Context){
	multi, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err, 
		})
		return 
	}

	files := multi.File["filename"]    // multi.File 是一个 map[string][]*multipart.FileHeader   

	for index, file := range files {
		// 获取完整文件名 (例如：test.txt)
		filenameWithSuffix := path.Base(file.Filename)
		// 获取文件名后缀 (例如：.txt)
		fileSuffix := path.Ext(filenameWithSuffix)
		// 将 完整的文件名去除后缀得到 不带后缀的文件名
		fileNameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
		// 拼接：文件名_index后缀,  得到类似于： ./test_1.txt
		dst := fmt.Sprintf("./%s_%d%s", fileNameOnly, index, fileSuffix)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "所有文件上传成功", 
	})

}


// gin -- uploadfile
func main() {
	router := gin.Default()

	// 处理multipart forms 提交文件时默认的内存限制是  32MiB   
	// // (大概意思就是 小于32M 的文件，都是先读取到内存中，然后从内存中落盘， 大于 32M 的文件则会先存储到临时文件中，然后从临时文件中落盘)
	// 可以通过下面的方式修改
	router.MaxMultipartMemory = 8 << 20    // 8 MiB

	router.LoadHTMLFiles("./upload.html", "./uploadmul.html")

	// 单文件上传
	router.GET("/upload", func(c *gin.Context){
		c.HTML(http.StatusOK, "upload.html",nil)
	})

	router.POST("/upload", uploadHandler)


	// 多文件上传
	router.GET("/uploadmul", func(c *gin.Context){
		c.HTML(http.StatusOK, "uploadmul.html", nil)
	})

	router.POST("/uploadmul", uploadMultiHandler)

	router.Run()
}