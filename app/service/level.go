/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : level
 */
package service

import (
	"easygoadmin/app/model"
	"easygoadmin/library/db"
	"easygoadmin/utils"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
)

var x *xorm.Engine

// 中间件管理服务
var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req *model.LevelPageReq) ([]model.Level, int64, error) {
	// 实例化引擎
	db := db.Instance().Engine()
	if db == nil {
		return nil, 0, errors.New("连接数据库失败")
	}
	// 初始化查询实例
	query := db.Table(model.TableName()).Where("mark=1")
	if req != nil {
		// 职级名称查询
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 查询记录总数
	totalQuery := query.Clone()
	count, err := totalQuery.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Asc("sort")
	// 分页
	query = query.Limit(req.Limit, 0)
	// 对象转换
	var list []model.Level
	query.Find(&list)
	return list, count, nil
}

func (s *levelService) Add(req *model.LevelAddReq, userId int) (int, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Level
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	_, err := entity.Insert()
	if err != nil {
		return 0, err
	}

	return entity.Id, nil
}

func (s *levelService) Update(req *model.LevelUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	fmt.Printf("记录ID:%d", req.Id)
	fmt.Println(req)
	// 查询记录
	entity := &model.Level{Id: req.Id}
	ok, err := entity.FindOne()
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("记录不存在")
	}
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *levelService) Delete(Ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	//idsArr := convert.ToInt64Array(Ids, ",")
	// 批量删除
	//rows, err := model.Level.Delete(idsArr)
	//if err != nil {
	//	return 0, err
	//}
	//return rows, nil
	return 0, nil
}

func (s *levelService) Status(req *model.LevelStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	//info, err := dao.Level.FindOne("id=?", req.Id)
	//if err != nil {
	//	return 0, err
	//}
	//if info == nil {
	//	return 0, gerror.New("记录不存在")
	//}
	//
	//// 设置状态
	//result, err := dao.Level.Data(g.Map{
	//	"status":      req.Status,
	//	"update_user": userId,
	//	"update_time": gtime.Now(),
	//}).Where(dao.Level.Columns.Id, info.Id).Update()
	//if err != nil {
	//	return 0, err
	//}
	//res, err := result.RowsAffected()
	//if err != nil {
	//	return 0, err
	//}
	//return res, nil
	return 0, nil
}
