// controller/jawaban_latihan_controller.go
package controller

import (
	"cbt-api/entity"
	"cbt-api/helper"
	"cbt-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// JawabanLatihanController is a contract for jawaban latihan controller
type JawabanLatihanController interface {
	CreateJawabanLatihan(ctx *gin.Context)
	CreateBatchJawabanLatihan(ctx *gin.Context)
	GetJawabanLatihanByID(ctx *gin.Context)
	GetJawabanLatihanBySiswaID(ctx *gin.Context)
	GetJawabanLatihanByLatihanID(ctx *gin.Context)
	GetSoalLatihanByID(ctx *gin.Context)
	GetSoalWithJawabanByLatihanID(ctx *gin.Context)
}

type jawabanLatihanController struct {
	jawabanLatihanService service.JawabanLatihanService
	soalLatihanService    service.SoalLatihanService
	jwtService            service.JWTService
}

// NewJawabanLatihanController creates a new instance of JawabanLatihanController
func NewJawabanLatihanController(
	jawabanLatihanService service.JawabanLatihanService,
	soalLatihanService service.SoalLatihanService,
	jwtService service.JWTService,
) JawabanLatihanController {
	return &jawabanLatihanController{
		jawabanLatihanService: jawabanLatihanService,
		soalLatihanService:    soalLatihanService,
		jwtService:            jwtService,
	}
}

// CreateJawabanLatihan creates a new jawaban for latihan
func (c *jawabanLatihanController) CreateJawabanLatihan(ctx *gin.Context) {
	var jawabanSiswa entity.JawabanSiswa
	err := ctx.ShouldBindJSON(&jawabanSiswa)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Set timestamps
	jawabanSiswa.CreatedAt = time.Now()
	jawabanSiswa.UpdatedAt = time.Now()

	result, err := c.jawabanLatihanService.CreateJawabanLatihan(jawabanSiswa)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to create jawaban latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban latihan created", result)
	ctx.JSON(http.StatusCreated, response)
}

// CreateBatchJawabanLatihan creates multiple jawaban for latihan at once
func (c *jawabanLatihanController) CreateBatchJawabanLatihan(ctx *gin.Context) {
	var request struct {
		JawabanList []entity.JawabanSiswa `json:"jawaban_list"`
		IdLatihan   string                `json:"id_latihan"` // Keep it as string in the JSON body
	}

	// Attempt to bind the incoming JSON payload
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// Log the error to see what went wrong during binding
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Parse id_latihan to uint64
	latihanId, err := strconv.ParseUint(request.IdLatihan, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid id_latihan format", "id_latihan should be a valid number", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Debugging: print the received IdLatihan after conversion

	// Set timestamps for all jawaban
	now := time.Now()
	for i := range request.JawabanList {
		request.JawabanList[i].CreatedAt = now
		request.JawabanList[i].UpdatedAt = now
	}

	// Call the service to handle the business logic
	results, err := c.jawabanLatihanService.CreateBatchJawabanLatihan(request.JawabanList, latihanId)
	if err != nil {
		// Log the error to understand what went wrong
		response := helper.BuildErrorResponse("Failed to create batch jawaban latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Return the success response
	response := helper.BuildResponse(true, "Batch jawaban latihan created", results)
	ctx.JSON(http.StatusCreated, response)
}


// GetJawabanLatihanByID gets a jawaban latihan by ID
func (c *jawabanLatihanController) GetJawabanLatihanByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.jawabanLatihanService.GetJawabanLatihanByID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban latihan retrieved", result)
	ctx.JSON(http.StatusOK, response)
}

// GetJawabanLatihanBySiswaID gets all jawaban latihan by siswa ID
func (c *jawabanLatihanController) GetJawabanLatihanBySiswaID(ctx *gin.Context) {
	id := ctx.Param("id_siswa")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	results, err := c.jawabanLatihanService.GetJawabanLatihanBySiswaID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban latihan retrieved", results)
	ctx.JSON(http.StatusOK, response)
}

// GetJawabanLatihanByLatihanID gets all jawaban latihan by latihan ID
func (c *jawabanLatihanController) GetJawabanLatihanByLatihanID(ctx *gin.Context) {
	id := ctx.Param("id_latihan")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	results, err := c.jawabanLatihanService.GetJawabanLatihanByLatihanID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban latihan retrieved", results)
	ctx.JSON(http.StatusOK, response)
}

// GetSoalLatihanByID gets all soal for a specific latihan ID
func (c *jawabanLatihanController) GetSoalLatihanByID(ctx *gin.Context) {
	id := ctx.Param("id_latihan")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	results, err := c.soalLatihanService.GetSoalLatihanByID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get soal latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Soal latihan retrieved", results)
	ctx.JSON(http.StatusOK, response)
}

// GetSoalWithJawabanByLatihanID gets all soal with their jawaban options for a specific latihan ID
func (c *jawabanLatihanController) GetSoalWithJawabanByLatihanID(ctx *gin.Context) {
	id := ctx.Param("id_latihan")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	// Get all soal for this latihan
	soalList, err := c.soalLatihanService.GetSoalLatihanByID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get soal latihan", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	// Get all jawaban options for these soal
	jawabanMap, err := c.soalLatihanService.GetSoalWithJawabanByLatihanID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban options", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	// Create a response structure that combines soal with their jawaban options
	type SoalWithJawaban struct {
		Soal    entity.Soal          `json:"soal"`
		Jawaban []entity.JawabanSoal `json:"jawaban"`
	}
	
	var result []SoalWithJawaban
	for _, soal := range soalList {
		jawaban, exists := jawabanMap[soal.IdSoal]
		if !exists {
			jawaban = []entity.JawabanSoal{} // Empty array if no jawaban found
		}
		
		result = append(result, SoalWithJawaban{
			Soal:    soal,
			Jawaban: jawaban,
		})
	}

	response := helper.BuildResponse(true, "Soal with jawaban retrieved", result)
	ctx.JSON(http.StatusOK, response)
}