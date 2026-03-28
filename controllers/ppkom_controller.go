package controllers

import (
	"fmt"
	"net/http"

	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
)

// Dapatkan daftar usulan yang masuk ke PPKOM
func GetUsulanPPKOM(c *gin.Context) {
	// userIdStr, _ := c.Get("user_id")
	// userID := uint(userIdStr.(float64))
	
	var listUsulan []models.Usulan

	// Secara bisnis: PPKOM melihat usulan yang sudah direkomendasikan PPTK
	// dan usulan usulan yg dia tolak atau selesaikan
	
	result := config.DB.
		Preload("KategoriBelanja").
		Where("status_kode IN (?)", []string{"DIDISPOSISI_PPKOM", "DIKEMBALIKAN_KE_PPTK", "DIDISPOSISI_PP", "REALISASI_SELESAI"}). 
		Order("created_at desc").
		Find(&listUsulan)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memuat usulan PPKOM"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil memuat usulan PPKOM",
		"data":    listUsulan,
	})
}

// Request untuk PPKOM menerima / Disposisi ke PP (Operator Pengadaan)
type SetujuiPPKOMRequest struct {
	Catatan string `json:"catatan"`
}

func SetujuiDanDisposisiKePP(c *gin.Context) {
	id := c.Param("id")
	var input SetujuiPPKOMRequest
	c.ShouldBindJSON(&input)

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	var usulan models.Usulan
	if err := config.DB.First(&usulan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Usulan tidak ditemukan"})
		return
	}

	if usulan.StatusKode != "DIDISPOSISI_PPKOM" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Usulan ini belum dalam antrean persetujuan PPKOM"})
		return
	}

	tx := config.DB.Begin()
	usulan.StatusKode = "DIDISPOSISI_PP"
	if err := tx.Save(&usulan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan status"})
		return
	}

	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("DIDISPOSISI_PPKOM"),
		StatusAkhir:  "DIDISPOSISI_PP",
		CatatanAlasan: fmt.Sprintf("Disetujui oleh PPKOM & diteruskan ke Pejabat Pengadaan (PP). %s", input.Catatan),
	}
	tx.Create(&history)
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Usulan berhasil disetujui, masuk ke antrean Pengadaan"})
}

type TolakPPKOMRequest struct {
	AlasanTolak string `json:"alasan_tolak" binding:"required"`
}

func TolakDanKembalikanKePPTK(c *gin.Context) {
	id := c.Param("id")
	var input TolakPPKOMRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Alasan penolakan harus diisi"})
		return
	}

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	var usulan models.Usulan
	if err := config.DB.First(&usulan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Usulan tidak ditemukan"})
		return
	}

	tx := config.DB.Begin()
	// Gunakan status DIKEMBALIKAN_KE_PPTK (bisa ditambahkan ke ENUM di DB nanti kalau belum ada, sementara pakai fallback logik jika tidak ada)
	// Namun di schema sql sebelumnya belum ada 'DIKEMBALIKAN_KE_PPTK'.
	// Karena di schema.sql hanya ada: 'MENUNGGU_PEP', 'DIDISPOSISI_PPTK', 'DIKEMBALIKAN_KE_PEP', 'GESER_TAHUN_DEPAN', 'DIDISPOSISI_PPKOM', 'DIDISPOSISI_PP', 'REALISASI_SELESAI'.
	// Kita akan kembalikan statusnya menjadi 'DIDISPOSISI_PPTK' lagi atau buat status khusus "TOLAK_PPKOM".
	// Mari kita asumsikan dikembalikan menjadi DIDISPOSISI_PPTK dengan catatan ditolak.
	
	usulan.StatusKode = "DIDISPOSISI_PPTK" 
	tx.Save(&usulan)

	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("DIDISPOSISI_PPKOM"),
		StatusAkhir:  "DIDISPOSISI_PPTK",
		CatatanAlasan: fmt.Sprintf("Ditolak/Dikembalikan oleh PPKOM karena: %s", input.AlasanTolak),
	}
	tx.Create(&history)
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Usulan dikembalikan ke PPTK"})
}
