package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	"prestasi/app/controller"
)

func main() {
	router := gin.Default()
	cfg := cors.DefaultConfig()

	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AddAllowedHeaders("*")

	cfg.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	cfg.AllowedHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

	//cfg.AllowAllOrigins =

	router.Use(cors.New(cfg))
	//router.Use(gin.Logger())
	///router.Use(gin.Recovery())

	//get
	router.GET("/lihatmahasiswa", controller.GetMahasiswa)
	router.GET("/lihatorganisasi", controller.GetOrganisasi)
	router.GET("/lihatpengabdian", controller.GetPengabdian)
	router.GET("/lihatprestasinon", controller.GetPrestasiNon)
	router.GET("/lihatprestasi", controller.GetPrestasi)
	router.GET("/lihatprodi", controller.GetProdi)

	//get 1 row
	router.GET("/lihatmahasiswa/:idmhs", controller.GetOneMahasiswa)
	router.GET("/lihatprodi/:kode_prodi", controller.GetOneProdi)
	router.GET("/lihatprestasi/:id_prestasi", controller.GetOnePrestasi)

	//insert
	router.POST("/addmahasiswa", controller.AddMahasiswa)
	router.POST("/addprodi", controller.AddProdi)
	router.POST("/addprestasi", controller.AddPrestasi)

	//edit
	router.PUT("/editmahasiswa/:idmhs", controller.EditMahasiswa)
	router.PUT("/editprodi/:kode_prodi", controller.EditProdi)
	router.PUT("/editprestasi/:id_prestasi", controller.EditPrestasi)

	//edit > setuju/tidak
	router.PUT("/editsetujuprestasi/:id_prestasi", controller.EditSetujuPrestasi)
	router.PUT("/edittsetujuprestasi/:id_prestasi", controller.EditTSetujuPrestasi)
	//delete
	router.DELETE("/deletemahasiswa/:idmhs", controller.DeleteMahasiswa)
	router.DELETE("/deleteprodi/:kode_prodi", controller.DeleteProdi)
	router.DELETE("/deleteprestasi/:id_prestasi", controller.DeletePrestasi)

	router.POST("/upload", controller.Upload)
	router.POST("/uploadpdf", controller.Pdfupload)

	router.Run(":8085")
}
