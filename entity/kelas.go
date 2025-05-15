package entity

import "time"

type Kelas struct {
    IdKelas    uint64    `gorm:"primary_key;autoIncrement" json:"id_kelas"`
    NamaKelas  string    `gorm:"type:varchar(255);not null" json:"nama_kelas"`
    // IdOperator uint64    `gorm:"type:bigint(20) unsigned" json:"id_operator"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Operator  Operator  `gorm:"foreignkey:IdOperator;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"operator"`
}

func (Kelas) TableName() string {
    return "kelas" // Nama tabel di database
}
