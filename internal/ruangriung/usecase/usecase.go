package usecase

import (
	"be-ruang-warga/internal/ruangriung/domain"

	"gorm.io/gorm"
)

type RuangRiungUsecase interface {
	Create(data *domain.RuangRiung) error
	GetAll() ([]domain.RuangRiung, error)
}

type ruangRiungUsecase struct {
	db *gorm.DB
}

func NewRuangRiungUsecase(db *gorm.DB) RuangRiungUsecase {
	return &ruangRiungUsecase{db: db}
}

func (r *ruangRiungUsecase) Create(data *domain.RuangRiung) error {
	return r.db.Create(data).Error
}

func (r *ruangRiungUsecase) GetAll() ([]domain.RuangRiung, error) {
	var schedules []domain.RuangRiung
	err := r.db.Find(&schedules).Error
	return schedules, err
}
