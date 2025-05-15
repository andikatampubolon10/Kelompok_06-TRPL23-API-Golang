package entity

import "time"

type Materi struct {
    IdMateri          uint64    `gorm:"primary_key;autoIncrement" json:"id_materi"`
    JudulMateri string  `gorm:"type:varchar(255);not null" json:"judul_materi"`
    Deskripsi       string  `gorm:"type:varchar(255);not null" json:"deskripsi"`
    // FilePath      string  `gorm:"type:varchar(255);not null" json:"file_path"`
	IdKursus      uint64   `gorm:"type:bigint(20) unsigned" json:"id_kursus"`
	// IdGuru      uint64   `gorm:"type:bigint(20) unsigned" json:"id_guru"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Guru   Guru  `gorm:"foreignkey:IdGuru;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"guru"`
    Kursus  Kursus `gorm:"foreignkey:IdKursus;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kursus"`
}

func (Materi) TableName() string {
    return "materi" // Nama tabel di database
}
