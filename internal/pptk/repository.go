package pptk

import (
	"siul-pbj-api/internal/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindUsulanByPPTK(userID uint) ([]domain.Usulan, error) {
	var list []domain.Usulan
	err := r.DB.
		Preload("KategoriBelanja").
		Where("pptk_user_id = ?", userID).
		Where("status_kode != 'MENUNGGU_PEP'").
		Order("created_at desc").
		Find(&list).Error
	return list, err
}

func (r *Repository) FindUsulanByID(id string) (*domain.Usulan, error) {
	var usulan domain.Usulan
	err := r.DB.First(&usulan, id).Error
	return &usulan, err
}

func (r *Repository) UpdateUsulanAndCreateHistory(usulan *domain.Usulan, history *domain.RiwayatUsulan) error {
	tx := r.DB.Begin()
	if err := tx.Save(usulan).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(history).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *Repository) FindPPKOMUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Where("role = ?", "PPKOM").Find(&users).Error
	return users, err
}
