package pptk

import (
	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	UC *Usecase
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	uc := NewUsecase(repo)
	return &Handler{UC: uc}
}

type DisposisiPPKOMRequest struct {
	PPKOMUserId uint   `json:"ppkom_user_id" binding:"required"`
	Catatan     string `json:"catatan"`
}

type ReturnPEPRequest struct {
	AlasanReturn string `json:"alasan_return" binding:"required"`
}

func (h *Handler) GetUsulanPPTK(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userID := uint(userIdStr.(float64))

	list, err := h.UC.GetUsulanPPTK(userID)
	if err != nil {
		response.InternalError(c, "Gagal memuat usulan PPTK")
		return
	}
	response.Success(c, "Berhasil memuat usulan PPTK", list)
}

func (h *Handler) DisposisiKePPKOM(c *gin.Context) {
	id := c.Param("id")
	var input DisposisiPPKOMRequest
	c.ShouldBindJSON(&input)

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.DisposisiKePPKOM(id, input.PPKOMUserId, input.Catatan, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, "Usulan berhasil didisposisi ke PPKOM", nil)
}

func (h *Handler) ReturnKePEP(c *gin.Context) {
	id := c.Param("id")
	var input ReturnPEPRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Alasan return harus diisi")
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.ReturnKePEP(id, input.AlasanReturn, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, "Usulan dikembalikan ke PEP", nil)
}
 
func (h *Handler) GetPPKOMUsers(c *gin.Context) {
	users, err := h.UC.GetPPKOMUsers()
	if err != nil {
		response.InternalError(c, "Gagal memuat data PPKOM")
		return
	}
	response.Success(c, "Berhasil memuat data PPKOM", users)
}
