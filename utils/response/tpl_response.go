// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

// 返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

// 返回一个错误的tpl响应
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

// 输出页面模板附加自定义函数
func (resp *TplResp) WriteTpl(params ...gin.H) {
	if params == nil || len(params) == 0 {
		fmt.Println("渲染")
		fmt.Println(resp.tpl)
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": 1})
	} else {
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
}
