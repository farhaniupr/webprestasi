package controller

import (
	"fmt"
	"net/http"
	"prestasi/app/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Nama                string     `json:"nama"`
	Idprestasi          int        `gorm:"primary_key";auto_increment;not_null json:"id_prestasi"`
	Idmhs               int        `json:"id_mhs"`
	Namakegiatan        string     `json:"nama_kegiatan"`
	Namapenyelenggaraan string     `json:"nama_penyelenggaraan"`
	URL                 string     `json:"url"`
	Jumlah              int        `json:"jumlah"`
	Kategorikegiatan    string     `json:"kategori_kegiatan"`
	Tingkatkegiatan     string     `json:"tingkat_kegiatan"`
	Hasilkegiatan       string     `json:"hasil_kegiatan"`
	Tempatkegiatan      string     `json:"tempat_kegiatan"`
	Tanggalawal         *time.Time `json:"tanggal_awal"`
	Tanggalakhir        *time.Time `json:"tanggal_akhir"`
	Unggahsertifikat    string     `json:"unggah_sertifikat"`
	Unggahsurattugas    string     `json:"unggah_surat_tugas"`
	Unggahfoto          string     `json:"unggah_foto"`
	Status              string     `json:"status"`
}

type Prestasi struct {
	Idprestasi          int        `gorm:"primary_key";auto_increment;not_null json:"id_prestasi"`
	Idmhs               int        `json:"id_mhs"`
	Namakegiatan        string     `json:"nama_kegiatan"`
	Namapenyelenggaraan string     `json:"nama_penyelenggaraan"`
	URL                 string     `json:"url"`
	Jumlah              int        `json:"jumlah"`
	Kategorikegiatan    string     `json:"kategori_kegiatan"`
	Tingkatkegiatan     string     `json:"tingkat_kegiatan"`
	Hasilkegiatan       string     `json:"hasil_kegiatan"`
	Tempatkegiatan      string     `json:"tempat_kegiatan"`
	Tanggalawal         *time.Time `json:"tanggal_awal"`
	Tanggalakhir        *time.Time `json:"tanggal_akhir"`
	Unggahsertifikat    string     `json:"unggah_sertifikat"`
	Unggahsurattugas    string     `json:"unggah_surat_tugas"`
	Unggahfoto          string     `json:"unggah_foto"`
	Status              string     `json:"status"`
}

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
	//var prestasi []models.Prestasi
	var result []Result

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, prestasis.jumlah, mahasiswas.nama, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
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
	//var prestasi models.Prestasi
	var result Result
	parprestasi := c.Param("id_prestasi")

	//models.DB.Where("id_prestasi = ?", parprestasi).First(&prestasi)

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, mahasiswas.nama, prestasis.jumlah, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Where("idprestasi = ?", parprestasi).Scan(&result)

	c.JSON(http.StatusOK, gin.H{"data": result})
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
	//var prestasiT Prestasi

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
