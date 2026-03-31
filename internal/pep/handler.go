package pep

import (
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
