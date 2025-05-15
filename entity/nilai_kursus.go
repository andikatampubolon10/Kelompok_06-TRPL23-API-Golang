package entity

import (
	"time"
)

type NilaiKursus struct {
	IdNilaiKursus     uint64    `gorm:"primary_key;autoIncrement" json:"id_nilai_kursus"`
	NilaiTipeUjian    float64   `gorm:"type:decimal(5,0);not null;default:0" json:"nilai_tipe_ujian"`
	IdKursus           uint64    `gorm:"type:bigint(20) unsigned" json:"id_kursus"`
	IdSiswa            uint64    `gorm:"type:bigint(20) unsigned" json:"id_siswa"`
	IdTipeUjian        uint64    `gorm:"type:bigint(20) unsigned" json:"id_tipe_ujian"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

	Kursus             Kursus      `gorm:"foreignkey:IdKursus;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kursus"`
	Siswa              Siswa       `gorm:"foreignkey:IdSiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"siswa"`
	TipeUjian          TipeUjian   `gorm:"foreignkey:IdTipeUjian;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_ujian"`
}

func (NilaiKursus) TableName() string {
    return "nilai_kursus" // Nama tabel di database
}
