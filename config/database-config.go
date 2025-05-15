package config

import (
	// 
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// dsn := "root:@tcp(127.0.0.1:3306)/pkm10?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:KEoDCfuvLDybxucWYnOWHSFHicWdJetR@tcp(metro.proxy.rlwy.net:50393)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	// mysql://root:KEoDCfuvLDybxucWYnOWHSFHicWdJetR@metro.proxy.rlwy.net:50393/railway
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal konek ke database:", err)
	}

	DB = db

	err = db.AutoMigrate(
        // &entity.Materi{},
		// &entity.JawabanSiswa{},
        // &entity.JawabanSoal{},
		// &entity.Soal{},
		// &entity.TipeNilai{},
        // &entity.Latihan{},
        // &entity.MataPelajaranSiswa{},
        // &entity.Kelas{},
		// &entity.MataPelajaran{},
		// &entity.KurikulumSiswa{},
		// &entity.Kurikulum{},
		// &entity.KursusSiswa{},
		// &entity.Kursus{},
		// &entity.Siswa{},
		// &entity.Users{},
		// &entity.Ujian{},
		// &entity.TipeUjian{},
		// &entity.TipeSoal{},
		// &entity.Kursus{},
		// &entity.Guru{},
		// &entity.TipeNilai{},
		// &entity.Kelas{},
		// &entity.Operator{},
		// &entity.Bisnis{},
		// &entity.Nilai{},
		// &entity.NilaiKursus{},
		
    )
	// if err != nil {
	// 	fmt.Println("Gagal AutoMigrate:", err)
	// }

	fmt.Println("Koneksi dan migrasi sukses")
    return DB
}
