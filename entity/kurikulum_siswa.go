package entity

import "time"

type KurikulumSiswa struct {
    IdKurikulumSiswa uint64    `gorm:"primary_key;autoIncrement" json:"id_kurikulum_siswa"`
    IdKurikulum      uint64    `gorm:"type:bigint(20) unsigned" json:"id_kurikulum"`
    IdSiswa          uint64    `gorm:"type:bigint(20) unsigned" json:"id_siswa"`
   CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Siswa      Siswa     `gorm:"foreignkey:IdSiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"siswa"`
    Kurikulum  Kurikulum `gorm:"foreignkey:IdKurikulum;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kurikulum"`
}
