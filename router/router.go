/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"fmt"
)

func init() {
	fmt.Println("路由已加载")

	//router := gin.Default()

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
