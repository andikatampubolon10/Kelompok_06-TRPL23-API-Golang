package entity

import (
	"time"
)

type Nilai struct {
	IdNilai     uint64    `gorm:"primary_key;autoIncrement" json:"id_nilai"`
	NilaiTotal  float64   `gorm:"type:decimal(5,0);not null" json:"nilai_total"`
	IdKursus    uint64    `gorm:"type:bigint(20);not null" json:"id_kursus"`
	IdSiswa     uint64    `gorm:"type:bigint(20);not null" json:"id_siswa"`
	IdTipeNilai uint64    `gorm:"type:bigint(20);not null" json:"id_tipe_nilai"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`

	// Relationships
	Kursus    Kursus    `gorm:"foreignKey:IdKursus;references:IdKursus" json:"kursus"`
	Siswa     Siswa     `gorm:"foreignKey:IdSiswa;references:IdSiswa" json:"siswa"`
	TipeNilai TipeNilai `gorm:"foreignKey:IdTipeNilai;references:IdTipeNilai" json:"tipe_nilai"`
}

func (Nilai) TableName() string {
    return "nilai" // Nama tabel di database
}
