/**
 *
 * @author 半城风雨
 * @since 2021/11/30
 * @File : analysis
 */
package controller

import (
	"easygoadmin/utils/response"
	"github.com/gin-gonic/gin"
)

var Analysis = new(analysisCtl)

type analysisCtl struct{}

func (c *analysisCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "analysis_index.html").WriteTpl()
}
