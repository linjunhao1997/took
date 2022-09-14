package user

import "gorm.io/gorm"

type Repository interface {
	FindOne(id int) (*User, error)
	Find(id ...int) ([]*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) FindOne(id int) (*User, error) {
	user := &User{Id: id}
	if err := r.db.Find(user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Find(id ...int) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.Find(&users, "id in (?)", id).Error; err != nil {
		return users, err
	}
	return users, nil
}
