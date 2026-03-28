package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
)

type UsulanRequest struct {
	NamaUsulan         string `json:"nama_usulan" binding:"required"`
	KategoriBelanjaID  uint   `json:"kategori_belanja_id" binding:"required"`
	TingkatKepentingan string `json:"tingkat_kepentingan" binding:"required"`
	Keterangan         string `json:"keterangan"`
}

func SubmitUsulan(c *gin.Context) {
	var input UsulanRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Validasi gagal", "details": err.Error()})
		return
	}

	// Generate Nomor Tiket e.g., USL-2026-64X89
	rngSrc := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSrc)
	ticketCode := fmt.Sprintf("USL-%d-%05d", time.Now().Year(), rng.Intn(99999))

	usulan := models.Usulan{
		KodeTiket:          ticketCode,
		NamaUsulan:         input.NamaUsulan,
		KategoriBelanjaID:  input.KategoriBelanjaID,
		TingkatKepentingan: input.TingkatKepentingan,
		SumberUsulan:       "Publik",
		StatusKode:         "MENUNGGU_PEP",
		Keterangan:         input.Keterangan,
	}

	// Begin Transaction
	tx := config.DB.Begin()

	if err := tx.Create(&usulan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan usulan"})
		return
	}

	// Create initial History
	history := models.RiwayatUsulan{
		UsulanID:     usulan.ID,
		StatusAkhir:  "MENUNGGU_PEP",
		CatatanAlasan: "Usulan baru masuk dari portal Publik",
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan riwayat"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Usulan berhasil disubmit",
		"data": gin.H{
			"nomor_tiket": usulan.KodeTiket,
		},
	})
}
