package file

import (
	"gorm.io/gorm"
	"time"
	"took/pkg/account/user"
)

type File struct {
	Id        int            `gorm:"column:id;primary_key" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Size      int            `gorm:"column:size" json:"size"`
	Type      string         `gorm:"column:type" json:"type"`
	Bucket    string         `gorm:"column:bucket" json:"bucket"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
	CreatorId int            `gorm:"column:creator_id" json:"creatorId"`

	Creator *user.User `json:"Creator"`
}
