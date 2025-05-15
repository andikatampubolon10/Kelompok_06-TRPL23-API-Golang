package entity

import "time"

type KursusSiswa struct {
    IdKursusSiswa  uint64    `gorm:"primary_key;autoIncrement" json:"id_kursus_siswa"`
    IdSiswa        uint64    `gorm:"type:bigint(20) unsigned" json:"id_siswa"`
    IdKursus       uint64    `gorm:"type:bigint(20) unsigned" json:"id_kursus"`
    CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Kursus Kursus `gorm:"foreignKey:IdKursus;references:IdKursus;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"kursus"`
    Siswa Siswa `gorm:"foreignKey:IdSiswa;references:IdSiswa;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"siswa"`
}

func (KursusSiswa) TableName() string {
    return "kursus_siswa" // Nama tabel di database
}
