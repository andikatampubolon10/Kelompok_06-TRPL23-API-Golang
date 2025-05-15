package repository

import (
	"cbt-api/entity"
	"gorm.io/gorm"
)

// GetJawabanSiswaRepository is a contract for jawaban siswa database operations
type GetJawabanSiswaRepository interface {
	FindByIdSiswa(idSiswa uint64) ([]entity.JawabanSiswa, error)
	FindByIdSiswaAndIdUjian(idSiswa uint64, idUjian uint64) ([]entity.JawabanSiswa, error)
}

// SoalRepository is a contract for soal database operations
type SoalRepository interface {
	FindById(id uint64) (entity.Soal, error)
	FindByIdUjian(idUjian uint64) ([]entity.Soal, error)
}

// JawabanSoalRepository is a contract for jawaban soal database operations
type JawabanSoalRepository interface {
	FindById(id uint64) (entity.JawabanSoal, error)
}

// jawabanSiswaRepository is a struct that implements GetJawabanSiswaRepository interface
type getjawabanSiswaRepository struct {
	db *gorm.DB
}

// soalRepository is a struct that implements SoalRepository interface
type soalRepository struct {
	db *gorm.DB
}

// jawabanSoalRepository is a struct that implements JawabanSoalRepository interface
type jawabanSoalRepository struct {
	db *gorm.DB
}

// NewJawabanSiswaRepository creates a new instance of GetJawabanSiswaRepository
func NewGetJawabanSiswaRepository(db *gorm.DB) GetJawabanSiswaRepository {
	return &jawabanSiswaRepository{db}
}

// NewSoalRepository creates a new instance of SoalRepository
func NewSoalRepository(db *gorm.DB) SoalRepository {
	return &soalRepository{db}
}

// NewJawabanSoalRepository creates a new instance of JawabanSoalRepository
func NewJawabanSoalRepository(db *gorm.DB) JawabanSoalRepository {
	return &jawabanSoalRepository{db}
}

// FindByIdSiswa finds all jawaban siswa by id_siswa
func (r *jawabanSiswaRepository) FindByIdSiswa(idSiswa uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswaList []entity.JawabanSiswa
	err := r.db.Where("id_siswa = ?", idSiswa).Find(&jawabanSiswaList).Error
	return jawabanSiswaList, err
}

// FindByIdSiswaAndIdUjian finds all jawaban siswa by id_siswa and id_ujian through Soal table
func (r *jawabanSiswaRepository) FindByIdSiswaAndIdUjian(idSiswa uint64, idUjian uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswaList []entity.JawabanSiswa
	
	// Using join to get jawaban siswa that match both id_siswa and id_ujian
	err := r.db.Joins("JOIN soal ON jawaban_siswa.id_soal = soal.id_soal").
		Where("jawaban_siswa.id_siswa = ? AND soal.id_ujian = ?", idSiswa, idUjian).
		Find(&jawabanSiswaList).Error
	
	return jawabanSiswaList, err
}

// FindById finds soal by id
func (r *soalRepository) FindById(id uint64) (entity.Soal, error) {
	var soal entity.Soal
	err := r.db.Where("id_soal = ?", id).First(&soal).Error
	return soal, err
}

// FindByIdUjian finds all soal for a specific ujian
func (r *soalRepository) FindByIdUjian(idUjian uint64) ([]entity.Soal, error) {
	var soalList []entity.Soal
	err := r.db.Where("id_ujian = ?", idUjian).Find(&soalList).Error
	return soalList, err
}

// FindById finds jawaban soal by id
func (r *jawabanSoalRepository) FindById(id uint64) (entity.JawabanSoal, error) {
	var jawabanSoal entity.JawabanSoal
	err := r.db.Where("id_jawaban_soal = ?", id).First(&jawabanSoal).Error
	return jawabanSoal, err
}