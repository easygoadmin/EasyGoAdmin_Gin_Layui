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

/**
 * 代码生成器-控制器
 * @author 半城风雨
 * @since 2021/11/15
 * @File : generate
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// 控制器管理对象
var Generate = new(generateCtl)

type generateCtl struct{}

func (c *generateCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "generate_index.html").WriteTpl()
}

func (c *generateCtl) List(ctx *gin.Context) {
	// 参数验证
	var req *dto.GeneratePageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用查询列表方法
	list, err := service.Generate.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: gconv.Int64(len(list)),
	})
}

func (c *generateCtl) Generate(ctx *gin.Context) {
	// 参数验证
	var req dto.GenerateFileReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用生成方法
	err := service.Generate.Generate(req, ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "模块生成成功",
	})
}

func (c *generateCtl) BatchGenerate(ctx *gin.Context) {
	// 生成对象
	var req *dto.BatchGenerateFileReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数分析
	tableList := strings.Split(req.Tables, ",")
	count := 0
	for _, item := range tableList {
		itemList := strings.Split(item, "|")
		// 组装参数对象
		var param dto.GenerateFileReq
		param.Name = itemList[0]
		param.Comment = itemList[1]
		// 调用生成方法
		err := service.Generate.Generate(param, ctx)
		if err != nil {
			continue
		}
		count++
	}
	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "本次共生成【" + strconv.Itoa(count) + "】个模块文件",
	})
}
