package controller

import (
	"net/http"
	"prestasi/app/models"

	"github.com/gin-gonic/gin"
)

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
