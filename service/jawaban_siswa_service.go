// service/jawaban_siswa_service.go
package service

import (
	"cbt-api/entity"
	"cbt-api/repository"
)

// JawabanSiswaService is a contract for jawaban siswa service
type JawabanSiswaService interface {
	CreateJawabanSiswa(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error)
	CreateBatchJawabanSiswa(jawabanList []entity.JawabanSiswa) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByID(id uint64) (entity.JawabanSiswa, error)
	GetJawabanSiswaBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByUjianID(ujianID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByUjianIDAndSiswaID(ujianID uint64, siswaID uint64) ([]entity.JawabanSiswa, error)
}

type jawabanSiswaService struct {
	jawabanSiswaRepository repository.JawabanSiswaRepository
}

// NewJawabanSiswaService creates a new instance of JawabanSiswaService
func NewJawabanSiswaService(jawabanSiswaRepo repository.JawabanSiswaRepository) JawabanSiswaService {
	return &jawabanSiswaService{
		jawabanSiswaRepository: jawabanSiswaRepo,
	}
}

func (s *jawabanSiswaService) CreateJawabanSiswa(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.CreateJawabanSiswa(jawabanSiswa)
}

func (s *jawabanSiswaService) CreateBatchJawabanSiswa(jawabanList []entity.JawabanSiswa) ([]entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.CreateBatchJawabanSiswa(jawabanList)
}

func (s *jawabanSiswaService) GetJawabanSiswaByID(id uint64) (entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.GetJawabanSiswaByID(id)
}

func (s *jawabanSiswaService) GetJawabanSiswaBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.GetJawabanSiswaBySiswaID(siswaID)
}

func (s *jawabanSiswaService) GetJawabanSiswaByUjianID(ujianID uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.GetJawabanSiswaByUjianID(ujianID)
}

func (s *jawabanSiswaService) GetJawabanSiswaByUjianIDAndSiswaID(ujianID uint64, siswaID uint64) ([]entity.JawabanSiswa, error) {
	return s.jawabanSiswaRepository.GetJawabanSiswaByUjianIDAndSiswaID(ujianID, siswaID)
}