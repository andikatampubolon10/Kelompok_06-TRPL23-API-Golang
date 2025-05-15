package entity

import "time"

type Bisnis struct {
    IdBisnis        uint64    `gorm:"primary_key;autoIncrement" json:"id_bisnis"`
    NamaSekolah     string    `gorm:"type:varchar(255);not null" json:"nama_sekolah"`
    JumlahPendapatan int      `gorm:"type:int(11);not null" json:"jumlah_pendapatan"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Bisnis) TableName() string {
    return "bisnis" // Nama tabel di database
}
