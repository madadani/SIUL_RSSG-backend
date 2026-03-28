package controllers

import (
	"fmt"
	"net/http"
	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
)

// Dapatkan semua daftar usulan yang masuk ke PEP
func GetUsulanMasuk(c *gin.Context) {
	var listUsulan []models.Usulan
	// Kita join dgn kategori jika mau melihat nama kategori
	result := config.DB.Preload("KategoriBelanja").Order("created_at desc").Find(&listUsulan)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memuat usulan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil memuat usulan",
		"data":    listUsulan,
	})
}

// Struct request untuk mendisposisi ke PPTK
type DisposisiRequest struct {
	PPTKUserId uint   `json:"pptk_user_id" binding:"required"`
	Catatan    string `json:"catatan"`
}

func DisposisiKePPTK(c *gin.Context) {
	var input DisposisiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Harap lengkapi tujuan PPTK"})
		return
	}

	id := c.Param("id")

	// Ambil data user yang sedang login (actor)
	actorIdStr, _ := c.Get("user_id")
	// Type assertion, jwt Float64
	actorID := uint(actorIdStr.(float64))

	// Cari usulan
	var usulan models.Usulan
	if err := config.DB.First(&usulan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Usulan tidak ditemukan"})
		return
	}

	if usulan.StatusKode != "MENUNGGU_PEP" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Usulan ini bukan dalam status Menunggu PEP"})
		return
	}

	// Update usulan di DB dan tambah riwayat
	tx := config.DB.Begin()

	// Update record
	usulan.StatusKode = "DIDISPOSISI_PPTK"
	usulan.PPTKUserId = &input.PPTKUserId
	if err := tx.Save(&usulan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal update usulan"})
		return
	}

	// Buat audit riwayat
	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("MENUNGGU_PEP"), // Workaround for pointer string
		StatusAkhir:  "DIDISPOSISI_PPTK",
		CatatanAlasan: fmt.Sprintf("Didisposisikan ke PPTK ID: %d. %s", input.PPTKUserId, input.Catatan),
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal membuat riwayat"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Usulan berhasil didisposisi ke PPTK",
	})
}

// Dapatkan daftar user PPTK untuk dropdown disposisi
func GetPPTKUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Where("role = ?", "pptk").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memuat data PPTK"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}
