package entity

import "time"

type TipeUjian struct {
    IdTipeUjian   uint64    `gorm:"primary_key;autoIncrement" json:"id_tipe_ujian"`
    NamaTipeUjian string    `gorm:"type:enum('Kuis', 'Ujian');not null" json:"nama_tipe_ujian"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (TipeUjian) TableName() string {
    return "tipe_ujian" // Nama tabel di database
}
