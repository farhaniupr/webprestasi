package controller

import (
	"fmt"
	"net/http"
	"prestasi/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Get
func GetMahasiswa(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	models.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

func GetProdi(c *gin.Context) {
	var prodi []models.Fakultasprodi
	models.DB.Find(&prodi)
	c.JSON(http.StatusOK, gin.H{"data": prodi})
}

func GetOrganisasi(c *gin.Context) {
	var organisasi []models.Organisasi
	models.DB.Find(&organisasi)
	c.JSON(http.StatusOK, gin.H{"data": organisasi})
}

func GetPengabdian(c *gin.Context) {
	var pengabdian []models.Pengabdian
	models.DB.Find(&pengabdian)
	c.JSON(http.StatusOK, gin.H{"data": pengabdian})
}

func GetPrestasiNon(c *gin.Context) {
	var prestasinon []models.Prestasinonkompetisi
	models.DB.Find(&prestasinon)
	c.JSON(http.StatusOK, gin.H{"data": prestasinon})
}

func GetPrestasi(c *gin.Context) {
	var prestasi []models.Prestasi
	models.DB.Find(&prestasi)
	c.JSON(http.StatusOK, gin.H{"data": prestasi})
}

//Get 1 Row
func GetOneMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	paridmhs := c.Param("idmhs")

	models.DB.Where("idmhs = ?", paridmhs).First(&mahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

func GetOneProdi(c *gin.Context) {
	var prodi models.Fakultasprodi
	parkopi := c.Param("kode_prodi")

	models.DB.Where("kodeprodi = ?", parkopi).First(&prodi)

	c.JSON(http.StatusOK, gin.H{"data": prodi})
}

func GetOnePrestasi(c *gin.Context) {
	var prestasi models.Prestasi
	parprestasi := c.Param("id_prestasi")

	models.DB.Where("id_prestasi = ?", parprestasi).First(&prestasi)

	c.JSON(http.StatusOK, gin.H{"data": prestasi})
}

//Insert
func AddMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := c.Bind(&mahasiswa); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	result := models.DB.Create(&mahasiswa)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func AddProdi(c *gin.Context) {
	var prodi models.Fakultasprodi

	if err := c.Bind(&prodi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	result := models.DB.Create(&prodi)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func AddPrestasi(c *gin.Context) {
	var prestasi models.Prestasi

	if err := c.Bind(&prestasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	result := models.DB.Create(&prestasi)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

//Edit
func EditMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := c.Bind(&mahasiswa); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var paridmhs string
	paridmhs = c.Param("idmhs")

	paridmhsConvert, err := strconv.Atoi(paridmhs)

	if err != nil {
		fmt.Println("eror")
	}

	mahasiswa.Idmhs = paridmhsConvert

	result := models.DB.Where("idmhs= ?", mahasiswa.Idmhs).Updates(&mahasiswa)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditProdi(c *gin.Context) {
	var prodi models.Fakultasprodi

	if err := c.Bind(&prodi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parkopi string
	parkopi = c.Param("kode_prodi")

	prodi.Kodeprodi = parkopi

	result := models.DB.Where("kodeprodi = ?", prodi.Kodeprodi).Updates(&prodi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditPrestasi(c *gin.Context) {
	var prestasi models.Prestasi

	if err := c.Bind(&prestasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parprestasi string
	parprestasi = c.Param("id_prestasi")

	parprestasiConvert, err := strconv.Atoi(parprestasi)

	if err != nil {
		fmt.Println("eror")
	}
	prestasi.Idprestasi = parprestasiConvert

	result := models.DB.Where("Idprestasi = ?", prestasi.Idprestasi).Updates(&prestasi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

//Delete
func DeleteMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := c.Bind(&mahasiswa); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	paridmhs := c.Param("idmhs")

	result := models.DB.Where("idmhs=?", paridmhs).Delete(&mahasiswa)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func DeleteProdi(c *gin.Context) {
	var prodi models.Fakultasprodi

	if err := c.Bind(&prodi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	parkopi := c.Param("kode_prodi")

	result := models.DB.Where("kodeprodi=?", parkopi).Delete(&prodi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func DeletePrestasi(c *gin.Context) {
	var prestasi models.Prestasi

	if err := c.Bind(&prestasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	parprestasi := c.Param("id_prestasi")

	result := models.DB.Where("idprestasi=?", parprestasi).Delete(&prestasi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}
