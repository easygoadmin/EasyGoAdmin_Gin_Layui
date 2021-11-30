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

	// 设置模板文件目录
	router.HTMLRender = LoadTemplates("views")

	// 指定模板加载目录
	//router.LoadHTMLGlob("views/**/*")

	// 设置静态资源路由
	router.Static("/resource", "./public/resource")
	router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	/* 登录注册 */
	index := router.Group("/")
	{
		index.GET("/", controller.Login.Login)
		index.Any("/login", controller.Login.Login)
		index.GET("/captcha", controller.Login.Captcha)
		index.GET("/index", controller.Index.Index)
		index.GET("/main", controller.Index.Main)
		index.Any("/userInfo", controller.Index.UserInfo)
		index.Any("/updatePwd", controller.Index.UpdatePwd)
		index.GET("/logout", controller.Index.Logout)
	}

	/* 用户管理 */
	user := router.Group("user")
	{
		user.GET("/index", controller.User.Index)
		user.POST("/list", controller.User.List)
		user.GET("/edit", controller.User.Edit)
		user.POST("/add", controller.User.Add)
		user.POST("/update", controller.User.Update)
		user.POST("/delete", controller.User.Delete)
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
		level.POST("/delete", controller.Level.Delete)
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
		position.POST("/delete", controller.Position.Delete)
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
		role.POST("/delete", controller.Role.Delete)
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
		dept.POST("/delete", controller.Dept.Delete)
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
		configdata.POST("/status", controller.ConfigData.Status)
	}

	/* 网站设置 */
	website := router.Group("website")
	{
		website.Any("/index", controller.Website.Index)
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
		example.POST("/status", controller.Example.Status)
		example.POST("/isVip", controller.Example.IsVip)
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
		example2.POST("/status", controller.Example2.Status)
	}

	// 启动
	router.Run(":9199")
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
		"widget":       widget.Widget,
		"query":        widget.Query,
		"add":          widget.Add,
		"edit":         widget.Edit,
		"delete":       widget.Delete,
		"dall":         widget.Dall,
		"expand":       widget.Expand,
		"collapse":     widget.Collapse,
		"addz":         widget.Addz,
		"switch":       widget.Switch,
		"select":       widget.Select,
		"submit":       widget.Submit,
		"icon":         widget.Icon,
		"transfer":     widget.Transfer,
		"upload_image": widget.UploadImage,
		"album":        widget.Album,
		"item":         widget.Item,
		"kindeditor":   widget.Kindeditor,
		"date":         widget.Date,
		"checkbox":     widget.Checkbox,
		"radio":        widget.Radio,
		"city":         widget.City,
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
		} else if strings.Contains(baseName, "config") {
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
