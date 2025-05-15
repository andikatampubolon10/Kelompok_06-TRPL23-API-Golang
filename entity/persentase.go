package entity

import (
	"time"
)

type Persentase struct {
	IdPersentase       uint64    `gorm:"primary_key;autoIncrement" json:"id_persentase"`
	Persentase        float64   `gorm:"type:decimal(5,0);not null;default:0" json:"persentase"`
	// IdTipepersentase   uint64    `gorm:"type:bigint(20) unsigned" json:"id_tipe_persentase"`
	IdKursus           uint64    `gorm:"type:bigint(20) unsigned" json:"id_kursus"`
	IdTipeUjian        uint64    `gorm:"type:bigint(20) unsigned" json:"id_tipe_ujian"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

	// TipePersentase     TipePersentase `gorm:"foreignkey:IdTipepersentase;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_persentase"`
	Kursus             Kursus         `gorm:"foreignkey:IdKursus;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kursus"`
	TipeUjian          TipeUjian      `gorm:"foreignkey:IdTipeUjian;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_ujian"`
}

func (Persentase) TableName() string {
    return "persentase" // Nama tabel di database
}
