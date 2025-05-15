package entity

import "time"

type TipeNilai struct {
    IdTipeNilai   uint64    `gorm:"primary_key;autoIncrement" json:"id_tipe_nilai"`
    Nilai         float64   `gorm:"type:decimal(5,0);not null" json:"nilai"`
    IdTipeUjian   uint64    `gorm:"type:int(11);not null" json:"id_tipe_ujian"`
    IdSiswa       uint64    `gorm:"type:int(11);not null" json:"id_siswa"`
    IdUjian       uint64    `gorm:"type:int(11);not null" json:"id_ujian"`
    CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    TipeUjian     TipeUjian `gorm:"foreignkey:IdTipeUjian;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"tipe_ujian"`
    Siswa         Siswa     `gorm:"foreignkey:IdSiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"siswa"`
    Ujian         Ujian     `gorm:"foreignkey:IdUjian;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"ujian"`
}

func (TipeNilai) TableName() string {
     return "tipe_nilai" // Nama tabel di database
 }