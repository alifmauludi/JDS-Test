package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByUsername(username string) (User, error)
	FindByID(user_id int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Omit("PlainPassword").Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(UserID int) (User, error) {
	var user User
	err := r.db.Where("user_id = ?", UserID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
