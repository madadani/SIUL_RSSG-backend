package domain

import "time"

// RealisasiLaporan records the final realization of a procurement by PP
type RealisasiLaporan struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UsulanID       uint      `gorm:"not null;unique" json:"usulan_id"`
	PPID           uint      `gorm:"column:pp_id;not null" json:"pp_id"`
	NamaVendor     string    `gorm:"size:200;not null" json:"nama_vendor"`
	NomorKontrak   string    `gorm:"size:100" json:"nomor_kontrak"`
	HargaFinal     float64   `gorm:"type:numeric(15,2);not null" json:"harga_final"`
	DokumenLaporan string    `gorm:"size:255" json:"dokumen_laporan"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Usulan         Usulan    `gorm:"foreignKey:UsulanID" json:"usulan"`
	PP             User      `gorm:"foreignKey:PPID" json:"pp"`
}

func (RealisasiLaporan) TableName() string { return "realisasi_laporan" }
