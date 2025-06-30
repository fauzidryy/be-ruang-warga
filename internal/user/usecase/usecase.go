package usecase

import (
	"be-ruang-warga/internal/user/domain"

	"gorm.io/gorm"
)

type UserUsecase interface {
	FindOrCreateUser(email, name string) (*domain.User, error)
	SubmitAdminReq(req *domain.AdminRequest) error
}

type userUsecase struct {
	db *gorm.DB
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{db: db}
}

func (u *userUsecase) FindOrCreateUser(email, name string) (*domain.User, error) {
	var user domain.User
	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		user = domain.User{
			Email: email,
			Name:  name,
			Role:  "warga",
		}
		if err := u.db.Create(&user).Error; err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (u *userUsecase) SubmitAdminReq(req *domain.AdminRequest) error {
	req.Status = "pending"
	return u.db.Create(&req).Error
}
