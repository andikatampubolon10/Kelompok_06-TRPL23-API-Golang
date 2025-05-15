package entity

import "time"

type Kurikulum struct {
    IdKurikulum  uint64    `gorm:"primary_key;autoIncrement" json:"id_kurikulum"`
    NamaKurikulum string   `gorm:"type:varchar(255);not null" json:"nama_kurikulum"`
    // IdOperator   uint64    `gorm:"type:bigint(20) unsigned" json:"id_operator"`
   CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Operator  Operator  `gorm:"foreignkey:IdOperator;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"operator"`
}

func (Kurikulum) TableName() string {
    return "kurikulum" // Nama tabel di database
}
