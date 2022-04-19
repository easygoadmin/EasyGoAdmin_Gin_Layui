// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 全局路由
 * @author 半城风雨
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	widget2 "easygoadmin/app/widget"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	router := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("MIIEogIBAAKCAQEA2"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	router.Use(sessions.Sessions("easygoadmin", store))
	// 登录验证中间件
	router.Use(middleware.CheckLogin())

	// 设置模板文件目录
	router.HTMLRender = LoadTemplates("views")

	// 指定模板加载目录
	//router.LoadHTMLGlob("views/**/*")

	// 设置静态资源路由
	router.Static("/resource", "./public/resource")
	router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	/* 文件上传 */
	upload := router.Group("upload")
	{
		// 上传图片
		upload.POST("/uploadImage", controller.Upload.UploadImage)
	}

	/* 登录注册 */
	index := router.Group("/")
	{
		index.GET("/", controller.Index.Index)
		index.Any("/login", controller.Login.Login)
		index.GET("/captcha", controller.Login.Captcha)
		index.GET("/index", controller.Index.Index)
		index.GET("/main", controller.Index.Main)
		index.Any("/userInfo", controller.Index.UserInfo)
		index.Any("/updatePwd", controller.Index.UpdatePwd)
		index.GET("/logout", controller.Index.Logout)
	}

	//router.GET("/logout", controller.Index.Logout)

	/* 用户管理 */
	user := router.Group("user")
	{
		user.GET("/index", controller.User.Index)
		user.POST("/list", controller.User.List)
		user.GET("/edit", controller.User.Edit)
		user.POST("/add", controller.User.Add)
		user.POST("/update", controller.User.Update)
		user.POST("/delete/:ids", controller.User.Delete)
		user.POST("/setStatus", controller.User.Status)
		user.POST("/resetPwd", controller.User.ResetPwd)
		user.GET("/checkUser", controller.User.CheckUser)
	}

	// 职级管理
	level := router.Group("/level")
	{
		level.GET("/index", controller.Level.Index)
		level.POST("/list", controller.Level.List)
		level.GET("/edit", controller.Level.Edit)
		level.POST("/add", controller.Level.Add)
		level.POST("/update", controller.Level.Update)
		level.POST("/delete/:ids", controller.Level.Delete)
		level.POST("/setStatus", controller.Level.Status)
	}

	/* 岗位管理 */
	position := router.Group("position")
	{
		position.GET("/index", controller.Position.Index)
		position.POST("/list", controller.Position.List)
		position.GET("/edit", controller.Position.Edit)
		position.POST("/add", controller.Position.Add)
		position.POST("/update", controller.Position.Update)
		position.POST("/delete/:ids", controller.Position.Delete)
		position.POST("/setStatus", controller.Position.Status)
		position.GET("/getPositionList", controller.Position.GetPositionList)
	}

	/* 角色路由 */
	role := router.Group("role")
	{
		role.GET("/index", controller.Role.Index)
		role.POST("/list", controller.Role.List)
		role.GET("/edit", controller.Role.Edit)
		role.POST("/add", controller.Role.Add)
		role.POST("/update", controller.Role.Update)
		role.POST("/delete/:ids", controller.Role.Delete)
		role.POST("/setStatus", controller.Role.Status)
		role.GET("/getRoleList", controller.Role.GetRoleList)
	}

	/* 角色菜单权限 */
	roleMenu := router.Group("rolemenu")
	{
		roleMenu.GET("/index/:roleId", controller.RoleMenu.Index)
		roleMenu.POST("/save", controller.RoleMenu.Save)
	}

	/* 部门管理 */
	dept := router.Group("dept")
	{
		dept.GET("/index", controller.Dept.Index)
		dept.POST("/list", controller.Dept.List)
		dept.GET("/edit", controller.Dept.Edit)
		dept.POST("/add", controller.Dept.Add)
		dept.POST("/update", controller.Dept.Update)
		dept.POST("/delete/:ids", controller.Dept.Delete)
		dept.GET("/getDeptList", controller.Dept.GetDeptList)
	}

	/* 菜单管理 */
	menu := router.Group("menu")
	{
		menu.GET("/index", controller.Menu.Index)
		menu.POST("/list", controller.Menu.List)
		menu.GET("/edit", controller.Menu.Edit)
		menu.POST("/add", controller.Menu.Add)
		menu.POST("/update", controller.Menu.Update)
		menu.POST("/delete/:ids", controller.Menu.Delete)
	}

	///* 登录日志 */
	//loginLog := router.Group("loginlog")
	//{
	//	loginLog.GET("/index", controller.LoginLog.Index)
	//	loginLog.GET("/list", controller.LoginLog.List)
	//	loginLog.POST("/delete/:ids", controller.LoginLog.Delete)
	//}
	//
	///* 操作日志 */
	//operLog := router.Group("operlog")
	//{
	//	operLog.GET("/list", controller.OperLog.List)
	//}

	/* 友链管理 */
	link := router.Group("link")
	{
		link.GET("/index", controller.Link.Index)
		link.POST("/list", controller.Link.List)
		link.GET("/edit", controller.Link.Edit)
		link.POST("/add", controller.Link.Add)
		link.POST("/update", controller.Link.Update)
		link.POST("/delete/:ids", controller.Link.Delete)
		link.POST("/setStatus", controller.Link.Status)
	}

	/* 城市管理 */
	city := router.Group("city")
	{
		city.GET("/index", controller.City.Index)
		city.POST("/list", controller.City.List)
		city.GET("/edit", controller.City.Edit)
		city.POST("/add", controller.City.Add)
		city.POST("/update", controller.City.Update)
		city.POST("/delete/:ids", controller.City.Delete)
		city.POST("/getChilds", controller.City.GetChilds)
	}

	/* 站点管理 */
	item := router.Group("item")
	{
		item.GET("/index", controller.Item.Index)
		item.POST("/list", controller.Item.List)
		item.GET("/edit", controller.Item.Edit)
		item.POST("/add", controller.Item.Add)
		item.POST("/update", controller.Item.Update)
		item.POST("/delete/:ids", controller.Item.Delete)
		item.POST("/setStatus", controller.Item.Status)
		item.GET("/getItemList", controller.Item.GetItemList)
	}

	/* 栏目管理 */
	itemcate := router.Group("itemcate")
	{
		itemcate.GET("/index", controller.ItemCate.Index)
		itemcate.POST("/list", controller.ItemCate.List)
		itemcate.GET("/edit", controller.ItemCate.Edit)
		itemcate.POST("/add", controller.ItemCate.Add)
		itemcate.POST("/update", controller.ItemCate.Update)
		itemcate.POST("/delete/:ids", controller.ItemCate.Delete)
		itemcate.GET("/getCateList", controller.ItemCate.GetCateList)
		itemcate.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
	}

	/* 通知管理 */
	notice := router.Group("notice")
	{
		notice.GET("/index", controller.Notice.Index)
		notice.POST("/list", controller.Notice.List)
		notice.GET("/edit", controller.Notice.Edit)
		notice.POST("/add", controller.Notice.Add)
		notice.POST("/update", controller.Notice.Update)
		notice.POST("/delete/:ids", controller.Notice.Delete)
		notice.POST("/setStatus", controller.Notice.Status)
	}

	/* 会员等级 */
	memberlevel := router.Group("memberlevel")
	{
		memberlevel.GET("/index", controller.MemberLevel.Index)
		memberlevel.POST("/list", controller.MemberLevel.List)
		memberlevel.GET("/edit", controller.MemberLevel.Edit)
		memberlevel.POST("/add", controller.MemberLevel.Add)
		memberlevel.POST("/update", controller.MemberLevel.Update)
		memberlevel.POST("/delete/:ids", controller.MemberLevel.Delete)
		memberlevel.GET("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	}

	/* 会员管理 */
	member := router.Group("member")
	{
		member.GET("/index", controller.Member.Index)
		member.POST("/list", controller.Member.List)
		member.GET("/edit", controller.Member.Edit)
		member.POST("/add", controller.Member.Add)
		member.POST("/update", controller.Member.Update)
		member.POST("/delete/:ids", controller.Member.Delete)
		member.POST("/setStatus", controller.Member.Status)
	}

	/* 广告位管理 */
	adsort := router.Group("adsort")
	{
		adsort.GET("/index", controller.AdSort.Index)
		adsort.POST("/list", controller.AdSort.List)
		adsort.GET("/edit", controller.AdSort.Edit)
		adsort.POST("/add", controller.AdSort.Add)
		adsort.POST("/update", controller.AdSort.Update)
		adsort.POST("/delete/:ids", controller.AdSort.Delete)
		adsort.GET("/getAdSortList", controller.AdSort.GetAdSortList)
	}

	/* 广告管理 */
	ad := router.Group("ad")
	{
		ad.GET("/index", controller.Ad.Index)
		ad.POST("/list", controller.Ad.List)
		ad.GET("/edit", controller.Ad.Edit)
		ad.POST("/add", controller.Ad.Add)
		ad.POST("/update", controller.Ad.Update)
		ad.POST("/delete/:ids", controller.Ad.Delete)
		ad.POST("/setStatus", controller.Ad.Status)
	}

	/* 字典管理 */
	dict := router.Group("dict")
	{
		dict.GET("/index", controller.Dict.Index)
		dict.POST("/list", controller.Dict.List)
		dict.POST("/add", controller.Dict.Add)
		dict.POST("/update", controller.Dict.Update)
		dict.POST("/delete/:ids", controller.Dict.Delete)
	}

	/* 字典项管理 */
	dictdata := router.Group("dictdata")
	{
		dictdata.POST("/list", controller.DictData.List)
		dictdata.POST("/add", controller.DictData.Add)
		dictdata.POST("/update", controller.DictData.Update)
		dictdata.POST("/delete/:ids", controller.DictData.Delete)
	}

	/* 配置管理 */
	config := router.Group("config")
	{
		config.GET("/index", controller.Config.Index)
		config.POST("/list", controller.Config.List)
		config.POST("/add", controller.Config.Add)
		config.POST("/update", controller.Config.Update)
		config.POST("/delete/:ids", controller.Config.Delete)
	}

	/* 配置项管理 */
	configdata := router.Group("configdata")
	{
		configdata.POST("/list", controller.ConfigData.List)
		configdata.POST("/add", controller.ConfigData.Add)
		configdata.POST("/update", controller.ConfigData.Update)
		configdata.POST("/delete/:ids", controller.ConfigData.Delete)
		configdata.POST("/setStatus", controller.ConfigData.Status)
	}

	/* 网站设置 */
	configweb := router.Group("configweb")
	{
		configweb.Any("/index", controller.ConfigWeb.Index)
	}

	/* 统计分析 */
	analysis := router.Group("analysis")
	{
		analysis.GET("/index", controller.Analysis.Index)
	}

	/* 代码生成器 */
	generate := router.Group("generate")
	{
		generate.GET("/index", controller.Generate.Index)
		generate.POST("/list", controller.Generate.List)
		generate.POST("/generate", controller.Generate.Generate)
		generate.POST("/batchGenerate", controller.Generate.BatchGenerate)
	}

	/* 演示一 */
	example := router.Group("example")
	{
		example.GET("/index", controller.Example.Index)
		example.POST("/list", controller.Example.List)
		example.GET("/edit", controller.Example.Edit)
		example.POST("/add", controller.Example.Add)
		example.POST("/update", controller.Example.Update)
		example.POST("/delete/:ids", controller.Example.Delete)
		example.POST("/setStatus", controller.Example.Status)
		example.POST("/setIsVip", controller.Example.IsVip)
	}

	/* 演示二 */
	example2 := router.Group("example2")
	{
		example2.GET("/index", controller.Example2.Index)
		example2.POST("/list", controller.Example2.List)
		example2.GET("/edit", controller.Example2.Edit)
		example2.POST("/add", controller.Example2.Add)
		example2.POST("/update", controller.Example2.Update)
		example2.POST("/delete/:ids", controller.Example2.Delete)
		example2.POST("/setStatus", controller.Example2.Status)
	}
	// 启动
	router.Run(":8080")
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// 非模板嵌套
	htmls, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, html := range htmls {
		r.AddFromGlob(filepath.Base(html), html)
	}

	// 布局模板
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 嵌套的内容模板
	includes, err := filepath.Glob(templatesDir + "/includes/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// template自定义函数
	funcMap := template.FuncMap{
		"StringToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"date2": func() string {
			return time.Now().Format("2006-01-02 15:04:05.00000")
		},
		"safe": func(str string) template.HTML {
			return template.HTML(str)
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
		"date":         widget2.Date,
		"checkbox":     widget2.Checkbox,
		"radio":        widget2.Radio,
		"city":         widget2.City,
	}

	// 将主模板，include页面，layout子模板组合成一个完整的html页面
	for _, include := range includes {
		// 文件名称
		baseName := filepath.Base(include)
		files := []string{}
		if strings.Contains(baseName, "edit") {
			files = append(files, templatesDir+"/layouts/form.html", include)
		} else if strings.Contains(baseName, "dict") {
			files = append(files, templatesDir+"/layouts/main.html", include)
			// 字典
			dict, _ := filepath.Glob(templatesDir + "/includes/dict/*.html")
			files = append(files, dict...)
		} else if baseName == "config" {
			files = append(files, templatesDir+"/layouts/main.html", include)
			// 字典
			dict, _ := filepath.Glob(templatesDir + "/includes/config/*.html")
			files = append(files, dict...)
		} else {
			files = append(files, templatesDir+"/layouts/layout.html", include)
		}
		files = append(files, layouts...)
		r.AddFromFilesFuncs(baseName, funcMap, files...)
	}
	return r
}
