package repository

import (
	"cbt-api/entity"
	"gorm.io/gorm"
)

// JawabanLatihanRepository is a contract for jawaban latihan repository
type JawabanLatihanRepository interface {
	CreateJawabanLatihan(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error)
	CreateBatchJawabanLatihan(jawabanList []entity.JawabanSiswa, idLatihan uint64) ([]entity.JawabanSiswa, error)
	GetJawabanLatihanByID(id uint64) (entity.JawabanSiswa, error)
	GetJawabanLatihanBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error)
	GetJawabanLatihanByLatihanID(latihanID uint64) ([]entity.JawabanSiswa, error)
}

// SoalLatihanRepository is a contract for soal latihan repository
type SoalLatihanRepository interface {
	GetSoalLatihanByID(latihanID uint64) ([]entity.Soal, error)
	GetSoalWithJawabanByLatihanID(latihanID uint64) (map[uint64][]entity.JawabanSoal, error)
}

type jawabanLatihanRepository struct {
	db *gorm.DB
}

type soalLatihanRepository struct {
	db *gorm.DB
}

// NewJawabanLatihanRepository creates a new instance of JawabanLatihanRepository
func NewJawabanLatihanRepository(db *gorm.DB) JawabanLatihanRepository {
	return &jawabanLatihanRepository{
		db: db,
	}
}

// NewSoalLatihanRepository creates a new instance of SoalLatihanRepository
func NewSoalLatihanRepository(db *gorm.DB) SoalLatihanRepository {
	return &soalLatihanRepository{
		db: db,
	}
}

// CreateBatchJawabanLatihan creates multiple jawaban for latihan at once
func (r *jawabanLatihanRepository) CreateBatchJawabanLatihan(jawabanList []entity.JawabanSiswa, idLatihan uint64) ([]entity.JawabanSiswa, error) {
	// Begin transaction
	tx := r.db.Begin()

	// Check for errors in transaction
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create all jawaban records
	err := tx.Create(&jawabanList).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return jawabanList, nil
}

func (r *jawabanLatihanRepository) CreateJawabanLatihan(jawabanSiswa entity.JawabanSiswa) (entity.JawabanSiswa, error) {
	err := r.db.Create(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanLatihanRepository) GetJawabanLatihanByID(id uint64) (entity.JawabanSiswa, error) {
	var jawabanSiswa entity.JawabanSiswa
	err := r.db.Preload("Soal").Preload("Siswa").Preload("JawabanSoal").Where("id_jawaban_siswa = ?", id).Take(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanLatihanRepository) GetJawabanLatihanBySiswaID(siswaID uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswa []entity.JawabanSiswa
	err := r.db.Joins("JOIN soal ON jawaban_siswa.id_soal = soal.id_soal").
		Joins("JOIN latihan_soal ON soal.id_soal = latihan_soal.id_soal").
		Preload("Soal").Preload("Siswa").Preload("JawabanSoal").
		Where("jawaban_siswa.id_siswa = ?", siswaID).
		Find(&jawabanSiswa).Error
	return jawabanSiswa, err
}

func (r *jawabanLatihanRepository) GetJawabanLatihanByLatihanID(latihanID uint64) ([]entity.JawabanSiswa, error) {
	var jawabanSiswa []entity.JawabanSiswa
	err := r.db.Joins("JOIN soal ON jawaban_siswa.id_soal = soal.id_soal").
		Joins("JOIN latihan_soal ON soal.id_soal = latihan_soal.id_soal").
		Preload("Soal").Preload("Siswa").Preload("JawabanSoal").
		Where("latihan_soal.id_latihan = ?", latihanID).
		Find(&jawabanSiswa).Error
	return jawabanSiswa, err
}

// SoalLatihanRepository functions to handle soal-related operations

func (r *soalLatihanRepository) GetSoalLatihanByID(latihanID uint64) ([]entity.Soal, error) {
	var soalList []entity.Soal
	err := r.db.Joins("JOIN latihan_soal ON soal.id_soal = latihan_soal.id_soal").
		Where("latihan_soal.id_latihan = ?", latihanID).
		Find(&soalList).Error
	return soalList, err
}

func (r *soalLatihanRepository) GetSoalWithJawabanByLatihanID(latihanID uint64) (map[uint64][]entity.JawabanSoal, error) {
	// Get all soal for this latihan
	var soalList []entity.Soal
	err := r.db.Joins("JOIN latihan_soal ON soal.id_soal = latihan_soal.id_soal").
		Where("latihan_soal.id_latihan = ?", latihanID).
		Find(&soalList).Error
	if err != nil {
		return nil, err
	}

	// Create a map to store soal ID to jawaban list
	result := make(map[uint64][]entity.JawabanSoal)

	// For each soal, get its jawaban options
	for _, soal := range soalList {
		var jawabanList []entity.JawabanSoal
		err := r.db.Where("id_soal = ?", soal.IdSoal).Find(&jawabanList).Error
		if err != nil {
			continue
		}

		result[soal.IdSoal] = jawabanList
	}

	return result, nil
}
