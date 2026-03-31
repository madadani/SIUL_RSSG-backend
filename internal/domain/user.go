package domain

import "time"

// User represents a system user (PEP, PPTK, PPKOM, PP)
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string    `gorm:"size:150;not null" json:"nama"`
	Username     string    `gorm:"size:50;unique;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         string    `gorm:"type:role_enum;not null" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// KewenanganPegawai maps user access to specific budget categories
type KewenanganPegawai struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            uint            `gorm:"not null" json:"user_id"`
	KategoriBelanjaID uint            `gorm:"not null" json:"kategori_belanja_id"`
	CreatedAt         time.Time       `json:"created_at"`
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	KategoriBelanja   KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori_belanja"`
}

func (KewenanganPegawai) TableName() string { return "kewenangan_pegawai" }
