package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	"prestasi/app/controller"
	"prestasi/app/middleware"
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
	router.GET("/lihatmahasiswa", middleware.Auth, controller.GetMahasiswa)
	router.GET("/lihatorganisasi", controller.GetOrganisasi)
	router.GET("/lihatpengabdian", controller.GetPengabdian)
	router.GET("/lihatprestasinon", controller.GetPrestasinon)
	router.GET("/lihatprestasi", controller.GetPrestasi)
	router.GET("/lihatprodi", controller.GetProdi)

	//get 1 row
	router.GET("/lihatmahasiswa/:idmhs", controller.GetOneMahasiswa)
	router.GET("/lihatprodi/:kode_prodi", controller.GetOneProdi)
	router.GET("/lihatprestasi/:id_prestasi", controller.GetOnePrestasi)
	router.GET("/lihatprestasinon/:id_prestasinon", controller.GetOnePrestasinon)
	router.GET("/lihatorganisasi/:id_organisasi", controller.GetOneOrganisasi)
	router.GET("/lihatpengabdian/:idpengabdian", controller.GetOnePengabdian)

	//insert
	router.POST("/addmahasiswa", controller.AddMahasiswa)
	router.POST("/addprodi", controller.AddProdi)
	router.POST("/addprestasi", controller.AddPrestasi)
	router.POST("/addprestasinon", controller.AddPrestasinon)
	router.POST("/addorganisasi", controller.AddOrganisasi)
	router.POST("/addpengabdian", controller.AddPengabdian)

	//edit
	router.PUT("/editmahasiswa/:idmhs", controller.EditMahasiswa)
	router.PUT("/editprodi/:kode_prodi", controller.EditProdi)
	router.PUT("/editprestasi/:id_prestasi", controller.EditPrestasi)
	router.PUT("/editprestasinon/:id_prestasinon", controller.EditPrestasinon)
	router.PUT("/editorganisasi/:id_organisasi", controller.EditOrganisasi)
	router.PUT("/editpengabdian/:idpengabdian", controller.EditPengabdian)

	//edit > setuju/tidak
	router.PUT("/editsetujuprestasi/:id_prestasi", controller.EditSetujuPrestasi)
	router.PUT("/edittsetujuprestasi/:id_prestasi", controller.EditTSetujuPrestasi)

	router.PUT("/editsetujuorganisasi/:id_organisasi", controller.EditSetujuOrganisasi)
	router.PUT("/edittsetujuorganisasi/:id_organisasi", controller.EditTSetujuOrganisasi)

	router.PUT("/editsetujuprestasinon/:id_prestasinon", controller.EditSetujuPrestasinon)
	router.PUT("/edittsetujuprestasinon/:id_prestasinon", controller.EditTSetujuPrestasinon)

	router.PUT("/editsetujupengabdian/:idpengabdian", controller.EditSetujuPengabdian)
	router.PUT("/edittsetujupengabdian/:idpengabdian", controller.EditTSetujuPengabdian)

	//delete
	router.DELETE("/deletemahasiswa/:idmhs", controller.DeleteMahasiswa)
	router.DELETE("/deleteprodi/:kode_prodi", controller.DeleteProdi)
	router.DELETE("/deleteprestasi/:id_prestasi", controller.DeletePrestasi)
	router.DELETE("/deleteprestasinon/:id_prestasinon", controller.DeletePrestasinon)
	router.DELETE("/deleteorganisasi/:id_organisasi", controller.DeleteOrganisasi)
	router.DELETE("/deletepengabdian/:idpengabdian", controller.DeletePengabdian)

	//REST API upload -- sdh digabung
	//router.POST("/upload", controller.Upload)
	//router.POST("/uploadpdf", controller.Pdfupload)

	router.POST("/login", controller.Sign)
	router.POST("/signup", controller.SignUp)

	router.Run(":8085")
}
