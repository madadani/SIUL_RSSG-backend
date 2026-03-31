package public

import (
	"net/http"
	"strconv"

	"siul-pbj-api/internal/domain"
	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for public endpoints
type Handler struct {
	UC *Usecase
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{UC: NewUsecase(db)}
}

func (h *Handler) SubmitUsulan(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		response.BadRequest(c, "Gagal memproses form data")
		return
	}

	namaPengusul := c.PostForm("nama_pengusul")
	noHP := c.PostForm("no_hp")
	unitRuangan := c.PostForm("unit_ruangan")
	namaUsulan := c.PostForm("nama_usulan")
	keterangan := c.PostForm("keterangan")
	jumlah, _ := strconv.Atoi(c.PostForm("jumlah"))
	satuan := c.PostForm("satuan")
	kegentingan := c.PostForm("kegentingan")
	tingkatKepentingan := c.PostForm("tingkat_kepentingan")

	if namaPengusul == "" || noHP == "" || namaUsulan == "" || jumlah == 0 || kegentingan == "" || tingkatKepentingan == "" {
		response.BadRequest(c, "Harap lengkapi field wajib (*)")
		return
	}

	kategoriBelanjaID := h.UC.MatchKategori(namaUsulan)
	fotoBarangPath := h.UC.SaveUploadedFile(c, "foto_barang")

	usulan := domain.Usulan{
		KodeTiket:          h.UC.GenerateTicketCode(),
		NamaPengusul:       namaPengusul,
		NoHP:               noHP,
		UnitRuangan:        unitRuangan,
		NamaUsulan:         namaUsulan,
		Keterangan:         keterangan,
		Jumlah:             jumlah,
		Satuan:             satuan,
		Kegentingan:        kegentingan,
		TingkatKepentingan: tingkatKepentingan,
		SumberUsulan:       "Unit Kerja",
		FotoBarang:         fotoBarangPath,
		KategoriBelanjaID:  kategoriBelanjaID,
		StatusKode:         "MENUNGGU_PEP",
	}

	if err := h.UC.CreateUsulanWithHistory(&usulan); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, "Usulan berhasil disubmit", gin.H{
		"nomor_tiket": usulan.KodeTiket,
	})
}

func (h *Handler) GetMasterKategori(c *gin.Context) {
	kategori, err := h.UC.GetAllKategori()
	if err != nil {
		response.InternalError(c, "Gagal memuat kategori")
		return
	}
	response.Success(c, "Berhasil memuat kategori", kategori)
}

func (h *Handler) GetMasterBarang(c *gin.Context) {
	barang, err := h.UC.GetAllMasterBarang()
	if err != nil {
		response.InternalError(c, "Gagal memuat master barang")
		return
	}
	response.Success(c, "Berhasil memuat master barang", barang)
}

func (h *Handler) GetDetailUsulanPublik(c *gin.Context) {
	nomorTiket := c.Param("nomor_tiket")
	usulan, riwayat, err := h.UC.GetUsulanByTiket(nomorTiket)
	if err != nil {
		response.NotFound(c, "Usulan tidak ditemukan. Periksa kembali Nomor Tiket Anda.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"usulan":  usulan,
			"riwayat": riwayat,
		},
	})
}

func (h *Handler) GetUsulanPublikList(c *gin.Context) {
	search := c.Query("search")
	page, limit := ParsePagination(c)

	usulanList, total, err := h.UC.GetUsulanPaginated(search, page, limit)
	if err != nil {
		response.InternalError(c, "Gagal memuat daftar usulan")
		return
	}

	response.SuccessWithMeta(c, "Berhasil memuat daftar usulan", usulanList, gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
