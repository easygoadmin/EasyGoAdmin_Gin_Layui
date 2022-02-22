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

/**
 * 系统配置-控制器
 * @author 半城风雨
 * @since 2021/11/15
 * @File : config_web
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gstr"
	"easygoadmin/utils/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// 控制器管理对象
var Website = new(website)

type website struct{}

func (c *website) Index(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		// 返回结果
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "演示环境，暂无权限操作",
		})
		return
		// key：string类型，value：interface{}  类型能存任何数据类型
		var jsonObj map[string]interface{}
		data, _ := ctx.GetRawData()
		json.Unmarshal(data, &jsonObj)
		// 遍历处理数据源
		for key, val := range jsonObj {
			// 参数处理
			if key == "checkbox" {
				// 复选框

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
			} else if strings.Contains(key, "upimage") {
				// 单图上传

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				if strings.Contains(gconv.String(val), "temp") {
					image, _ := utils.SaveImage(gconv.String(val), "config")
					// 赋值给参数
					val = image
				} else {
					// 赋值给参数
					val = gstr.Replace(gconv.String(val), utils.ImageUrl(), "")
				}
			} else if strings.Contains(key, "upimgs") {
				// 多图上传
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				// 图片地址处理
				urlArr := gstr.Split(gconv.String(val), ",")
				list := make([]string, 0)
				for _, v := range urlArr {
					if strings.Contains(gconv.String(v), "temp") {
						image, _ := utils.SaveImage(v, "config")
						list = append(list, image)
					} else {
						image := gstr.Replace(v, utils.ImageUrl(), "")
						list = append(list, image)
					}
				}
				// 数组转字符串，逗号分隔
				val = strings.Join(list, ",")
			} else if strings.Contains(key, "ueditor") {
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]

				// 富文本处理(待完善)
				// TODO...
			}

			var info model.ConfigData
			has, err := utils.XormDb.Where("code=?", key).Get(&info)
			if err != nil || !has {
				continue
			}

			// 更新记录
			entity := &model.ConfigData{Id: info.Id}
			entity.Value = gconv.String(val)
			entity.UpdateUser = utils.Uid(ctx)
			entity.UpdateTime = time.Now().Unix()
			entity.Update()
		}

		// 返回结果
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: 0,
			Msg:  "保存成功",
		})
		return
	}
	// 配置ID
	configId := ctx.Query("configId")
	if configId == "" {
		configId = "1"
	}

	// 获取配置列表
	configData := make([]model.Config, 0)
	utils.XormDb.Where("mark=1").Find(&configData)
	configList := make(map[string]string)
	for _, v := range configData {
		configList[gconv.String(v.Id)] = v.Name
	}

	// 获取配置项列表
	itemData := make([]model.ConfigData, 0)
	utils.XormDb.Where("config_id=? and status=1 and mark=1", configId).OrderBy("sort asc").Find(&itemData)
	itemList := make([]map[string]interface{}, 0)
	for _, v := range itemData {
		item := make(map[string]interface{})
		item["id"] = v.Id
		item["title"] = v.Title
		item["code"] = v.Code
		item["value"] = v.Value
		item["type"] = v.Type

		if v.Type == "checkbox" {
			// 复选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, val := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "radio" {
			// 单选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList

		} else if v.Type == "select" {
			// 下拉选择框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "image" {
			// 单图片
			item["value"] = utils.GetImageUrl(v.Value)
		} else if v.Type == "images" {
			// 多图片
			list := gstr.Split(v.Value, ",")
			itemList := make([]string, 0)
			for _, v := range list {
				// 图片地址
				item := utils.GetImageUrl(v)
				itemList = append(itemList, item)
			}
			item["value"] = itemList
		}
		itemList = append(itemList, item)
	}

	// 渲染模板
	response.BuildTpl(ctx, "website_index.html").WriteTpl(gin.H{
		"configId":   configId,
		"configList": configList,
		"itemList":   itemList,
	})
}
