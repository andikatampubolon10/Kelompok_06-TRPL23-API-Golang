package entity

import "time"

type MataPelajaranSiswa struct {
    IdMataPelajaranSiswa uint64    `gorm:"primary_key;autoIncrement" json:"id_mata_pelajaran_siswa"`
    IdSiswa              uint64    `gorm:"type:bigint(20) unsigned" json:"id_siswa"`
    IdMataPelajaran      uint64    `gorm:"type:bigint(20) unsigned" json:"id_mata_pelajaran"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Siswa           Siswa            `gorm:"foreignkey:IdSiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"siswa"`
    MataPelajaran  MataPelajaran    `gorm:"foreignkey:IdMataPelajaran;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"mata_pelajaran"`
}
