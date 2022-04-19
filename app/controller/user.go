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
 * 用户管理-控制器
 * @author 半城风雨
 * @since 2021/11/11
 * @File : user
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

var User = new(userCtl)

type userCtl struct{}

func (c *userCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "user_index.html").WriteTpl()
}

func (c *userCtl) List(ctx *gin.Context) {
	// 参数
	var req *dto.UserPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用获取列表方法
	list, count, err := service.User.GetList(req)
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
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}

func (c *userCtl) Edit(ctx *gin.Context) {
	// 获取职级
	levelAll := make([]model.Level, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&levelAll)
	levelList := make(map[int]string, 0)
	for _, v := range levelAll {
		levelList[v.Id] = v.Name
	}
	// 获取岗位
	positionAll := make([]model.Position, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&positionAll)
	positionList := make(map[int]string, 0)
	for _, v := range positionAll {
		positionList[v.Id] = v.Name
	}
	// 获取部门列表
	deptData, _ := service.Dept.GetDeptTreeList()
	deptList := service.Dept.MakeList(deptData)
	// 获取角色
	roleData := make([]model.Role, 0)
	utils.XormDb.Where("status=1 and mark=1").Find(&roleData)
	roleList := make(map[int]string)
	for _, v := range roleData {
		roleList[v.Id] = v.Name
	}

	// 记录ID
	id := gconv.Int(ctx.Query("id"))
	if id > 0 {
		// 编辑
		info := &model.User{Id: id}
		has, err := info.Get()
		if err != nil || !has {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		var userInfo = vo.UserInfoVo{}
		userInfo.User = *info
		// 头像
		userInfo.Avatar = utils.GetImageUrl(info.Avatar)

		// 角色ID
		var userRoleList []model.UserRole
		utils.XormDb.Where("user_id=?", utils.Uid(ctx)).Find(&userRoleList)
		roleIds := make([]interface{}, 0)
		for _, v := range userRoleList {
			roleIds = append(roleIds, v.RoleId)
		}
		userInfo.RoleIds = roleIds

		// 渲染模板
		response.BuildTpl(ctx, "user_edit.html").WriteTpl(gin.H{
			"info":         userInfo,
			"genderList":   utils.GENDER_LIST,
			"levelList":    levelList,
			"positionList": positionList,
			"deptList":     deptList,
			"roleList":     roleList,
		})
	} else {
		// 添加
		response.BuildTpl(ctx, "user_edit.html").WriteTpl(gin.H{
			"genderList":   utils.GENDER_LIST,
			"levelList":    levelList,
			"positionList": positionList,
			"deptList":     deptList,
			"roleList":     roleList,
		})
	}
}

func (c *userCtl) Add(ctx *gin.Context) {
	// 参数
	var req *dto.UserAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.User.Add(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *userCtl) Update(ctx *gin.Context) {
	// 参数
	var req *dto.UserUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.User.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *userCtl) Delete(ctx *gin.Context) {
	// 参数
	ids := ctx.Param("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}

	// 调用删除方法
	rows, err := service.User.Delete(ids)
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (c *userCtl) Status(ctx *gin.Context) {
	// 参数
	var req *dto.UserStatusReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用设置方法
	rows, err := service.User.Status(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}

func (c *userCtl) ResetPwd(ctx *gin.Context) {
	// 参数验证
	var req *dto.UserResetPwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用重置密码方法
	rows, err := service.User.ResetPwd(gconv.Int(req.Id), utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "重置密码成功",
	})
}

func (c *userCtl) CheckUser(ctx *gin.Context) {
	// 参数验证
	var req *dto.CheckUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用检查用户方法
	user, err := service.User.CheckUser(req)
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
		Msg:  "查询成功",
		Data: user,
	})
}
