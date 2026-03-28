package controllers

import (
	"fmt"
	"net/http"

	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
)

// Dapatkan daftar usulan yang sudah didisposisikan ke PPTK
func GetUsulanPPTK(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userID := uint(userIdStr.(float64))

	var listUsulan []models.Usulan

	// Secara bisnis: PPTK melihat usulan yang didisposisikan ke mereka (pptk_user_id)
	
	result := config.DB.
		Preload("KategoriBelanja").
		Where("pptk_user_id = ?", userID).
		Where("status_kode != 'MENUNGGU_PEP'"). 
		Order("created_at desc").
		Find(&listUsulan)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memuat usulan PPTK"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil memuat usulan PPTK",
		"data":    listUsulan,
	})
}

// Struct untuk Disposisi PPTK -> PPKOM
type DisposisiPPKOMRequest struct {
	Catatan string `json:"catatan"`
}

func DisposisiKePPKOM(c *gin.Context) {
	id := c.Param("id")
	var input DisposisiPPKOMRequest
	c.ShouldBindJSON(&input)

	actorIdStr, _ := c.Get("user_id")
	actorID := uint(actorIdStr.(float64))

	var usulan models.Usulan
	if err := config.DB.First(&usulan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Usulan tidak ditemukan"})
		return
	}

	if usulan.StatusKode != "DIDISPOSISI_PPTK" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Usulan ini belum dalam antrean Anda"})
		return
	}

	// Update usulan
	tx := config.DB.Begin()
	usulan.StatusKode = "DIDISPOSISI_PPKOM"
	if err := tx.Save(&usulan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan usulan"})
		return
	}

	// Riwayat
	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("DIDISPOSISI_PPTK"),
		StatusAkhir:  "DIDISPOSISI_PPKOM",
		CatatanAlasan: fmt.Sprintf("Anggaran tersedia & diteruskan ke PPKOM. %s", input.Catatan),
	}
	tx.Create(&history)
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Usulan berhasil didisposisi ke PPKOM"})
}

// Return ke PEP
type ReturnPEPRequest struct {
	AlasanReturn string `json:"alasan_return" binding:"required"`
}

func ReturnKePEP(c *gin.Context) {
	id := c.Param("id")
	var input ReturnPEPRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Alasan return harus diisi"})
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
	usulan.StatusKode = "DIKEMBALIKAN_KE_PEP"
	tx.Save(&usulan)

	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		ActorID:      &actorID,
		StatusAwal:   func(s string) *string { return &s }("DIDISPOSISI_PPTK"),
		StatusAkhir:  "DIKEMBALIKAN_KE_PEP",
		CatatanAlasan: fmt.Sprintf("Dikembalikan ke PEP karena: %s", input.AlasanReturn),
	}
	tx.Create(&history)
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Usulan dikembalikan ke PEP"})
}
