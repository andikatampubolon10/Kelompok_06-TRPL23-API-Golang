// service/jawaban_latihan_service.go
package service

import (
	"cbt-api/entity"
	"cbt-api/repository"
)

// JawabanLatihanService is a contract for jawaban latihan service
type JawabanLatihanService interface {
	CreateJawabanLatihan(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error)
	CreateBatchJawabanLatihan(jawabanList []entity.JawabanSiswa, idLatihan uint64) ([]entity.JawabanSiswa, error)
	GetJawabanLatihanByID(id uint64) (entity.JawabanSiswa, error)
	GetJawabanLatihanBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanLatihanByLatihanID(latihanID uint64) ([]entity.JawabanSiswa, error)
}

// SoalLatihanService is a contract for soal latihan service
type SoalLatihanService interface {
	GetSoalLatihanByID(latihanID uint64) ([]entity.Soal, error)
	GetSoalWithJawabanByLatihanID(latihanID uint64) (map[uint64][]entity.JawabanSoal, error)
}

type jawabanLatihanService struct {
	jawabanLatihanRepository repository.JawabanLatihanRepository
}

type soalLatihanService struct {
	soalLatihanRepository repository.SoalLatihanRepository
}

// NewJawabanLatihanService creates a new instance of JawabanLatihanService
func NewJawabanLatihanService(jawabanLatihanRepo repository.JawabanLatihanRepository) JawabanLatihanService {
	return &jawabanLatihanService{
		jawabanLatihanRepository: jawabanLatihanRepo,
	}
}

// NewSoalLatihanService creates a new instance of SoalLatihanService
func NewSoalLatihanService(soalLatihanRepo repository.SoalLatihanRepository) SoalLatihanService {
	return &soalLatihanService{
		soalLatihanRepository: soalLatihanRepo,
	}
}

func (s *jawabanLatihanService) CreateJawabanLatihan(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error) {
	return s.jawabanLatihanRepository.CreateJawabanLatihan(jawabanSiswa)
}

func (s *jawabanLatihanService) CreateBatchJawabanLatihan(jawabanList []entity.JawabanSiswa, idLatihan uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanLatihanRepository.CreateBatchJawabanLatihan(jawabanList, idLatihan)
}

func (s *jawabanLatihanService) GetJawabanLatihanByID(id uint64) (entity.JawabanSiswa, error) {
	return s.jawabanLatihanRepository.GetJawabanLatihanByID(id)
}

func (s *jawabanLatihanService) GetJawabanLatihanBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanLatihanRepository.GetJawabanLatihanBySiswaID(siswaID)
}

func (s *jawabanLatihanService) GetJawabanLatihanByLatihanID(latihanID uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanLatihanRepository.GetJawabanLatihanByLatihanID(latihanID)
}

func (s *soalLatihanService) GetSoalLatihanByID(latihanID uint64) ([]entity.Soal, error) {
	return s.soalLatihanRepository.GetSoalLatihanByID(latihanID)
}

func (s *soalLatihanService) GetSoalWithJawabanByLatihanID(latihanID uint64) (map[uint64][]entity.JawabanSoal, error) {
	return s.soalLatihanRepository.GetSoalWithJawabanByLatihanID(latihanID)
}