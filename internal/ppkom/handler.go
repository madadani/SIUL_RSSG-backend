package ppkom

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

type SetujuiRequest struct {
	PPUserId uint   `json:"pp_user_id" binding:"required"`
	Catatan  string `json:"catatan"`
}

type TolakRequest struct {
	AlasanTolak string `json:"alasan_tolak" binding:"required"`
}

func (h *Handler) GetUsulanPPKOM(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userID := uint(userIdStr.(float64))

	list, err := h.UC.GetUsulanPPKOM(userID)
	if err != nil {
		response.InternalError(c, "Gagal memuat usulan PPKOM")
		return
	}
	response.Success(c, "Berhasil memuat usulan PPKOM", list)
}

func (h *Handler) SetujuiDanDisposisiKePP(c *gin.Context) {
	id := c.Param("id")
	var input SetujuiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Harap lengkapi tujuan Pejabat Pengadaan (PP)")
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.SetujuiDanDisposisiKePP(id, input.PPUserId, input.Catatan, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, "Usulan berhasil disetujui, masuk ke antrean Pengadaan", nil)
}

func (h *Handler) TolakDanKembalikanKePPTK(c *gin.Context) {
	id := c.Param("id")
	var input TolakRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Alasan penolakan harus diisi")
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.TolakDanKembalikanKePPTK(id, input.AlasanTolak, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, "Usulan dikembalikan ke PPTK", nil)
}

func (h *Handler) GetPPUsers(c *gin.Context) {
	users, err := h.UC.GetPPUsers()
	if err != nil {
		response.InternalError(c, "Gagal memuat data Pejabat Pengadaan")
		return
	}
	response.Success(c, "Berhasil memuat data Pejabat Pengadaan", users)
}
