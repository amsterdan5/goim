package model

import "gorm.io/plugin/soft_delete"

type baseModel struct {
	ID         uint64                `gorm:"column:id"`
	CreateTime uint64                `gorm:"column:create_time"` // 创建时间
	Updatetime uint64                `gorm:"column:update_time"` // 更新时间
	CreateBy   uint64                `gorm:"column:create_by"`
	UpdateBy   uint64                `gorm:"column:update_by"`
	IsDel      soft_delete.DeletedAt `gorm:"column:is_delete,softDelete:flag"` // 已删除, 0:否, 1:是
}
