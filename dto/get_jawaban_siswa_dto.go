package dto

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// JawabanResponse represents a jawaban response
type JawabanResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

// JawabanDetailDTO represents detailed jawaban data
type JawabanDetailDTO struct {
	IDJawabanSiswa uint64  `json:"id_jawaban_siswa"`
	JawabanSiswa   string  `json:"jawaban_siswa"`
	Benar          bool    `json:"benar"`
	NilaiPerSoal   float64 `json:"nilai_per_soal"`
	IDSoal         uint64  `json:"id_soal"`
	IDSiswa        uint64  `json:"id_siswa"`
	IDUjian        uint64  `json:"id_ujian"`
	Grade          float64 `json:"grade"`
	IDJawabanSoal  uint64  `json:"id_jawaban_soal"`
}