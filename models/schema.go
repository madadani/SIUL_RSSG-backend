package models

import (
	"time"
)

// Users
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string    `gorm:"size:150;not null" json:"nama"`
	Username     string    `gorm:"size:50;unique;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         string    `gorm:"type:role_enum;not null" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Kategori & Kewenangan
type KategoriBelanja struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	KodeKategori string    `gorm:"size:50;unique;not null" json:"kode_kategori"`
	NamaKategori string    `gorm:"size:150;not null" json:"nama_kategori"`
	CreatedAt    time.Time `json:"created_at"`
}

func (KategoriBelanja) TableName() string { return "kategori_belanja" }

type KewenanganPegawai struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            uint            `gorm:"not null" json:"user_id"`
	KategoriBelanjaID uint            `gorm:"not null" json:"kategori_belanja_id"`
	CreatedAt         time.Time       `json:"created_at"`
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	KategoriBelanja   KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori_belanja"`
}

func (KewenanganPegawai) TableName() string { return "kewenangan_pegawai" }

// Master Data
type MasterRincianBelanja struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	KategoriBelanjaID uint            `gorm:"not null" json:"kategori_belanja_id"`
	NamaBarang        string          `gorm:"size:255;not null" json:"nama_barang"`
	Spesifikasi       string          `gorm:"type:text" json:"spesifikasi"`
	Satuan            string          `gorm:"size:50" json:"satuan"`
	HargaReferensi    float64         `gorm:"type:numeric(15,2);not null" json:"harga_referensi"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	KategoriBelanja   KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori"`
}

func (MasterRincianBelanja) TableName() string { return "master_rincian_belanja" }

// Anggaran
type Anggaran struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	TahunAnggaran     int             `gorm:"not null" json:"tahun_anggaran"`
	KategoriBelanjaID uint            `gorm:"not null" json:"kategori_belanja_id"`
	PaguAwal          float64         `gorm:"type:numeric(15,2);not null" json:"pagu_awal"`
	SisaAnggaran      float64         `gorm:"type:numeric(15,2);not null" json:"sisa_anggaran"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	KategoriBelanja   KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori"`
}

func (Anggaran) TableName() string { return "anggaran" }

// Usulan
type Usulan struct {
	ID                 uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	KodeTiket          string          `gorm:"size:50;unique;not null" json:"kode_tiket"`
	NamaUsulan         string          `gorm:"size:200;not null" json:"nama_usulan"`
	Keterangan         string          `gorm:"type:text" json:"keterangan"`
	KategoriBelanjaID  uint            `gorm:"not null" json:"kategori_belanja_id"`
	TingkatKepentingan string          `gorm:"type:tingkat_kepentingan_enum;not null" json:"tingkat_kepentingan"`
	SumberUsulan       string          `gorm:"type:sumber_usulan_enum;not null" json:"sumber_usulan"`
	StatusKode         string          `gorm:"type:status_usulan_enum;default:'MENUNGGU_PEP';not null" json:"status_kode"`
	PPTKUserId         *uint           `json:"pptk_user_id"` // Assigned PPTK
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	KategoriBelanja    KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori"`
	PPTKUser           User            `gorm:"foreignKey:PPTKUserId" json:"pptk_user"`
}

func (Usulan) TableName() string { return "usulan" }

// Riwayat Usulan
type RiwayatUsulan struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UsulanID      uint      `gorm:"not null" json:"usulan_id"`
	ActorID       *uint     `json:"actor_id"` // Nullable if from public
	StatusAwal    *string   `gorm:"type:status_usulan_enum" json:"status_awal"`
	StatusAkhir   string    `gorm:"type:status_usulan_enum;not null" json:"status_akhir"`
	CatatanAlasan string    `gorm:"type:text" json:"catatan_alasan"`
	CreatedAt     time.Time `json:"created_at"`
}

func (RiwayatUsulan) TableName() string { return "riwayat_usulan" }

// Realisasi Laporan
type RealisasiLaporan struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UsulanID        uint      `gorm:"not null;unique" json:"usulan_id"`
	PPID            uint      `gorm:"column:pp_id;not null" json:"pp_id"`
	NamaVendor      string    `gorm:"size:200;not null" json:"nama_vendor"`
	NomorKontrak    string    `gorm:"size:100" json:"nomor_kontrak"`
	HargaFinal      float64   `gorm:"type:numeric(15,2);not null" json:"harga_final"`
	DokumenLaporan  string    `gorm:"size:255" json:"dokumen_laporan"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Usulan          Usulan    `gorm:"foreignKey:UsulanID" json:"usulan"`
	PP              User      `gorm:"foreignKey:PPID" json:"pp"`
}

func (RealisasiLaporan) TableName() string { return "realisasi_laporan" }
