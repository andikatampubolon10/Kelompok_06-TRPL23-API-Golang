package entity

import "time"

type Ujian struct {
    IdUjian        uint64    `gorm:"primary_key;autoIncrement" json:"id_ujian"`
    NamaUjian      string    `gorm:"type:varchar(255);not null" json:"nama_ujian"`
    Acak           string    `gorm:"type:enum('Aktif', 'Tidak Aktif');not null" json:"acak"`
    StatusJawaban  string    `gorm:"type:enum('Aktif', 'Tidak Aktif');not null" json:"status_jawaban"`
    Grade          float64   `gorm:"type:double" json:"grade"`

    PasswordMasuk  string    `gorm:"type:varchar(255);not null" json:"password_masuk"`
    PasswordKeluar string    `gorm:"type:varchar(255);not null" json:"password_keluar"`
    WaktuMulai     time.Time `gorm:"type:timestamp" json:"waktu_mulai"`
    WaktuSelesai   time.Time `gorm:"type:timestamp" json:"waktu_selesai"`
    Durasi         int       `gorm:"type:int(11)" json:"durasi"`
    TanggalUjian   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"tanggal_ujian"`

    IdKursus       uint64    `gorm:"type:bigint(20) unsigned;not null" json:"id_kursus"`
    // IdGuru         uint64    `gorm:"type:bigint(20) unsigned;not null" json:"id_guru"`
    IdTipeUjian    uint64    `gorm:"type:bigint(20) unsigned;not null" json:"id_tipe_ujian"`

    CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

    // Relations
    Kursus         Kursus     `gorm:"foreignKey:IdKursus;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"kursus"`
    // Guru           Guru       `gorm:"foreignKey:IdGuru;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"guru"`
    TipeUjian      TipeUjian  `gorm:"foreignKey:IdTipeUjian;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"tipe_ujian"`
}

func (Ujian) TableName() string {
    return "ujian" // Nama tabel di database
}
