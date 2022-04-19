/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : main
 */
package main

import (
	widget2 "easygoadmin/app/widget"
	_ "easygoadmin/boot"
	cfg "easygoadmin/library/cfg"
	_ "easygoadmin/router"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 开始调试模式
	gin.SetMode("debug")

	// 实例化配置
	config := cfg.Instance()
	if config == nil {
		fmt.Printf("参数错误")
		return
	}

	//router := gin.Default()
	//
	//// 设置静态资源路由
	//router.Static("/resource", "./public/resource")
	//router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")
	//router.HTMLRender = LoadTemplates("views")
	//router.GET("level/index", controller.Level.Index)
	//router.GET("position/index", controller.Position.Index)
	//router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	//// 非模板嵌套
	//htmls, err := filepath.Glob(templatesDir + "/*.html")
	//if err != nil {
	//	panic(err.Error())
	//}
	//for _, html := range htmls {
	//	r.AddFromGlob(filepath.Base(html), html)
	//}

	// 布局模板
	layouts, err := filepath.Glob(templatesDir + "/public/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 嵌套的内容模板
	includes, err := filepath.Glob(templatesDir + "/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// template自定义函数
	funcMap := template.FuncMap{
		"StringToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"date": func() string {
			return time.Now().Format("2006-01-02 15:04:05.00000")
		},
		"widget":       widget2.Widget,
		"query":        widget2.Query,
		"add":          widget2.Add,
		"edit":         widget2.Edit,
		"delete":       widget2.Delete,
		"dall":         widget2.Dall,
		"expand":       widget2.Expand,
		"collapse":     widget2.Collapse,
		"addz":         widget2.Addz,
		"switch":       widget2.Switch,
		"select":       widget2.Select,
		"submit":       widget2.Submit,
		"icon":         widget2.Icon,
		"transfer":     widget2.Transfer,
		"upload_image": widget2.UploadImage,
		"album":        widget2.Album,
		"item":         widget2.Item,
		"kindeditor":   widget2.Kindeditor,
		//"date":         widget.Date,
		"checkbox": widget2.Checkbox,
		"radio":    widget2.Radio,
		"city":     widget2.City,
	}

	// 将主模板，include页面，layout子模板组合成一个完整的html页面
	for _, include := range includes {
		files := []string{}
		files = append(files, templatesDir+"/public/base.html", include)
		files = append(files, layouts...)
		r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
	}
	return r
}
