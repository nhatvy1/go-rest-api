package user_module

import (
	"database/sql/driver"
	"user-service/pkg/commons"
)

type Status string

const (
	Active  Status = "active"
	Blocked Status = "blocked"
	Pending Status = "pending"
)

func (s *Status) Scan(value interface{}) error {
	*s = Status(value.([]byte))
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}

type User struct {
	commons.BaseModel
	FirstName string `json:"first_name" gorm:"column:firstName;not null;" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" gorm:"column:lastName;not null;" validate:"required,min=1,max=50"`
	Email     string `json:"email" gorm:"column:email;not null;unique;" validate:"email,required"`
	Password  string `json:"password" gorm:"column:password;not null;" validate:"required"`
	Status    Status `json:"status" gorm:"column:status;type:ENUM('active', 'blocked', 'pending');default:'active';"`
}
