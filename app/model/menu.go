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

package model

import "easygoadmin/utils"

type Menu struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string `json:"name" xorm:"not null comment('菜单名称') index VARCHAR(30)"`
	Icon       string `json:"icon" xorm:"comment('图标') VARCHAR(50)"`
	Url        string `json:"url" xorm:"comment('URL地址') VARCHAR(150)"`
	Param      string `json:"param" xorm:"comment('参数') VARCHAR(150)"`
	Pid        int    `json:"pid" xorm:"not null default 0 comment('上级ID') index INT(10)"`
	Type       int    `json:"type" xorm:"not null default 0 comment('类型：1模块 2导航 3菜单 4节点') TINYINT(1)"`
	Permission string `json:"permission" xorm:"comment('权限标识') VARCHAR(150)"`
	Status     int    `json:"status" xorm:"not null default 1 comment('是否显示：1显示 2不显示') TINYINT(1)"`
	Target     int    `json:"target" xorm:"not null default 1 comment('打开方式：1内部打开 2外部打开') TINYINT(1)"`
	Note       string `json:"note" xorm:"comment('菜单备注') VARCHAR(255)"`
	Sort       int    `json:"sort" xorm:"default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int    `json:"create_user" xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime int64  `json:"create_time" xorm:"default 0 comment('创建时间') INT(11)"`
	UpdateUser int    `json:"update_user" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime int64  `json:"update_time" xorm:"default 0 comment('更新时间') INT(11)"`
	Mark       int    `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Menu) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Menu) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Menu) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Menu) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Menu{})
}

//批量删除
func (r *Menu) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Menu{})
}
