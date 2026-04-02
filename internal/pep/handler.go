package pep

import (
	"siul-pbj-api/internal/domain"
	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for PEP endpoints
type Handler struct {
	UC *Usecase
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	uc := NewUsecase(repo)
	return &Handler{UC: uc}
}

type DisposisiRequest struct {
	PPTKUserId uint   `json:"pptk_user_id" binding:"required"`
	Catatan    string `json:"catatan"`
}

func (h *Handler) GetUsulanMasuk(c *gin.Context) {
	list, err := h.UC.GetAllUsulan()
	if err != nil {
		response.InternalError(c, "Gagal memuat usulan")
		return
	}
	response.Success(c, "Berhasil memuat usulan", list)
}

func (h *Handler) DisposisiKePPTK(c *gin.Context) {
	var input DisposisiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Harap lengkapi tujuan PPTK")
		return
	}

	id := c.Param("id")
	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.DisposisiKePPTK(id, input.PPTKUserId, input.Catatan, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, "Usulan berhasil didisposisi ke PPTK", nil)
}

func (h *Handler) GetPPTKUsers(c *gin.Context) {
	users, err := h.UC.GetPPTKUsers()
	if err != nil {
		response.InternalError(c, "Gagal memuat data PPTK")
		return
	}
	response.Success(c, "Berhasil memuat data PPTK", users)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.UC.GetAllUsers()
	if err != nil {
		response.InternalError(c, "Gagal memuat semua data user")
		return
	}
	response.Success(c, "Berhasil memuat semua user", users)
}

func (h *Handler) GetDetailAnggaran(c *gin.Context) {
	list, err := h.UC.GetDetailAnggaran()
	if err != nil {
		response.InternalError(c, "Gagal memuat detail anggaran")
		return
	}
	response.Success(c, "Berhasil memuat detail anggaran", list)
}

func (h *Handler) AcknowledgeReturn(c *gin.Context) {
	id := c.Param("id")
	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.AcknowledgeReturn(id, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, "Usulan return berhasil diketahui", nil)
}

// Master Kategori CRUD
func (h *Handler) CreateKategori(c *gin.Context) {
	var input domain.KategoriBelanja
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Data tidak valid")
		return
	}
	if err := h.UC.Repo.DB.Create(&input).Error; err != nil {
		response.InternalError(c, "Gagal menambah kategori")
		return
	}
	response.Created(c, "Kategori berhasil ditambahkan", input)
}

func (h *Handler) UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var input domain.KategoriBelanja
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Data tidak valid")
		return
	}
	if err := h.UC.Repo.DB.Model(&domain.KategoriBelanja{}).Where("id = ?", id).Updates(input).Error; err != nil {
		response.InternalError(c, "Gagal update kategori")
		return
	}
	response.Success(c, "Kategori berhasil diupdate", nil)
}

func (h *Handler) DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	if err := h.UC.Repo.DB.Delete(&domain.KategoriBelanja{}, id).Error; err != nil {
		response.InternalError(c, "Gagal menghapus kategori")
		return
	}
	response.Success(c, "Kategori berhasil dihapus", nil)
}

// Master Barang CRUD
func (h *Handler) CreateBarang(c *gin.Context) {
	var input domain.MasterRincianBelanja
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Data tidak valid")
		return
	}
	if err := h.UC.Repo.DB.Create(&input).Error; err != nil {
		response.InternalError(c, "Gagal menambah master barang")
		return
	}
	response.Created(c, "Barang berhasil ditambahkan", input)
}

func (h *Handler) UpdateBarang(c *gin.Context) {
	id := c.Param("id")
	var input domain.MasterRincianBelanja
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Data tidak valid")
		return
	}
	if err := h.UC.Repo.DB.Model(&domain.MasterRincianBelanja{}).Where("id = ?", id).Updates(input).Error; err != nil {
		response.InternalError(c, "Gagal update master barang")
		return
	}
	response.Success(c, "Barang berhasil diupdate", nil)
}

func (h *Handler) DeleteBarang(c *gin.Context) {
	id := c.Param("id")
	if err := h.UC.Repo.DB.Delete(&domain.MasterRincianBelanja{}, id).Error; err != nil {
		response.InternalError(c, "Gagal menghapus master barang")
		return
	}
	response.Success(c, "Barang berhasil dihapus", nil)
}
