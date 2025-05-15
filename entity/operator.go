package entity

import "time"

type Operator struct {
    IdOperator  uint64    `gorm:"primary_key;autoIncrement" json:"id_operator"`
    NamaSekolah string    `gorm:"type:varchar(255);not null" json:"nama_sekolah"`
    Durasi      int       `gorm:"type:int(11);not null" json:"durasi"`
    Status      string    `gorm:"type:enum('Aktif', 'Tidak Aktif');not null" json:"status"`
    IdUser      uint64    `gorm:"type:bigint(20) unsigned" json:"id_user"`
    CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    Users       Users     `gorm:"foreignkey:id_user;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}

func (Operator) TableName() string {
    return "operator" // Nama tabel di database
}
