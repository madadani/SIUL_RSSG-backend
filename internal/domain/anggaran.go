package domain

import "time"

// KategoriBelanja represents budget categories (e.g., BBJ, BMPM)
type KategoriBelanja struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	KodeKategori string    `gorm:"size:50;unique;not null" json:"kode_kategori"`
	NamaKategori string    `gorm:"size:150;not null" json:"nama_kategori"`
	CreatedAt    time.Time `json:"created_at"`
}

func (KategoriBelanja) TableName() string { return "kategori_belanja" }

// MasterRincianBelanja represents the master data of budget line items
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

// DetailAnggaran represents the budget allocation details for a procurement item
type DetailAnggaran struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID uint      `gorm:"column:category_id;not null" json:"category_id"`
	Nama       string    `gorm:"size:255;not null" json:"nama"`
	Nominal    float64   `gorm:"type:numeric(15,2);not null" json:"nominal"`
	PPTKID     uint      `gorm:"column:pptk_id;not null" json:"pptk_id"`
	PPKOMID    *uint     `gorm:"column:ppkom_id" json:"ppkom_id"`
	PPID       *uint     `gorm:"column:pp_id" json:"pp_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category KategoriBelanja `gorm:"foreignKey:CategoryID" json:"kategori"`
	PPTK     User            `gorm:"foreignKey:PPTKID" json:"pptk_user"`
	PPKOM    User            `gorm:"foreignKey:PPKOMID" json:"ppkom_user"`
	PP       User            `gorm:"foreignKey:PPID" json:"pp_user"`
}

func (DetailAnggaran) TableName() string { return "detail_anggaran" }
