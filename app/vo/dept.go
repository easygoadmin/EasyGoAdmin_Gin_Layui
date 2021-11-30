package vo

import "easygoadmin/app/model"

// 部门树结构
type DeptTreeNode struct {
	model.Dept
	Children []*DeptTreeNode `json:"children"` // 子栏目
}
