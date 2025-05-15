package entity

import "time"

type Siswa struct {
    IdSiswa     uint64    `gorm:"primary_key;autoIncrement" json:"id_siswa"`
    NamaSiswa   string    `gorm:"type:varchar(255);not null" json:"nama_siswa"`
    NIS         int       `gorm:"type:int(11);not null" json:"nis"`
    Status      string    `gorm:"type:enum('Aktif', 'Tidak Aktif');not null" json:"status"`
    IdUser      uint64    `gorm:"type:bigint(20) unsigned" json:"id_user"`
    // IdOperator  uint64    `gorm:"type:bigint(20) unsigned" json:"id_operator"`
    IdKelas     uint64    `gorm:"type:bigint(20) unsigned" json:"id_kelas"`
    CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`


    User     Users   `gorm:"foreignkey:IdUser;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
    // Operator Operator `gorm:"foreignkey:IdOperator;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"operator"`
    Kelas    Kelas    `gorm:"foreignkey:IdKelas;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kelas"`
}

func (Siswa) TableName() string {
    return "siswa" // Nama tabel di database
}
