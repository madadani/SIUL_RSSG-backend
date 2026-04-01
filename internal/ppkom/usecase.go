package ppkom

import (
	"fmt"
	"time"

	"siul-pbj-api/internal/domain"
)

type Usecase struct {
	Repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{Repo: repo}
}

func (u *Usecase) GetUsulanPPKOM(userID uint) ([]domain.Usulan, error) {
	return u.Repo.FindUsulanPPKOM(userID)
}

func (u *Usecase) SetujuiDanDisposisiKePP(id string, ppUserID uint, catatan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}
	if usulan.StatusKode != "DIDISPOSISI_PPKOM" {
		return fmt.Errorf("usulan ini belum dalam antrean persetujuan PPKOM")
	}

	usulan.StatusKode = "DIDISPOSISI_PP"
	usulan.PPUserId = &ppUserID
	usulan.CatatanPPKOM = catatan
	now := time.Now()
	usulan.SetujuPPKOMAt = &now

	statusAwal := "DIDISPOSISI_PPKOM"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIDISPOSISI_PP",
		CatatanAlasan: fmt.Sprintf("Disetujui oleh PPKOM & diteruskan ke Pejabat Pengadaan (PP) ID: %d. %s", ppUserID, catatan),
	}
	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) TolakDanKembalikanKePPTK(id string, alasan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}

	usulan.StatusKode = "DIDISPOSISI_PPTK"
	usulan.AlasanTolak = alasan
	now := time.Now()
	usulan.RejectAt = &now

	statusAwal := "DIDISPOSISI_PPKOM"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIDISPOSISI_PPTK",
		CatatanAlasan: fmt.Sprintf("Ditolak/Dikembalikan oleh PPKOM karena: %s", alasan),
	}
	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) GetPPUsers() ([]domain.User, error) {
	return u.Repo.FindPPUsers()
}
