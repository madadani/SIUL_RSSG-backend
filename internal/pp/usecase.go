package pp

import (
	"fmt"

	"siul-pbj-api/internal/domain"
)

type Usecase struct {
	Repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{Repo: repo}
}

func (u *Usecase) GetUsulanPP(userID uint) ([]domain.Usulan, error) {
	return u.Repo.FindUsulanPP(userID)
}

func (u *Usecase) RealisasiUsulan(id string, namaVendor string, nomorKontrak string, hargaFinal float64, kategoriId *uint, catatan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}
	if usulan.StatusKode != "DIDISPOSISI_PP" {
		return fmt.Errorf("usulan ini belum dalam antrean PP")
	}

	usulan.StatusKode = "REALISASI_SELESAI"
	if kategoriId != nil && *kategoriId > 0 {
		usulan.KategoriBelanjaID = *kategoriId
	}

	realisasi := &domain.RealisasiLaporan{
		UsulanID:     usulan.ID,
		PPID:         actorID,
		NamaVendor:   namaVendor,
		NomorKontrak: nomorKontrak,
		HargaFinal:   hargaFinal,
	}

	statusAwal := "DIDISPOSISI_PP"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "REALISASI_SELESAI",
		CatatanAlasan: fmt.Sprintf("Realisasi selesai. Vendor: %s, Kontrak: %s, Harga: %.2f. %s", namaVendor, nomorKontrak, hargaFinal, catatan),
	}

	return u.Repo.RealisasiTransaction(usulan, realisasi, history)
}
