package models

import (
	"github.com/x-funs/go-fun/alias"
	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt alias.DateTime `json:"created_at" swaggertype:"string"`
	UpdatedAt alias.DateTime `json:"updated_at" swaggertype:"string"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" swaggertype:"string" gorm:"index"`
}
