package entity

import "time"

type JawabanSiswa struct {
    IdJawabanSiswa uint64    `gorm:"primary_key;autoIncrement" json:"id_jawaban_siswa"`
    JawabanText    string    `gorm:"type:varchar(255);not null" json:"jawaban_siswa"`
    IdSoal         uint64    `gorm:"type:bigint(20) unsigned" json:"id_soal"`
    IdSiswa        uint64    `gorm:"type:bigint(20) unsigned" json:"id_siswa"`
    IdJawabanSoal  uint64    `gorm:"type:bigint(20) unsigned" json:"id_jawaban_soal"`
    CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Soal      Soal      `gorm:"foreignkey:IdSoal;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"soal"`
    Siswa     Siswa     `gorm:"foreignkey:IdSiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"siswa"`
    JawabanSoal JawabanSoal `gorm:"foreignkey:IdJawabanSoal;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"jawaban_soal"`
}

func (JawabanSiswa) TableName() string {
    return "jawaban_siswa" // Nama tabel di database
}
