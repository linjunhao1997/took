package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int    `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Email    string `gorm:"column:email" json:"email"`
	//Roles  []*Role     `gorm:"many2many:sys_user_r_sys_role" json:"roles"`
	Disabled       int8           `gorm:"column:disabled" json:"disabled"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"CreatedAt"`
	lastModifiedAt time.Time      `gorm:"column:last_modified_at" json:"lastModifiedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}
