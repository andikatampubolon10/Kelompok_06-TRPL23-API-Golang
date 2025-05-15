package entity

import "time"

type Guru struct {
    IdGuru     uint64    `gorm:"primary_key;autoIncrement" json:"id_guru"`
    NamaGuru   string    `gorm:"type:varchar(255);not null" json:"nama_guru"`
    NIP        int       `gorm:"type:int(11);not null" json:"nip"`
    Status     string    `gorm:"type:enum('Aktif', 'Tidak Aktif');not null" json:"status"`
    IdUser     uint64   `gorm:"type:bigint(20) unsigned" json:"id_user"` // Foreign Key ke Users
    IdOperator uint64    `gorm:"type:bigint(20) unsigned" json:"id_operator"` // Foreign Key ke Operators
    CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Foreign key relations
    Users      Users     `gorm:"foreignkey:id_user;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
    Operator   Operator  `gorm:"foreignkey:id_operator;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"operator"`
}

func (Guru) TableName() string {
    return "guru" // Nama tabel di database
}
