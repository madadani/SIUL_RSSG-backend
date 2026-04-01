package pp

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

type RealisasiRequest struct {
	NamaVendor   string  `json:"nama_vendor" binding:"required"`
	NomorKontrak string  `json:"nomor_kontrak"`
	HargaFinal   float64 `json:"harga_final" binding:"required"`
	Catatan      string  `json:"catatan"`
}

func (h *Handler) GetUsulanPP(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userID := uint(userIdStr.(float64))

	list, err := h.UC.GetUsulanPP(userID)
	if err != nil {
		response.InternalError(c, "Gagal memuat usulan PP")
		return
	}
	response.Success(c, "Berhasil memuat usulan PP", list)
}

func (h *Handler) RealisasiUsulan(c *gin.Context) {
	id := c.Param("id")
	var input RealisasiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Data realisasi tidak lengkap (vendor & harga wajib diisi)")
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	if err := h.UC.RealisasiUsulan(id, input.NamaVendor, input.NomorKontrak, input.HargaFinal, input.Catatan, actorID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, "Realisasi berhasil dicatat, proses pengadaan selesai!", nil)
}
