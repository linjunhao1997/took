package file

import "gorm.io/gorm"

type Repository interface {
	Save(file *File) error
	FindOne(id int) (*File, error)
	Find(id ...int) ([]*File, error)
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{
		db,
	}
}

func (r *fileRepository) Save(file *File) error {
	return r.db.Save(file).Error
}

func (r *fileRepository) FindOne(id int) (*File, error) {
	file := &File{}
	if err := r.db.First(file, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (r *fileRepository) Find(id ...int) ([]*File, error) {
	files := make([]*File, 0)
	if err := r.db.Find(&files, "id in (?)", id).Error; err != nil {
		return files, err
	}
	return files, nil
}
