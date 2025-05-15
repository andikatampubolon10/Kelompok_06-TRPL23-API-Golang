package entity

import "time"

type MataPelajaran struct {
    IdMataPelajaran   uint64    `gorm:"primary_key;autoIncrement" json:"id_mata_pelajaran"`
    NamaMataPelajaran string  `gorm:"type:varchar(255);not null" json:"nama_mata_pelajaran"`
    // IdOperator       uint64   `gorm:"type:bigint(20) unsigned" json:"id_operator"`
    IdKurikulum       uint64   `gorm:"type:bigint(20) unsigned" json:"id_kurikulum"`
    CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Operator   Operator  `gorm:"foreignkey:IdOperator;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"operator"`
    Kurikulum  Kurikulum `gorm:"foreignkey:IdKurikulum;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"kurikulum"`
}

func (MataPelajaran) TableName() string {
    return "mata_pelajaran" // Nama tabel di database
}
