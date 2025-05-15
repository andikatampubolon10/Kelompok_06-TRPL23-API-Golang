package entity

import "time"

type Kursus struct {
    IdKursus   uint64    `gorm:"primary_key;autoIncrement" json:"id_kursus"`
    NamaKursus string   `gorm:"type:varchar(255);not null" json:"nama_kursus"`
    Password    string    `gorm:"type:varchar(255);not null" json:"-"`
    Image       string    `gorm:"type:varchar(255);not null" json:"image"`
    ImageUrl       string    `gorm:"type:varchar(255);not null" json:"image_url"`
    // IdGuru     uint64    `gorm:"type:bigint(20) unsigned" json:"id_guru"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Guru   Guru   `gorm:"foreignkey:IdGuru;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"guru"`
}

func (Kursus) TableName() string {
    return "kursus" // Nama tabel di database
}
