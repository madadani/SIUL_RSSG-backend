package domain

import "time"

// Usulan represents a procurement request submitted by a unit
type Usulan struct {
	ID                 uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	KodeTiket          string          `gorm:"size:50;unique;not null" json:"kode_tiket"`
	NamaPengusul       string          `gorm:"size:150;not null" json:"nama_pengusul"`
	NoHP               string          `gorm:"size:20;not null" json:"no_hp"`
	UnitRuangan        string          `gorm:"size:100" json:"unit_ruangan"`
	NamaUsulan         string          `gorm:"size:200;not null" json:"nama_usulan"`
	Keterangan         string          `gorm:"type:text" json:"keterangan"`
	Jumlah             int             `gorm:"not null" json:"jumlah"`
	Satuan             string          `gorm:"size:50" json:"satuan"`
	Kegentingan        string          `gorm:"size:50;not null" json:"kegentingan"`
	TingkatKepentingan string          `gorm:"type:tingkat_kepentingan_enum;not null" json:"tingkat_kepentingan"`
	SumberUsulan       string          `gorm:"type:sumber_usulan_enum;default:'Unit Kerja';not null" json:"sumber_usulan"`
	FotoBarang         string          `gorm:"size:255" json:"foto_barang"`
	KategoriBelanjaID  uint            `gorm:"not null" json:"kategori_belanja_id"`
	StatusKode         string          `gorm:"type:status_usulan_enum;default:'MENUNGGU_PEP';not null" json:"status_kode"`
	
	// History Tracking Fields
	CatatanPEP         string          `gorm:"type:text" json:"catatan_pep"`
	DisposisiPEPAt     *time.Time      `json:"disposisi_pep_at"`
	
	CatatanPPTK        string          `gorm:"type:text" json:"catatan_pptk"`
	DisposisiPPTKAt    *time.Time      `json:"disposisi_pptk_at"`
	AlasanReturn       string          `gorm:"type:text" json:"alasan_return"`
	ReturnAt           *time.Time      `json:"return_at"`
	IsReturnDiketahui  bool            `gorm:"default:false" json:"is_return_diketahui"`
	
	CatatanPPKOM       string          `gorm:"type:text" json:"catatan_ppkom"`
	SetujuPPKOMAt      *time.Time      `json:"setuju_ppkom_at"`
	AlasanTolak        string          `gorm:"type:text" json:"alasan_tolak"`
	RejectAt           *time.Time      `json:"reject_at"`

	PPTKUserId         *uint           `json:"pptk_user_id"`
	PPKOMUserId        *uint           `json:"ppkom_user_id"`
	PPUserId           *uint           `json:"pp_user_id"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	
	KategoriBelanja    KategoriBelanja `gorm:"foreignKey:KategoriBelanjaID" json:"kategori"`
	PPTKUser           User            `gorm:"foreignKey:PPTKUserId" json:"pptk_user"`
	PPKOMUser          User            `gorm:"foreignKey:PPKOMUserId" json:"ppkom_user"`
	PPUser             User            `gorm:"foreignKey:PPUserId" json:"pp_user"`
}

func (Usulan) TableName() string { return "usulan" }

// RiwayatUsulan tracks status changes of a procurement request
type RiwayatUsulan struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UsulanID      uint      `gorm:"not null" json:"usulan_id"`
	ActorID       *uint     `json:"actor_id"`
	StatusAwal    *string   `gorm:"type:status_usulan_enum" json:"status_awal"`
	StatusAkhir   string    `gorm:"type:status_usulan_enum;not null" json:"status_akhir"`
	CatatanAlasan string    `gorm:"type:text" json:"catatan_alasan"`
	CreatedAt     time.Time `json:"created_at"`
}

func (RiwayatUsulan) TableName() string { return "riwayat_usulan" }
