package controller

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"prestasi/app/models"
	"strconv"
	"strings"
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

type ResultOrganisasi struct {
	Idorganisasi   int    `json:"id_organisasi"`
	Idmhs          int    `json:"id_mhs"`
	Nama           string `json:"nama"`
	Namaorganisasi string `json:"nama_organisasi"`
	Jabatanmhs     string `json:"jabatan_mhs"`
	Periode        string `json:"periode"`
	Unggahsk       string `json:"unggahsk"`
	Status         string `json:"status"`
}

type ResultPengabdian struct {
	Idpengabdian  int        `gorm:"primary_key";auto_increment;not_null json:"id_pengabdian"`
	Idmhs         int        `json:"id_mhs"`
	Nama          string     `json:"nama"`
	Namaprogram   string     `json:"nama_program"`
	Tahunkegiatan *time.Time `json:"tahun_kegiatan"`
	Unggahsurat   string     `json:"unggah_surat"`
	Status        string     `json:"status"`
}

type Pengabdian struct {
	Idpengabdian  int    `gorm:"primary_key";auto_increment;not_null json:"id_pengabdian"`
	Idmhs         int    `json:"id_mhs"`
	Nama          string `json:"nama"`
	Namaprogram   string `json:"nama_program"`
	Tahunkegiatan string `json:"tahun_kegiatan"`
	Unggahsurat   string `json:"unggah_surat"`
	Status        string `json:"status"`
}
type Prestasinonkompetisi struct {
	Idprestasinon       int        `gorm:"primary_key";auto_increment;not_null json:"id_prestasi_non"`
	Idmhs               int        `json:"id_mhs"`
	Nama                string     `json:"nama"`
	Namapenyelenggaraan string     `json:"nama_penyelenggaraan"`
	Namakegiatan        string     `json:"nama_kegiatan"`
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

type ResultCount struct {
	Jumlah int `json:"jumlah"`
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
	var organisasi []ResultOrganisasi
	//models.DB.Find(&organisasi)
	models.DB.Table("organisasis").Select("organisasis.idorganisasi, organisasis.idmhs, organisasis.namaorganisasi, organisasis.jabatanmhs, organisasis.periode, organisasis.unggahsk, organisasis.status, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = organisasis.idmhs").Scan(&organisasi)
	c.JSON(http.StatusOK, gin.H{"data": organisasi})
}

func GetOrganisasiST(c *gin.Context) {
	var organisasi []ResultOrganisasi
	//models.DB.Find(&organisasi)
	models.DB.Table("organisasis").Where("status = ?", "setuju").Select("organisasis.idorganisasi, organisasis.idmhs, organisasis.namaorganisasi, organisasis.jabatanmhs, organisasis.periode, organisasis.unggahsk, organisasis.status, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = organisasis.idmhs").Scan(&organisasi)
	c.JSON(http.StatusOK, gin.H{"data": organisasi})
}

func GetOrganisasiTST(c *gin.Context) {
	var organisasi []ResultOrganisasi
	//models.DB.Find(&organisasi)
	models.DB.Table("organisasis").Where("status = ?", "").Select("organisasis.idorganisasi, organisasis.idmhs, organisasis.namaorganisasi, organisasis.jabatanmhs, organisasis.periode, organisasis.unggahsk, organisasis.status, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = organisasis.idmhs").Scan(&organisasi)
	c.JSON(http.StatusOK, gin.H{"data": organisasi})
}

func GetPengabdian(c *gin.Context) {
	var pengabdian []Pengabdian
	//models.DB.Find(&pengabdian)

	models.DB.Table("pengabdians").Select("pengabdians.idpengabdian, pengabdians.idmhs, pengabdians.namaprogram, pengabdians.tahunkegiatan, pengabdians.unggahsurat, pengabdians.status, mahasiswas.nama, mahasiswas.idmhs").
		Joins("left join mahasiswas on mahasiswas.idmhs = pengabdians.idmhs").Scan(&pengabdian)

	c.JSON(http.StatusOK, gin.H{"data": pengabdian})
}

func GetPengabdianTST(c *gin.Context) {
	var pengabdian []Pengabdian
	//models.DB.Find(&pengabdian)

	models.DB.Table("pengabdians").Where("status = ?", "").Select("pengabdians.idpengabdian, pengabdians.idmhs, pengabdians.namaprogram, pengabdians.tahunkegiatan, pengabdians.unggahsurat, pengabdians.status, mahasiswas.nama, mahasiswas.idmhs").
		Joins("left join mahasiswas on mahasiswas.idmhs = pengabdians.idmhs").Scan(&pengabdian)

	c.JSON(http.StatusOK, gin.H{"data": pengabdian})
}

func GetPengabdianST(c *gin.Context) {
	var pengabdian []Pengabdian
	//models.DB.Find(&pengabdian)

	models.DB.Table("pengabdians").Where("status = ?", "setuju").Select("pengabdians.idpengabdian, pengabdians.idmhs, pengabdians.namaprogram, pengabdians.tahunkegiatan, pengabdians.unggahsurat, pengabdians.status, mahasiswas.nama, mahasiswas.idmhs").
		Joins("left join mahasiswas on mahasiswas.idmhs = pengabdians.idmhs").Scan(&pengabdian)

	c.JSON(http.StatusOK, gin.H{"data": pengabdian})
}

func GetPrestasi(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Result

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, prestasis.jumlah, mahasiswas.nama, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasiConditionTS(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Result

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, prestasis.jumlah, mahasiswas.nama, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Where("status = ?", "").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasiConditionST(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Result

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, prestasis.jumlah, mahasiswas.nama, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Where("status = ?", "setuju").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasiConditionTST(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Result

	models.DB.Table("prestasis").Select("prestasis.idmhs, prestasis.namakegiatan, prestasis.jumlah, mahasiswas.nama, prestasis.idprestasi, prestasis.namapenyelenggaraan, prestasis.url, prestasis.kategorikegiatan, prestasis.tingkatkegiatan, prestasis.hasilkegiatan, prestasis.tempatkegiatan, prestasis.tanggalawal, prestasis.tanggalakhir, prestasis.unggahsertifikat, prestasis.unggahsurattugas, prestasis.unggahfoto, prestasis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasis.idmhs").Where("status = ?", "").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasinon(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Prestasinonkompetisi

	models.DB.Table("prestasinonkompetisis").Select("prestasinonkompetisis.idmhs, prestasinonkompetisis.namakegiatan, mahasiswas.nama, prestasinonkompetisis.idprestasinon, prestasinonkompetisis.namapenyelenggaraan,	 prestasinonkompetisis.tempatkegiatan, prestasinonkompetisis.tanggalawal, prestasinonkompetisis.tanggalakhir, prestasinonkompetisis.unggahsertifikat, prestasinonkompetisis.unggahsurattugas, prestasinonkompetisis.unggahfoto, prestasinonkompetisis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasinonkompetisis.idmhs").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasinonST(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Prestasinonkompetisi

	models.DB.Table("prestasinonkompetisis").Select("prestasinonkompetisis.idmhs, prestasinonkompetisis.namakegiatan, mahasiswas.nama, prestasinonkompetisis.idprestasinon, prestasinonkompetisis.namapenyelenggaraan,	 prestasinonkompetisis.tempatkegiatan, prestasinonkompetisis.tanggalawal, prestasinonkompetisis.tanggalakhir, prestasinonkompetisis.unggahsertifikat, prestasinonkompetisis.unggahsurattugas, prestasinonkompetisis.unggahfoto, prestasinonkompetisis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasinonkompetisis.idmhs").Where("status = ?", "setuju").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPrestasinonTST(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result []Prestasinonkompetisi

	models.DB.Table("prestasinonkompetisis").Select("prestasinonkompetisis.idmhs, prestasinonkompetisis.namakegiatan, mahasiswas.nama, prestasinonkompetisis.idprestasinon, prestasinonkompetisis.namapenyelenggaraan,	 prestasinonkompetisis.tempatkegiatan, prestasinonkompetisis.tanggalawal, prestasinonkompetisis.tanggalakhir, prestasinonkompetisis.unggahsertifikat, prestasinonkompetisis.unggahsurattugas, prestasinonkompetisis.unggahfoto, prestasinonkompetisis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasinonkompetisis.idmhs").Where("status = ?", "").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

//Get 1 Row
func GetOnePengabdian(c *gin.Context) {
	var pengabdian Pengabdian
	parpengabdian := c.Param("idpengabdian")

	//models.DB.Where("idpengabdian = ?", parpengabdian).First(&pengabdian)
	models.DB.Table("pengabdians").Select("pengabdians.idpengabdian, pengabdians.idmhs, pengabdians.namaprogram, pengabdians.tahunkegiatan, pengabdians.unggahsurat, pengabdians.status, mahasiswas.idmhs, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = pengabdians.idmhs").Where("idpengabdian = ?", parpengabdian).Scan(&pengabdian)

	c.JSON(http.StatusOK, gin.H{"data": pengabdian})
}

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

func GetOnePrestasinon(c *gin.Context) {
	//var prestasi []models.Prestasi
	var result Prestasinonkompetisi
	parprestasinon := c.Param("id_prestasinon")

	models.DB.Table("prestasinonkompetisis").Select("prestasinonkompetisis.idmhs, prestasinonkompetisis.namakegiatan, mahasiswas.nama, prestasinonkompetisis.idprestasinon, prestasinonkompetisis.namapenyelenggaraan,	 prestasinonkompetisis.tempatkegiatan, prestasinonkompetisis.tanggalawal, prestasinonkompetisis.tanggalakhir, prestasinonkompetisis.unggahsertifikat, prestasinonkompetisis.unggahsurattugas, prestasinonkompetisis.unggahfoto, prestasinonkompetisis.status").
		Joins("left join mahasiswas on mahasiswas.idmhs = prestasinonkompetisis.idmhs").Where("idprestasinon = ?", parprestasinon).Scan(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetOneOrganisasi(c *gin.Context) {
	//var prestasi models.Prestasi
	var result ResultOrganisasi
	parorganisasi := c.Param("id_organisasi")

	//models.DB.Where("id_prestasi = ?", parprestasi).First(&prestasi)

	models.DB.Table("organisasis").Select("organisasis.idorganisasi, organisasis.idmhs, organisasis.namaorganisasi, organisasis.jabatanmhs, organisasis.periode, organisasis.unggahsk, organisasis.status, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = organisasis.idmhs").Where("idorganisasi = ?", parorganisasi).Scan(&result)

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

	//START png
	readerFoto := base64.NewDecoder(base64.StdEncoding, strings.NewReader(prestasi.Unggahfoto))
	m, formatString, err := image.Decode(readerFoto)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	_, err = os.Stat("web/dist/image/")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("web/dist/image/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	pngFilename := "web/dist/image/" + prestasi.Namakegiatan + "_foto.png"

	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", pngFilename, "created")

	prestasi.Unggahfoto = prestasi.Namakegiatan + "_foto.png"
	//END png

	//START pdf
	readerSertifikat, err := base64.StdEncoding.DecodeString(prestasi.Unggahsertifikat)
	if err != nil {
		log.Fatal(err)
	}

	PdfSertifikatFilename := "web/dist/image/" + prestasi.Namakegiatan + "_SE.pdf"

	g, err := os.OpenFile(PdfSertifikatFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := g.Write(readerSertifikat); err != nil {
		panic(err)
	}
	if err := g.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSertifikatFilename, "created")
	prestasi.Unggahsertifikat = prestasi.Namakegiatan + "_SE.pdf"
	//end pdf

	//START pdf
	readerSurattugas, err := base64.StdEncoding.DecodeString(prestasi.Unggahsurattugas)
	if err != nil {
		log.Fatal(err)
	}

	PdfSTFilename := "web/dist/image/" + prestasi.Namakegiatan + "_ST.pdf"

	h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := h.Write(readerSurattugas); err != nil {
		panic(err)
	}
	if err := h.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSTFilename, "created")
	prestasi.Unggahsurattugas = prestasi.Namakegiatan + "_ST.pdf"
	//end pdf

	result := models.DB.Create(&prestasi)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func AddPrestasinon(c *gin.Context) {
	var prestasinon models.Prestasinonkompetisi
	//var prestasiT Prestasi

	if err := c.Bind(&prestasinon); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	//START png
	readerFoto := base64.NewDecoder(base64.StdEncoding, strings.NewReader(prestasinon.Unggahfoto))
	m, formatString, err := image.Decode(readerFoto)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	_, err = os.Stat("web/dist/image/")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("web/dist/image/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	pngFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_foto.png"

	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", pngFilename, "created")

	prestasinon.Unggahfoto = prestasinon.Namakegiatan + "_foto.png"
	//END png

	//START pdf
	readerSertifikat, err := base64.StdEncoding.DecodeString(prestasinon.Unggahsertifikat)
	if err != nil {
		log.Fatal(err)
	}

	PdfSertifikatFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_SE.pdf"

	g, err := os.OpenFile(PdfSertifikatFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := g.Write(readerSertifikat); err != nil {
		panic(err)
	}
	if err := g.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSertifikatFilename, "created")
	prestasinon.Unggahsertifikat = prestasinon.Namakegiatan + "_SE.pdf"
	//end pdf

	//START pdf
	readerSurattugas, err := base64.StdEncoding.DecodeString(prestasinon.Unggahsurattugas)
	if err != nil {
		log.Fatal(err)
	}

	PdfSTFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_ST.pdf"

	h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := h.Write(readerSurattugas); err != nil {
		panic(err)
	}
	if err := h.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSTFilename, "created")
	prestasinon.Unggahsurattugas = prestasinon.Namakegiatan + "_ST.pdf"
	//end pdf

	result := models.DB.Create(&prestasinon)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func AddOrganisasi(c *gin.Context) {
	var organisasi models.Organisasi
	//var prestasiT Prestasi

	if err := c.Bind(&organisasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	//START pdf
	readerSK, err := base64.StdEncoding.DecodeString(organisasi.Unggahsk)
	if err != nil {
		log.Fatal(err)
	}

	PdfSTFilename := "web/dist/image/" + organisasi.Namaorganisasi + "_SK.pdf"

	h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := h.Write(readerSK); err != nil {
		panic(err)
	}
	if err := h.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSTFilename, "created")
	organisasi.Unggahsk = organisasi.Namaorganisasi + "_SK.pdf"
	//end pdf

	result := models.DB.Create(&organisasi)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
}

func AddPengabdian(c *gin.Context) {
	var pengabdian models.Pengabdian
	//var prestasiT Prestasi

	if err := c.Bind(&pengabdian); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	//START pdf
	readerSK, err := base64.StdEncoding.DecodeString(pengabdian.Unggahsurat)
	if err != nil {
		log.Fatal(err)
	}

	PdfSTFilename := "web/dist/image/" + pengabdian.Namaprogram + "_Surat.pdf"

	h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := h.Write(readerSK); err != nil {
		panic(err)
	}
	if err := h.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Pdf file", PdfSTFilename, "created")
	pengabdian.Unggahsurat = pengabdian.Namaprogram + "_Surat.pdf"
	//end pdf

	result := models.DB.Create(&pengabdian)
	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
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

	//START png
	if len(prestasi.Unggahfoto) > 0 {
		readerFoto := base64.NewDecoder(base64.StdEncoding, strings.NewReader(prestasi.Unggahfoto))
		m, formatString, err := image.Decode(readerFoto)
		if err != nil {
			log.Fatal(err)
		}

		bounds := m.Bounds()
		fmt.Println(bounds, formatString)

		_, err = os.Stat("web/dist/image/")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("web/dist/image/", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		pngFilename := "web/dist/image/" + prestasi.Namakegiatan + "_foto.png"

		f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = png.Encode(f, m)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Png file", pngFilename, "created")

		prestasi.Unggahfoto = prestasi.Namakegiatan + "_foto.png"
		//END png
	}

	if len(prestasi.Unggahsertifikat) > 0 {
		//START pdf
		readerSertifikat, err := base64.StdEncoding.DecodeString(prestasi.Unggahsertifikat)
		if err != nil {
			log.Fatal(err)
		}

		PdfSertifikatFilename := "web/dist/image/" + prestasi.Namakegiatan + "_SE.pdf"

		g, err := os.OpenFile(PdfSertifikatFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := g.Write(readerSertifikat); err != nil {
			panic(err)
		}
		if err := g.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSertifikatFilename, "created")
		prestasi.Unggahsertifikat = prestasi.Namakegiatan + "_SE.pdf"
		//end pdf
	}

	if len(prestasi.Unggahsurattugas) > 0 {
		//START pdf
		readerSurattugas, err := base64.StdEncoding.DecodeString(prestasi.Unggahsurattugas)
		if err != nil {
			log.Fatal(err)
		}

		PdfSTFilename := "web/dist/image/" + prestasi.Namakegiatan + "_ST.pdf"

		h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := h.Write(readerSurattugas); err != nil {
			panic(err)
		}
		if err := h.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSTFilename, "created")
		prestasi.Unggahsurattugas = prestasi.Namakegiatan + "_ST.pdf"
		//end pdf
	}

	result := models.DB.Where("Idprestasi = ?", prestasi.Idprestasi).Updates(&prestasi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditPrestasinon(c *gin.Context) {
	var prestasinon models.Prestasinonkompetisi

	if err := c.Bind(&prestasinon); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parprestasinon string
	parprestasinon = c.Param("id_prestasinon")
	parprestasinonConvert, err := strconv.Atoi(parprestasinon)

	if err != nil {
		fmt.Println("eror")
	}
	prestasinon.Idprestasinon = parprestasinonConvert

	//START png
	if len(prestasinon.Unggahfoto) > 0 {
		readerFoto := base64.NewDecoder(base64.StdEncoding, strings.NewReader(prestasinon.Unggahfoto))
		m, formatString, err := image.Decode(readerFoto)
		if err != nil {
			log.Fatal(err)
		}

		bounds := m.Bounds()
		fmt.Println(bounds, formatString)

		_, err = os.Stat("web/dist/image/")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("web/dist/image/", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		pngFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_foto.png"

		f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = png.Encode(f, m)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Png file", pngFilename, "created")

		prestasinon.Unggahfoto = prestasinon.Namakegiatan + "_foto.png"
		//END png
	}
	//START pdf

	if len(prestasinon.Unggahsertifikat) > 0 {
		readerSertifikat, err := base64.StdEncoding.DecodeString(prestasinon.Unggahsertifikat)
		if err != nil {
			log.Fatal(err)
		}

		PdfSertifikatFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_SE.pdf"

		g, err := os.OpenFile(PdfSertifikatFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := g.Write(readerSertifikat); err != nil {
			panic(err)
		}
		if err := g.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSertifikatFilename, "created")
		prestasinon.Unggahsertifikat = prestasinon.Namakegiatan + "_SE.pdf"
		//end pdf
	}

	if len(prestasinon.Unggahsurattugas) > 0 {
		//START pdf
		readerSurattugas, err := base64.StdEncoding.DecodeString(prestasinon.Unggahsurattugas)
		if err != nil {
			log.Fatal(err)
		}

		PdfSTFilename := "web/dist/image/" + prestasinon.Namakegiatan + "_ST.pdf"

		h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := h.Write(readerSurattugas); err != nil {
			panic(err)
		}
		if err := h.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSTFilename, "created")
		prestasinon.Unggahsurattugas = prestasinon.Namakegiatan + "_ST.pdf"
		//end pdf
	}

	result := models.DB.Where("idprestasinon = ?", parprestasinon).Updates(&prestasinon)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditOrganisasi(c *gin.Context) {
	var organisasi models.Organisasi

	if err := c.Bind(&organisasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parorganisasi string
	parorganisasi = c.Param("id_organisasi")

	parorganisasiConvert, err := strconv.Atoi(parorganisasi)

	if err != nil {
		fmt.Println("eror")
	}
	organisasi.Idorganisasi = parorganisasiConvert

	if len(organisasi.Unggahsk) > 0 {
		//START pdf
		readerSK, err := base64.StdEncoding.DecodeString(organisasi.Unggahsk)
		if err != nil {
			log.Fatal(err)
		}

		PdfSTFilename := "web/dist/image/" + organisasi.Namaorganisasi + "_SK.pdf"

		h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := h.Write(readerSK); err != nil {
			panic(err)
		}
		if err := h.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSTFilename, "created")
		organisasi.Unggahsk = organisasi.Namaorganisasi + "_SK.pdf"
		//end pdf
	}

	result := models.DB.Where("idorganisasi = ?", organisasi.Idorganisasi).Updates(&organisasi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditPengabdian(c *gin.Context) {
	var pengabdian models.Pengabdian

	if err := c.Bind(&pengabdian); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parpengabdian string
	parpengabdian = c.Param("idpengabdian")

	parpengabdianConvert, err := strconv.Atoi(parpengabdian)

	if err != nil {
		fmt.Println("eror")
	}
	pengabdian.Idpengabdian = parpengabdianConvert

	if len(pengabdian.Unggahsurat) > 0 {
		//START pdf
		readerSK, err := base64.StdEncoding.DecodeString(pengabdian.Unggahsurat)
		if err != nil {
			log.Fatal(err)
		}

		PdfSTFilename := "web/dist/image/" + pengabdian.Namaprogram + "_Surat.pdf"

		h, err := os.OpenFile(PdfSTFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		if _, err := h.Write(readerSK); err != nil {
			panic(err)
		}
		if err := h.Sync(); err != nil {
			panic(err)
		}
		fmt.Println("Pdf file", PdfSTFilename, "created")
		pengabdian.Unggahsurat = pengabdian.Namaprogram + "_Surat.pdf"
		//end pdf
	}

	result := models.DB.Where("idpengabdian = ?", pengabdian.Idpengabdian).Updates(&pengabdian)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditSetujuPrestasi(c *gin.Context) {
	var prestasi models.Prestasi
	//var prestasit models.Prestasi

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

	result := models.DB.Model(&models.Prestasi{}).Where("Idprestasi = ?", prestasi.Idprestasi).Update("status", "setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditTSetujuPrestasi(c *gin.Context) {
	var prestasi models.Prestasi
	//var prestasit models.Prestasi

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

	result := models.DB.Model(&models.Prestasi{}).Where("Idprestasi = ?", prestasi.Idprestasi).Update("status", "tidak setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditSetujuOrganisasi(c *gin.Context) {
	var organisasi models.Organisasi
	//var prestasit models.Prestasi

	if err := c.Bind(&organisasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parorganisasi string
	parorganisasi = c.Param("id_organisasi")

	parorganisasiConvert, err := strconv.Atoi(parorganisasi)

	if err != nil {
		fmt.Println("eror")
	}
	organisasi.Idorganisasi = parorganisasiConvert

	result := models.DB.Model(&models.Organisasi{}).Where("idorganisasi = ?", organisasi.Idorganisasi).Update("status", "setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditTSetujuOrganisasi(c *gin.Context) {
	var organisasi models.Organisasi
	//var prestasit models.Prestasi

	if err := c.Bind(&organisasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parorganisasi string
	parorganisasi = c.Param("id_organisasi")

	parorganisasiConvert, err := strconv.Atoi(parorganisasi)

	if err != nil {
		fmt.Println("eror")
	}
	organisasi.Idorganisasi = parorganisasiConvert

	result := models.DB.Model(&models.Organisasi{}).Where("idorganisasinon = ?", organisasi.Idorganisasi).Update("status", "tidak setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditSetujuPrestasinon(c *gin.Context) {
	var prestasinon models.Prestasinonkompetisi
	//var prestasit models.Prestasi

	if err := c.Bind(&prestasinon); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parprestasinon string
	parprestasinon = c.Param("id_prestasinon")

	parprestasinonConvert, err := strconv.Atoi(parprestasinon)

	if err != nil {
		fmt.Println("eror")
	}
	prestasinon.Idprestasinon = parprestasinonConvert

	result := models.DB.Model(&models.Prestasinonkompetisi{}).Where("Idprestasinon = ?", prestasinon.Idprestasinon).Update("status", "setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditTSetujuPrestasinon(c *gin.Context) {
	var prestasinon models.Prestasinonkompetisi
	//var prestasit models.Prestasi

	if err := c.Bind(&prestasinon); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parprestasinon string
	parprestasinon = c.Param("id_prestasinon")

	parprestasinonConvert, err := strconv.Atoi(parprestasinon)

	if err != nil {
		fmt.Println("eror")
	}
	prestasinon.Idprestasinon = parprestasinonConvert

	result := models.DB.Model(&models.Prestasinonkompetisi{}).Where("Idprestasinon = ?", prestasinon.Idprestasinon).Update("status", "tidak setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditSetujuPengabdian(c *gin.Context) {
	var pengabdian models.Pengabdian
	//var prestasit models.Prestasi

	if err := c.Bind(&pengabdian); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parpengabdian string
	parpengabdian = c.Param("idpengabdian")

	parpengabdianConvert, err := strconv.Atoi(parpengabdian)

	if err != nil {
		fmt.Println("eror")
	}
	pengabdian.Idpengabdian = parpengabdianConvert

	result := models.DB.Model(&models.Pengabdian{}).Where("idpengabdian = ?", pengabdian.Idpengabdian).Update("status", "setuju")

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func EditTSetujuPengabdian(c *gin.Context) {
	var pengabdian models.Pengabdian
	//var prestasit models.Prestasi

	if err := c.Bind(&pengabdian); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	var parpengabdian string
	parpengabdian = c.Param("idpengabdian")

	parpengabdianConvert, err := strconv.Atoi(parpengabdian)

	if err != nil {
		fmt.Println("eror")
	}
	pengabdian.Idpengabdian = parpengabdianConvert

	result := models.DB.Model(&models.Pengabdian{}).Where("idpengabdian = ?", pengabdian.Idpengabdian).Update("status", "tidak setuju")

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

func DeletePrestasinon(c *gin.Context) {
	var prestasinon models.Prestasinonkompetisi

	if err := c.Bind(&prestasinon); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	parprestasinon := c.Param("id_prestasinon")

	result := models.DB.Where("idprestasinon=?", parprestasinon).Delete(&prestasinon)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func DeleteOrganisasi(c *gin.Context) {
	var organisasi models.Organisasi

	if err := c.Bind(&organisasi); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	parorganisasi := c.Param("id_organisasi")

	result := models.DB.Where("idorganisasi=?", parorganisasi).Delete(&organisasi)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

func DeletePengabdian(c *gin.Context) {
	var pengabdian models.Pengabdian

	if err := c.Bind(&pengabdian); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	parpengabdian := c.Param("idpengabdian")

	result := models.DB.Where("idpengabdian=?", parpengabdian).Delete(&pengabdian)

	chech := result.RowsAffected
	if chech == 1 {
		c.JSON(http.StatusOK, gin.H{"Status": "Berhasil"})
	}
}

//upload
func Upload(c *gin.Context) {
	var foto models.Foto

	if err := c.Bind(&foto); err != nil {
		c.JSON(http.StatusOK, gin.H{"Eror": err.Error()})
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(foto.Value))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	_, err = os.Stat("web/dist/image/")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("web/dist/image/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	//parm := c.Param("username")
	//Encode from image format to writer
	pngFilename := "web/dist/image/test.jpg"

	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	//model.DB.Model(&model.User{}).Where("username = ?", parm).Update("foto", pngFilename)

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", pngFilename, "created")
}

func Pdfupload(c *gin.Context) {
	var foto models.Foto

	if err := c.Bind(&foto); err != nil {
		c.JSON(http.StatusOK, gin.H{"Eror": err.Error()})
	}

	reader, err := base64.StdEncoding.DecodeString(foto.Value)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("web/dist/image/")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("web/dist/image/", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	//parm := c.Param("username")
	//Encode from image format to writer
	pngFilename := "web/dist/image/test.pdf"

	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := f.Write(reader); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("Png file", pngFilename, "created")
}

func EvaluasiUnivPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiUnivPrestasi

	models.DB.Raw("SELECT year(tanggalawal) as year, count(*) as prestasi from prestasis group by year(tanggalawal)").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiUnivPrestasiNon(c *gin.Context) {
	var evaluasi []models.EvaluasiUnivPrestasinon

	models.DB.Raw("SELECT year(tanggalawal) as year, count(*) as prestasinon from prestasinonkompetisis group by year(tanggalawal)").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiUnivPengabdian(c *gin.Context) {
	var evaluasi []models.EvaluasiUnivPengabdian

	models.DB.Raw("SELECT year(tanggalawal) as year, count(*) as pengabdian from pengabdians group by year(tanggalawal)").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasEkoPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasEkonomi
	//models.DB.Raw("SELECT year(p.tanggalawal) as year, count(*) as ekonomi FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Ekonomi dan Bisnis' GROUP by p.tanggalawal").Scan(&evaluasi)
	models.DB.Raw("select a.year as year, sum(a.ekonomi) as ekonomi from (SELECT year(p.tanggalawal) as year, count(*) as ekonomi FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Ekonomi dan Bisnis' GROUP by p.tanggalawal) a group by a.year").Scan((&evaluasi))
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasTeknikPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasTeknik
	models.DB.Raw("select a.year as year, sum(a.teknik) as teknik from (SELECT year(p.tanggalawal) as year, count(*) as teknik FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Teknik' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasPertaPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasPertanian
	models.DB.Raw("select a.year as year, sum(a.pertanian) as pertanian from (SELECT year(p.tanggalawal) as year, count(*) as pertanian FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Pertanian' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasFisipPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasFisip
	models.DB.Raw("select a.year as year, sum(a.fisip) as fisip from (SELECT year(p.tanggalawal) as year, count(*) as fisip FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Ilmu Sosial dan Ilmu Politik' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasHukumPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasHukum
	models.DB.Raw("select a.year as year, sum(a.hukum) as hukum from (SELECT year(p.tanggalawal) as year, count(*) as hukum FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Hukum' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasGuruPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasGuru
	models.DB.Raw("select a.year as year, sum(a.keguruan) as keguruan from (SELECT year(p.tanggalawal) as year, count(*) as keguruan FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Keguruan dan Ilmu Pendidikan' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasMtkPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasMtk
	models.DB.Raw("select a.year as year, sum(a.matematika) as matematika from (SELECT year(p.tanggalawal) as year, count(*) as matematika FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Matematika dan Ilmu Pengetahuan Alam' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func EvaluasiFakultasDokterPrestasi(c *gin.Context) {
	var evaluasi []models.EvaluasiFakultasKedokteran
	models.DB.Raw("select a.year as year, sum(a.kedokteran) as kedokteran from (SELECT year(p.tanggalawal) as year, count(*) as kedokteran FROM mahasiswas m INNER join prestasis p on m.idmhs = p.idmhs where m.fakultas='Fakultas Kedokteran' GROUP by p.tanggalawal) a group by a.year").Scan(&evaluasi)
	c.JSON(http.StatusOK, gin.H{"data": evaluasi})
}

func GetCountMahasiswa(c *gin.Context) {
	var ResultCount []ResultCount
	models.DB.Raw("SELECT COUNT(*) as jumlah from mahasiswas").Scan(&ResultCount)
	c.JSON(http.StatusOK, gin.H{"data": ResultCount})
}

func GetCountPrestasi(c *gin.Context) {
	var ResultCount []ResultCount
	models.DB.Raw("SELECT COUNT(*) as jumlah from prestasis").Scan(&ResultCount)
	c.JSON(http.StatusOK, gin.H{"data": ResultCount})
}

func GetCountPrestasinon(c *gin.Context) {
	var ResultCount []ResultCount
	models.DB.Raw("SELECT COUNT(*) as jumlah from prestasinonkompetisis").Scan(&ResultCount)
	c.JSON(http.StatusOK, gin.H{"data": ResultCount})
}

func GetCountPengabdian(c *gin.Context) {
	var ResultCount []ResultCount
	models.DB.Raw("SELECT COUNT(*) as jumlah from pengabdians").Scan(&ResultCount)
	c.JSON(http.StatusOK, gin.H{"data": ResultCount})
}
