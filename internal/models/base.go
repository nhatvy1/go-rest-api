package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement;"`
	FirstName string     `json:"first_name" gorm:"column:firstName;not null;" validate:"required,min=1,max=50"`
	LastName  string     `json:"last_name" gorm:"column:lastName;not null;" validate:"required,min=1,max=50"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;default:null"`
}
