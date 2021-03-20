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

	//get evaluasi
	router.GET("/evaluasiunivprestasi", controller.EvaluasiUnivPrestasi)
	router.GET("/evaluasiunivpengabdian", controller.EvaluasiUnivPengabdian)
	router.GET("/evaluasiunivprestasinon", controller.EvaluasiUnivPrestasiNon)
	//router.GET("/evaluasiunivorganisasi", controller.EvaluasiUnivOrganisasi)

	router.GET("/evaluasifakultasKedokteran", controller.EvaluasiFakultasDokterPrestasi)
	router.GET("/evaluasifakultasFisip", controller.EvaluasiFakultasFisipPrestasi)
	router.GET("/evaluasifakultasMtk", controller.EvaluasiFakultasMtkPrestasi)
	router.GET("/evaluasifakultasGuru", controller.EvaluasiFakultasGuruPrestasi)
	router.GET("/evaluasifakultasEkonomi", controller.EvaluasiFakultasEkoPrestasi)
	router.GET("/evaluasifakultasPertanian", controller.EvaluasiFakultasPertaPrestasi)
	router.GET("/evaluasifakultasHukum", controller.EvaluasiFakultasHukumPrestasi)
	router.GET("/evaluasifakultasTeknik", controller.EvaluasiFakultasTeknikPrestasi)
	//router.GET("/evaluasifakultas", controla)

	//get
	router.GET("/lihatmahasiswa", controller.GetMahasiswa)
	router.GET("/lihatorganisasi", middleware.Auth, controller.GetOrganisasi)
	router.GET("/lihatpengabdian", middleware.Auth, controller.GetPengabdian)
	router.GET("/lihatprestasinon", middleware.Auth, controller.GetPrestasinon)
	router.GET("/lihatprestasi", middleware.Auth, controller.GetPrestasi)
	router.GET("/lihatprodi", middleware.Auth, controller.GetProdi)

	//get 1 row
	router.GET("/lihatmahasiswa/:idmhs", middleware.Auth, controller.GetOneMahasiswa)
	router.GET("/lihatprodi/:kode_prodi", middleware.Auth, controller.GetOneProdi)
	router.GET("/lihatprestasi/:id_prestasi", middleware.Auth, controller.GetOnePrestasi)
	router.GET("/lihatprestasinon/:id_prestasinon", middleware.Auth, controller.GetOnePrestasinon)
	router.GET("/lihatorganisasi/:id_organisasi", middleware.Auth, controller.GetOneOrganisasi)
	router.GET("/lihatpengabdian/:idpengabdian", middleware.Auth, controller.GetOnePengabdian)

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
