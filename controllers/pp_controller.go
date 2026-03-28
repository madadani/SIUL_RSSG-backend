package controllers

import (
	"fmt"
	"net/http"

	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
)

// Dapatkan daftar usulan yang masuk ke PP (Pejabat Pengadaan)
func GetUsulanPP(c *gin.Context) {
	var listUsulan []models.Usulan

	// PP melihat usulan yang sudah disetujui PPKOM (status DIDISPOSISI_PP) dan yang sudah selesai
	result := config.DB.
		Preload("KategoriBelanja").
		Where("status_kode IN (?)", []string{"DIDISPOSISI_PP", "REALISASI_SELESAI"}).
		Order("created_at desc").
		Find(&listUsulan)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memuat usulan PP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil memuat usulan PP",
		"data":    listUsulan,
	})
}

// Realisasi: PP menandai usulan sebagai selesai
type RealisasiRequest struct {
	NamaVendor   string  `json:"nama_vendor" binding:"required"`
	NomorKontrak string  `json:"nomor_kontrak"`
	HargaFinal   float64 `json:"harga_final" binding:"required"`
	Catatan      string  `json:"catatan"`
}

func RealisasiUsulan(c *gin.Context) {
	id := c.Param("id")
	var input RealisasiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Data realisasi tidak lengkap (vendor & harga wajib diisi)"})
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	var usulan models.Usulan
	if err := config.DB.First(&usulan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Usulan tidak ditemukan"})
		return
	}

	if usulan.StatusKode != "DIDISPOSISI_PP" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Usulan ini belum dalam antrean PP"})
		return
	}

	tx := config.DB.Begin()

	// Update status usulan
	usulan.StatusKode = "REALISASI_SELESAI"
	if err := tx.Save(&usulan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan status"})
		return
	}

	// Simpan laporan realisasi
	realisasi := models.RealisasiLaporan{
		UsulanID:     usulan.ID,
		PPID:         actorID,
		NamaVendor:   input.NamaVendor,
		NomorKontrak: input.NomorKontrak,
		HargaFinal:   input.HargaFinal,
	}
	if err := tx.Create(&realisasi).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan data realisasi"})
		return
	}

	// Riwayat
	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("DIDISPOSISI_PP"),
		StatusAkhir:  "REALISASI_SELESAI",
		CatatanAlasan: fmt.Sprintf("Realisasi selesai. Vendor: %s, Kontrak: %s, Harga: %.2f. %s", input.NamaVendor, input.NomorKontrak, input.HargaFinal, input.Catatan),
	}
	tx.Create(&history)
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Realisasi berhasil dicatat, proses pengadaan selesai!"})
}
