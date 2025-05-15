package entity

import (
	"time"
)

type Soal struct {
	IdSoal    uint64    `gorm:"primary_key;autoIncrement" json:"id_soal"`
	Soal      string    `gorm:"type:varchar(255);not null" json:"soal"`
	Image     string    `gorm:"type:varchar(255);nullable" json:"image,omitempty"`
	ImageUrl       string    `gorm:"type:varchar(255);not null" json:"image_url"`
	NilaiPerSoal float64  `gorm:"type:decimal(5,2);nullable" json:"nilai_per_soal"`
	IdUjian   uint64    `gorm:"type:bigint(20) unsigned" json:"id_ujian"`
	IdTipeSoal uint64  `gorm:"type:bigint(20) unsigned" json:"id_tipe_soal"`
	IdLatihan uint64    `gorm:"type:bigint(20) unsigned" json:"id_latihan"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

	Ujian    Ujian   `gorm:"foreignkey:IdUjian;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"ujian"`
	TipeSoal TipeSoal `gorm:"foreignkey:IdTipeSoal;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_soal"`
	Latihan  Latihan `gorm:"foreignkey:IdLatihan;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"latihan"`
}

func (Soal) TableName() string {
    return "soal" // Nama tabel di database
}