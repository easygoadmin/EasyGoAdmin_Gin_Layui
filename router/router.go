/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/widget"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	r := gin.Default()

	// 设置模板函数
	r.SetFuncMap(template.FuncMap{
		"widget":   widget.Widget,
		"query":    widget.Query,
		"add":      widget.Add,
		"edit":     widget.Edit,
		"delete":   widget.Delete,
		"dall":     widget.Dall,
		"expand":   widget.Expand,
		"collapse": widget.Collapse,
		"addz":     widget.Addz,
	})

	// 设置静态资源路由
	r.Static("/resource", "./public/resource")
	r.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	// 指定模板加载目录
	r.LoadHTMLGlob("views/**/*")

	// 职级管理
	level := r.Group("/level")
	{
		level.GET("/index", controller.Level.Index)
	}

	// 启动
	r.Run()

	//// 职级管理
	//level := router.Group("/level")
	//{
	//	level.GET("/index", controller.Level.Index)
	//}

	//r := gin.New()
	//r.Use(gin.Recovery())

	//// 职级路由
	//r.Group("level", func(context *gin.Context) {
	//	r.GET("/index", controller.Level.Index)
	//})

	//r := gin.Default()
	//r.GET("/level/index", controller.Level.Index)

	//// 路由设置
	//auth := r.Group("/")
	//r.Use(middleware.CheckLogin())
	//{
	//	// 用户模块的路由接口
	//	auth.GET("/level/index", controller.Level.Index)
	//}

}
