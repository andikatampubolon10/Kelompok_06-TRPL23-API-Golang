package entity

import "time"

type TipeSoal struct {
    IdTipeSoal   uint64    `gorm:"primary_key;autoIncrement" json:"id_tipe_soal"`
    NamaTipeUjian string   `gorm:"type:enum('Pilihan_Berganda', 'Benar_Salah', 'Isian');not null" json:"nama_tipe_ujian"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (TipeSoal) TableName() string {
    return "tipe_soal" // Nama tabel di database
}
