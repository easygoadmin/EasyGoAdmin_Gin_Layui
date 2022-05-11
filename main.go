// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

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
