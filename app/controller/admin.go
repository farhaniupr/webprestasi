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
	var organisasi []ResultOrganisasi
	//models.DB.Find(&organisasi)
	models.DB.Table("organisasis").Select("organisasis.idorganisasi, organisasis.idmhs, organisasis.namaorganisasi, organisasis.jabatanmhs, organisasis.periode, organisasis.unggahsk, organisasis.status, mahasiswas.nama").
		Joins("left join mahasiswas on mahasiswas.idmhs = organisasis.idmhs").Scan(&organisasi)
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

	result := models.DB.Where("idorganisasi = ?", organisasi.Idorganisasi).Updates(&organisasi)

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

	result := models.DB.Model(&models.Organisasi{}).Where("idorganisasi = ?", organisasi.Idorganisasi).Update("status", "tidak setuju")

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
