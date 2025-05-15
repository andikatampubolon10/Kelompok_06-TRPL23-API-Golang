// repository/jawaban_siswa_repository.go
package repository

import (
	"cbt-api/entity"
	"gorm.io/gorm"
)

// JawabanSiswaRepository is a contract for jawaban siswa repository
type JawabanSiswaRepository interface {
	CreateJawabanSiswa(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error)
	CreateBatchJawabanSiswa(jawabanList []entity.JawabanSiswa) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByID(id uint64) (entity.JawabanSiswa, error)
	GetJawabanSiswaBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByUjianID(ujianID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanSiswaByUjianIDAndSiswaID(ujianID uint64, siswaID uint64) ([]entity.JawabanSiswa, error)
}

type jawabanSiswaRepository struct {
	db *gorm.DB
}

// NewJawabanSiswaRepository creates a new instance of JawabanSiswaRepository
func NewJawabanSiswaRepository(db *gorm.DB) JawabanSiswaRepository {
	return &jawabanSiswaRepository{
		db: db,
	}
}

func (r *jawabanSiswaRepository) CreateJawabanSiswa(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error) {
	err := r.db.Create(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanSiswaRepository) CreateBatchJawabanSiswa(jawabanList []entity.JawabanSiswa) ([]entity.JawabanSiswa, error) {
	err := r.db.Create(&jawabanList).Error
	return jawabanList, err
}

func (r *jawabanSiswaRepository) GetJawabanSiswaByID(id uint64) (entity.JawabanSiswa, error) {
	var jawabanSiswa entity.JawabanSiswa
	err := r.db.Preload("Soal").Preload("Siswa").Preload("JawabanSoal").Where("id_jawaban_siswa = ?", id).Take(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanSiswaRepository) GetJawabanSiswaBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswa []entity.JawabanSiswa
	err := r.db.Preload("Soal").Preload("Siswa").Preload("JawabanSoal").Where("id_siswa = ?", siswaID).Find(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanSiswaRepository) GetJawabanSiswaByUjianID(ujianID uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswa []entity.JawabanSiswa
	err := r.db.Joins("JOIN soal ON jawaban_siswa.id_soal = soal.id_soal").
		Preload("Soal").Preload("Siswa").Preload("JawabanSoal").
		Where("soal.id_ujian = ?", ujianID).
		Find(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanSiswaRepository) GetJawabanSiswaByUjianIDAndSiswaID(ujianID uint64, siswaID uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswa []entity.JawabanSiswa
	err := r.db.Joins("JOIN soal ON jawaban_siswa.id_soal = soal.id_soal").
		Joins("JOIN jawaban_soal ON jawaban_siswa.id_jawaban_soal = jawaban_soal.id_jawaban_soal").
		Where("soal.id_ujian = ? AND jawaban_siswa.id_siswa = ?", ujianID, siswaID).
		Preload("Soal").
		Preload("JawabanSoal").
		Find(&jawabanSiswa).Error
	return jawabanSiswa, err
}