package controller

import (
	"cbt-api/entity"
	"github.com/gin-gonic/gin"
	"cbt-api/helper"   // Add this import
	"cbt-api/service"
	"net/http"
	"strconv"
	"time"
)

// JawabanSiswaController is a contract for jawaban siswa controller
type JawabanSiswaController interface {
	CreateJawabanSiswa(ctx *gin.Context)
	CreateBatchJawabanSiswa(ctx *gin.Context)
	GetJawabanSiswaByID(ctx *gin.Context)
	GetJawabanSiswaBySiswaID(ctx *gin.Context)
	GetJawabanSiswaByUjianID(ctx *gin.Context)
}

type jawabanSiswaController struct {
	jawabanSiswaService service.JawabanSiswaService
	jwtService          service.JWTService
}

// NewJawabanSiswaController creates a new instance of JawabanSiswaController
func NewJawabanSiswaController(jawabanSiswaService service.JawabanSiswaService, jwtService service.JWTService) JawabanSiswaController {
	return &jawabanSiswaController{
		jawabanSiswaService: jawabanSiswaService,
		jwtService:          jwtService,
	}
}

// CreateJawabanSiswa creates a new jawaban siswa
func (c *jawabanSiswaController) CreateJawabanSiswa(ctx *gin.Context) {
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

	result, err := c.jawabanSiswaService.CreateJawabanSiswa(jawabanSiswa)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to create jawaban siswa", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban siswa created", result)
	ctx.JSON(http.StatusCreated, response)
}

// CreateBatchJawabanSiswa creates multiple jawaban siswa at once
func (c *jawabanSiswaController) CreateBatchJawabanSiswa(ctx *gin.Context) {
	var request struct {
		JawabanList []entity.JawabanSiswa `json:"jawaban_list"`
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Set timestamps for all jawaban
	now := time.Now()
	for i := range request.JawabanList {
		request.JawabanList[i].CreatedAt = now
		request.JawabanList[i].UpdatedAt = now
	}

	results, err := c.jawabanSiswaService.CreateBatchJawabanSiswa(request.JawabanList)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to create batch jawaban siswa", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Batch jawaban siswa created", results)
	ctx.JSON(http.StatusCreated, response)
}

// GetJawabanSiswaByID gets a jawaban siswa by ID
func (c *jawabanSiswaController) GetJawabanSiswaByID(ctx *gin.Context) {
	id := ctx.Param("id_siswa")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.jawabanSiswaService.GetJawabanSiswaByID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban siswa", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban siswa retrieved", result)
	ctx.JSON(http.StatusOK, response)
}

// GetJawabanSiswaBySiswaID gets all jawaban siswa by siswa ID
func (c *jawabanSiswaController) GetJawabanSiswaBySiswaID(ctx *gin.Context) {
	id := ctx.Param("id_siswa")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	results, err := c.jawabanSiswaService.GetJawabanSiswaBySiswaID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban siswa", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban siswa retrieved", results)
	ctx.JSON(http.StatusOK, response)
}

// GetJawabanSiswaByUjianID gets all jawaban siswa by ujian ID
func (c *jawabanSiswaController) GetJawabanSiswaByUjianID(ctx *gin.Context) {
	id := ctx.Param("id_ujian")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid ID", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	results, err := c.jawabanSiswaService.GetJawabanSiswaByUjianID(idUint)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get jawaban siswa", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "Jawaban siswa retrieved", results)
	ctx.JSON(http.StatusOK, response)
}