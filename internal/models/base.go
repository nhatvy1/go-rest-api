package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;default:null"`
}
