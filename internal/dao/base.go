package dao

import "gorm.io/gorm"

type AllDao struct {
	db *gorm.DB
}

func NewAllDao(db *gorm.DB) AllDao {
	return AllDao{
		db: db,
	}
}

// 返回db
func (a AllDao) DB() *gorm.DB {
	return a.db
}
