package usecase

import (
	"be-ruang-warga/internal/ruangriung/domain"

	"gorm.io/gorm"
)

type RuangRiungUsecase interface {
	GetAll() ([]domain.RuangRiung, error)
	Create(data *domain.RuangRiung) error
	Update(id string, data *domain.RuangRiung) error
	Delete(id string) error
}

type ruangRiungUsecase struct {
	db *gorm.DB
}

func NewRuangRiungUsecase(db *gorm.DB) RuangRiungUsecase {
	return &ruangRiungUsecase{db: db}
}

func (r *ruangRiungUsecase) GetAll() ([]domain.RuangRiung, error) {
	var schedules []domain.RuangRiung
	err := r.db.Find(&schedules).Error
	return schedules, err
}

func (r *ruangRiungUsecase) Create(data *domain.RuangRiung) error {
	return r.db.Create(data).Error
}

func (r *ruangRiungUsecase) Update(id string, data *domain.RuangRiung) error {
	var existing domain.RuangRiung
	if err := r.db.First(&existing, id).Error; err != nil {
		return err
	}

	existing.Title = data.Title
	existing.Description = data.Description
	existing.ScheduleTime = data.ScheduleTime
	existing.Location = data.Location
	existing.PosterPath = data.PosterPath
	existing.Performers = data.Performers

	return r.db.Save(&existing).Error
}

func(r *ruangRiungUsecase) Delete(id string) error {
	var data domain.RuangRiung
	return r.db.Delete(&data, id).Error
}
