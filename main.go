package main

import (
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/controller"
	"cbt-api/service"
	"cbt-api/repository"
)

func main() {
	db := config.ConnectDatabase() 

	// Inisialisasi Gin router
	r := gin.Default()

	// Initialize repositories
	jawabanSiswaRepo := repository.NewJawabanSiswaRepository(db)
	jawabanLatihanRepo := repository.NewJawabanLatihanRepository(db)
	soalLatihanRepo := repository.NewSoalLatihanRepository(db)

	// Initialize services
	jawabanSiswaService := service.NewJawabanSiswaService(jawabanSiswaRepo)
	jawabanLatihanService := service.NewJawabanLatihanService(jawabanLatihanRepo)
	soalLatihanService := service.NewSoalLatihanService(soalLatihanRepo)
	jwtService := service.NewJWTService()

	// Initialize controllers
	jawabanSiswaController := controller.NewJawabanSiswaController(jawabanSiswaService, jwtService)
	jawabanLatihanController := controller.NewJawabanLatihanController(
		jawabanLatihanService,
		soalLatihanService,
		jwtService,
	)

	// Setup routes
	setupRoutes(r, jawabanSiswaController, jawabanLatihanController)

	// Rute lainnya
	r.POST("/login", controller.Login)
	r.GET("/kursus-siswa/:id_siswa", controller.GetKursusBySiswa)
	r.GET("/ujian-materi-kursus/:id_kursus", controller.GetUjianAndMateriByKursus)
	r.GET("/profil/:id_siswa", controller.GetSiswaWithKelas)
	r.GET("/kelas", controller.GetAllKelas)
	r.GET("/kurikulum", controller.GetAllKurikulum)
	r.POST("/login-ujian/:id_ujian", controller.AccessUjian)
	r.GET("/soal/:id_soal", controller.GetSoalWithJawaban)
	r.GET("/soal-ujian/:id_ujian", controller.GetSoalByUjian)
	r.GET("/soal-latihan/:id_latihan", controller.GetSoalByLatihan)
	r.GET("/mata-pelajaran/:id_kurikulum", controller.GetMataPelajaranByKurikulum)
	r.GET("/kursus/available/:id_siswa", controller.GetAvailableKursusForSiswa)
	r.GET("/kursus/:id_kursus", controller.GetKursusById)
	r.POST("/kursus/access/:id_kursus", controller.EnrollKursus)
	r.POST("/kursus_siswa/enroll", controller.EnrollKursusSiswa)
	r.GET("/api/latihan-soal/:id_kurikulum/:id_kelas/:id_mata_pelajaran", controller.GetLatihanSoal)
	r.GET("/jawaban-siswa/:id_ujian/:id_siswa", controller.GetJawabanSiswa)
	r.GET("/api/total-nilai-by-ujian/:id_ujian/:id_siswa", controller.GetTotalNilaiSiswaByUjian)
	r.POST("/api/tipe-nilai", controller.CreateTipeNilai)
	r.GET("/api/ujian/:idUjian", controller.GetUjianById)
	r.GET("/nilai-siswa/:id_ujian/:id_siswa", controller.CalculateScore)
	r.GET("api/calculate-and-save-score/:id_ujian/:id_siswa/:id_tipe_ujian", controller.CalculateScore)
	r.GET("/jawaban-soal/:id_jawaban_soal", controller.GetJawabanSoalByID)
	r.POST("/keluar-ujian/:id_ujian", controller.ExitUjian)
	r.GET("/kursus/detail/:id_siswa/:id_kursus", controller.GetKursusWithUjianAndNilai)
	r.POST("/check-attempt-ujian", controller.CheckQuizAttempt)
	r.GET("/check-attempt-ujian/:id_ujian/:id_siswa", controller.CheckQuizAttempt)
	r.GET("/sum_nilai_ujian_kursus/:id_kursus/:id_siswa", controller.GetTotalNilaiByTipeUjian)
	r.POST("/nilai_kursus/:id_kursus/:id_siswa", controller.PostNilaiKursus)
	r.POST("/nilai/:id_kursus/:id_siswa", controller.PostNilai)
	r.GET("/nilai-kursus-siswa/:id_kursus/:id_siswa", controller.GetNilaiByKursusAndSiswa)

	// Jalankan server
	r.Run("192.168.56.1:8080")
}

// Daftarkan semua endpoint jawaban siswa dan jawaban latihan ke router utama
func setupRoutes(
	router *gin.Engine, 
	jawabanSiswaController controller.JawabanSiswaController,
	jawabanLatihanController controller.JawabanLatihanController,
) {
	// Jawaban Siswa routes (for exams)
	jawabanSiswaRoutes := router.Group("api/jawaban-siswa")
	{
		jawabanSiswaRoutes.POST("/", jawabanSiswaController.CreateJawabanSiswa)
		jawabanSiswaRoutes.POST("/batch", jawabanSiswaController.CreateBatchJawabanSiswa)
		jawabanSiswaRoutes.GET("/:id", jawabanSiswaController.GetJawabanSiswaByID)
		jawabanSiswaRoutes.GET("/siswa/:id_siswa", jawabanSiswaController.GetJawabanSiswaBySiswaID)
		jawabanSiswaRoutes.GET("/ujian/:id_ujian", jawabanSiswaController.GetJawabanSiswaByUjianID)
	}
	
	// Jawaban Latihan routes (for practice questions)
	jawabanLatihanRoutes := router.Group("api/jawaban-latihan")
	{
		jawabanLatihanRoutes.POST("/", jawabanLatihanController.CreateJawabanLatihan)
		jawabanLatihanRoutes.POST("/batch", jawabanLatihanController.CreateBatchJawabanLatihan)
		jawabanLatihanRoutes.GET("/:id", jawabanLatihanController.GetJawabanLatihanByID)
		jawabanLatihanRoutes.GET("/siswa/:id_siswa", jawabanLatihanController.GetJawabanLatihanBySiswaID)
		jawabanLatihanRoutes.GET("/latihan/:id_latihan", jawabanLatihanController.GetJawabanLatihanByLatihanID)
	}
	
	// Soal Latihan routes
	soalLatihanRoutes := router.Group("api/soal-latihan")
	{
		soalLatihanRoutes.GET("/:id_latihan", jawabanLatihanController.GetSoalLatihanByID)
		soalLatihanRoutes.GET("/with-jawaban/:id_latihan", jawabanLatihanController.GetSoalWithJawabanByLatihanID)
	}
}