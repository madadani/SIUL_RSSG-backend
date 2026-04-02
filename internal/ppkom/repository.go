package ppkom

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

func (r *Repository) FindUsulanPPKOM(userID uint) ([]domain.Usulan, error) {
	var list []domain.Usulan
	err := r.DB.Preload("KategoriBelanja").Preload("PPTKUser").Preload("PPKOMUser").Preload("PPUser").
		Where("status_kode = ?", "DIDISPOSISI_PPKOM").
		Where("ppkom_user_id = ?", userID).
		Order("created_at desc").
		Find(&list).Error
	return list, err
}

func (r *Repository) FindUsulanByID(id string) (*domain.Usulan, error) {
	var usulan domain.Usulan
	err := r.DB.Preload("KategoriBelanja").Preload("PPTKUser").Preload("PPKOMUser").Preload("PPUser").First(&usulan, id).Error
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

func (r *Repository) FindPPUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Where("role = ?", "pp").Find(&users).Error
	return users, err
}
