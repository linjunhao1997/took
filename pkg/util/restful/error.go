package restful

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound  = gorm.ErrRecordNotFound
	ErrInvalidArgument = errors.New("invalid argument")
	ErrBadRoute        = errors.New("bad route")
)
