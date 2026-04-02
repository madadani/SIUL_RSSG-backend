package pptk

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

func (u *Usecase) GetUsulanPPTK(userID uint) ([]domain.Usulan, error) {
	return u.Repo.FindUsulanByPPTK(userID)
}

func (u *Usecase) DisposisiKePPKOM(id string, ppkomUserID uint, catatan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}
	if usulan.StatusKode != "DIDISPOSISI_PPTK" {
		return fmt.Errorf("usulan ini belum dalam antrean Anda")
	}

	usulan.StatusKode = "DIDISPOSISI_PPKOM"
	usulan.PPKOMUserId = &ppkomUserID
	usulan.CatatanPPTK = catatan
	now := time.Now()
	usulan.DisposisiPPTKAt = &now
	statusAwal := "DIDISPOSISI_PPTK"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIDISPOSISI_PPKOM",
		CatatanAlasan: fmt.Sprintf("Anggaran tersedia & diteruskan ke PPKOM. %s", catatan),
	}
	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) ReturnKePEP(id string, alasan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}

	usulan.StatusKode = "DIKEMBALIKAN_KE_PEP"
	usulan.AlasanReturn = alasan
	now := time.Now()
	usulan.ReturnAt = &now
	statusAwal := "DIDISPOSISI_PPTK"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIKEMBALIKAN_KE_PEP",
		CatatanAlasan: fmt.Sprintf("Dikembalikan ke PEP karena: %s", alasan),
	}
	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) AcknowledgeReturn(id string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}

	if usulan.StatusKode != "DIKEMBALIKAN_KE_PPTK" {
		return fmt.Errorf("hanya usulan yang dikembalikan yang bisa diketahui")
	}

	usulan.IsReturnDiketahui = true

	statusAwal := "DIKEMBALIKAN_KE_PPTK"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIKEMBALIKAN_KE_PPTK",
		CatatanAlasan: "PPTK telah mengetahui pengembalian usulan",
	}

	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) GetPPKOMUsers() ([]domain.User, error) {
	return u.Repo.FindPPKOMUsers()
}
