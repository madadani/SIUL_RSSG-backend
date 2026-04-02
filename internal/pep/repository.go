package pep

import (
	"siul-pbj-api/internal/domain"

	"gorm.io/gorm"
)

// Repository handles all PEP-related database operations
type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindAllUsulan() ([]domain.Usulan, error) {
	var list []domain.Usulan
	err := r.DB.Preload("KategoriBelanja").Preload("PPTKUser").Preload("PPKOMUser").Preload("PPUser").Order("created_at desc").Find(&list).Error
	return list, err
}

func (r *Repository) FindUsulanByID(id string) (*domain.Usulan, error) {
	var usulan domain.Usulan
	err := r.DB.Preload("KategoriBelanja").Preload("PPTKUser").Preload("PPKOMUser").Preload("PPUser").First(&usulan, id).Error
	return &usulan, err
}

func (r *Repository) FindPPTKUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Where("role = ?", "pptk").Find(&users).Error
	return users, err
}

func (r *Repository) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Order("role, nama").Find(&users).Error
	return users, err
}

func (r *Repository) FindAllDetailAnggaran() ([]domain.DetailAnggaran, error) {
	var list []domain.DetailAnggaran
	err := r.DB.Preload("Category").Preload("PPTK").Preload("PPKOM").Preload("PP").Order("created_at desc").Find(&list).Error
	return list, err
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
