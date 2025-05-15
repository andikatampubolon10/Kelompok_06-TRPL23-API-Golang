package entity

import "time"

type JawabanSoal struct {
    IdJawabanSoal uint64    `gorm:"primary_key;autoIncrement" json:"id_jawaban_soal"`
    Jawaban       string    `gorm:"type:varchar(255);not null" json:"jawaban"`
    Benar         bool      `gorm:"type:tinyint(1);not null" json:"benar"`
    IdSoal        uint64    `gorm:"type:bigint(20) unsigned" json:"id_soal"`
    IdTipeSoal    uint64    `gorm:"type:bigint(20) unsigned" json:"id_tipe_soal"`
    CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Soal      Soal      `gorm:"foreignkey:IdSoal;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"soal"`
    TipeSoal  TipeSoal  `gorm:"foreignkey:IdTipeSoal;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_soal"`
}

func (JawabanSoal) TableName() string {
    return "jawaban_soal" // Nama tabel di database
}
