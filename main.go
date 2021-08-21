/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : main
 */
package main

import (
	_ "easygoadmin/boot"
	_ "easygoadmin/router"
)

func main() {
	//r := gin.Default()
	//// 指定模板加载目录
	//r.LoadHTMLGlob("views/**/*")
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	//r.GET("/", func(context *gin.Context) {
	//
	//	context.HTML(http.StatusOK, "level/index.html", gin.H{
	//
	//		"title": "main.html title",
	//
	//		"content_before": "content 内容上部分",
	//
	//		"content_text": "content 内容部分",
	//
	//		"content_after": "content 内容下部分",
	//	})
	//
	//})

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
