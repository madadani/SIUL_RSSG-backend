package public

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	"siul-pbj-api/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Usecase holds business logic for public-facing operations
type Usecase struct {
	DB *gorm.DB
}

func NewUsecase(db *gorm.DB) *Usecase {
	return &Usecase{DB: db}
}

// GenerateTicketCode creates a unique ticket number like USL-2026-64X89
func (u *Usecase) GenerateTicketCode() string {
	rngSrc := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSrc)
	return fmt.Sprintf("USL-%d-%05d", time.Now().Year(), rng.Intn(99999))
}

// MatchKategori tries to auto-assign a budget category based on the item name
func (u *Usecase) MatchKategori(namaUsulan string) uint {
	var matchedItem domain.MasterRincianBelanja
	kategoriID := uint(1) // Default fallback
	if err := u.DB.Where("nama_barang ILIKE ?", "%"+namaUsulan+"%").First(&matchedItem).Error; err == nil {
		kategoriID = matchedItem.KategoriBelanjaID
	}
	return kategoriID
}

// SaveUploadedFile handles optional file upload and returns the saved path
func (u *Usecase) SaveUploadedFile(c *gin.Context, fieldName string) string {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return ""
	}
	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	uploadPath := "./uploads/" + filename
	if err := c.SaveUploadedFile(file, uploadPath); err == nil {
		return "/uploads/" + filename
	}
	return ""
}

// CreateUsulanWithHistory creates a new Usulan and its initial RiwayatUsulan in a transaction
func (u *Usecase) CreateUsulanWithHistory(usulan *domain.Usulan) error {
	tx := u.DB.Begin()

	if err := tx.Create(usulan).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal menyimpan usulan: %w", err)
	}

	history := domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		StatusAkhir:   "MENUNGGU_PEP",
		CatatanAlasan: "Usulan baru masuk dari portal Publik",
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal menyimpan riwayat awal: %w", err)
	}

	return tx.Commit().Error
}

// GetAllKategori returns all budget categories
func (u *Usecase) GetAllKategori() ([]domain.KategoriBelanja, error) {
	var kategori []domain.KategoriBelanja
	err := u.DB.Order("id").Find(&kategori).Error
	return kategori, err
}

// GetAllMasterBarang returns all master items with their categories
func (u *Usecase) GetAllMasterBarang() ([]domain.MasterRincianBelanja, error) {
	var barang []domain.MasterRincianBelanja
	err := u.DB.Preload("KategoriBelanja").Order("nama_barang").Find(&barang).Error
	return barang, err
}

// GetUsulanByTiket fetches a single usulan by ticket code with its history
func (u *Usecase) GetUsulanByTiket(nomorTiket string) (*domain.Usulan, []domain.RiwayatUsulan, error) {
	var usulan domain.Usulan
	if err := u.DB.Preload("KategoriBelanja").Where("kode_tiket = ?", nomorTiket).First(&usulan).Error; err != nil {
		return nil, nil, err
	}

	var riwayat []domain.RiwayatUsulan
	u.DB.Where("usulan_id = ?", usulan.ID).Order("created_at desc").Find(&riwayat)

	return &usulan, riwayat, nil
}

// GetUsulanPaginated returns a paginated list of usulan with optional search
func (u *Usecase) GetUsulanPaginated(search string, page, limit int) ([]domain.Usulan, int64, error) {
	query := u.DB.Model(&domain.Usulan{})

	if search != "" {
		query = query.Where("nama_usulan ILIKE ? OR nama_pengusul ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	query.Count(&total)

	offset := (page - 1) * limit
	var usulanList []domain.Usulan
	err := query.Preload("KategoriBelanja").Order("created_at desc").Offset(offset).Limit(limit).Find(&usulanList).Error

	return usulanList, total, err
}

// ParsePagination extracts page and limit from query params with defaults
func ParsePagination(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 1 {
		limit = 10
	}
	return page, limit
}
