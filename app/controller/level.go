/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : level
 */
package controller

import (
	"easygoadmin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Level = new(levelCtl)

type levelCtl struct{}

func (ctl *levelCtl) Index(c *gin.Context) {
	fmt.Println("测试路由")
	// 渲染模板
	response.BuildTpl(c, "/level/index.html").WriteTpl()
}
