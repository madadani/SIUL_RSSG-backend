package pep

import (
	"fmt"

	"time"
	"siul-pbj-api/internal/domain"
)

// Usecase holds business logic for the PEP role
type Usecase struct {
	Repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{Repo: repo}
}

func (u *Usecase) GetAllUsulan() ([]domain.Usulan, error) {
	return u.Repo.FindAllUsulan()
}

func (u *Usecase) DisposisiKePPTK(id string, pptkUserID uint, catatan string, actorID uint) error {
	usulan, err := u.Repo.FindUsulanByID(id)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan")
	}

	if usulan.StatusKode != "MENUNGGU_PEP" {
		return fmt.Errorf("usulan ini bukan dalam status Menunggu PEP")
	}

	usulan.StatusKode = "DIDISPOSISI_PPTK"
	usulan.PPTKUserId = &pptkUserID
	usulan.CatatanPEP = catatan
	now := time.Now()
	usulan.DisposisiPEPAt = &now

	statusAwal := "MENUNGGU_PEP"
	history := &domain.RiwayatUsulan{
		UsulanID:      usulan.ID,
		ActorID:       &actorID,
		StatusAwal:    &statusAwal,
		StatusAkhir:   "DIDISPOSISI_PPTK",
		CatatanAlasan: fmt.Sprintf("Didisposisikan ke PPTK ID: %d. %s", pptkUserID, catatan),
	}

	return u.Repo.UpdateUsulanAndCreateHistory(usulan, history)
}

func (u *Usecase) GetPPTKUsers() ([]domain.User, error) {
	return u.Repo.FindPPTKUsers()
}

func (u *Usecase) GetAllUsers() ([]domain.User, error) {
	return u.Repo.FindAllUsers()
}

func (u *Usecase) GetDetailAnggaran() ([]domain.DetailAnggaran, error) {
	return u.Repo.FindAllDetailAnggaran()
}
