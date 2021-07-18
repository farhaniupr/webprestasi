package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"prestasi/app/controller"
	"prestasi/app/middleware"
)

// CORS Middleware
func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	router := gin.Default()

	router.Use(CORS)

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
	router.GET("/lihatmahasiswa", middleware.Auth, controller.GetMahasiswa)
	router.GET("/lihatorganisasi", middleware.Auth, controller.GetOrganisasi)
	router.GET("/lihatpengabdian", middleware.Auth, controller.GetPengabdian)
	router.GET("/lihatprestasinon", middleware.Auth, controller.GetPrestasinon)
	router.GET("/lihatprestasi", middleware.Auth, controller.GetPrestasi)
	router.GET("/lihatprodi", middleware.Auth, controller.GetProdi)

	//get for dashboard
	router.GET("/lihatprestasidash", middleware.Auth, controller.GetPrestasiConditionTS)

	//get for conidition
	router.GET("/lihatprestasi_st", middleware.Auth, controller.GetPrestasiConditionST)
	router.GET("/lihatprestasi_tst", middleware.Auth, controller.GetPrestasiConditionTST)

	router.GET("/lihatpengabdian_tst", middleware.Auth, controller.GetPengabdianTST)
	router.GET("/lihatpengabdian_st", middleware.Auth, controller.GetPengabdianST)

	router.GET("/lihatprestasinon_st", middleware.Auth, controller.GetPrestasinonST)
	router.GET("/lihatprestasinon_tst", middleware.Auth, controller.GetPrestasinonTST)

	router.GET("/lihatorganisasi_st", middleware.Auth, controller.GetOrganisasiST)
	router.GET("/lihatorganisasi_tst", middleware.Auth, controller.GetOrganisasiTST)

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

	router.GET("/countmahasiswa", controller.GetCountMahasiswa)
	router.GET("/countprestasi", controller.GetCountPrestasi)
	router.GET("/countprestasinon", controller.GetCountPrestasinon)
	router.GET("/countpengabdian", controller.GetCountPengabdian)
	//REST API upload -- sdh digabung
	//router.POST("/upload", controller.Upload)
	//router.POST("/uploadpdf", controller.Pdfupload)

	//API utk client luar
	router.GET("/api.evaluasiunivprestasi", controller.EvaluasiUnivPrestasi)
	router.GET("/api.evaluasiunivpengabdian", controller.EvaluasiUnivPengabdian)
	router.GET("/api.evaluasiunivprestasinon", controller.EvaluasiUnivPrestasiNon)
	router.GET("/api.evaluasifakultasKedokteran", controller.EvaluasiFakultasDokterPrestasi)
	router.GET("/api.evaluasifakultasFisip", controller.EvaluasiFakultasFisipPrestasi)
	router.GET("/api.evaluasifakultasMtk", controller.EvaluasiFakultasMtkPrestasi)
	router.GET("/api.evaluasifakultasGuru", controller.EvaluasiFakultasGuruPrestasi)
	router.GET("/api.evaluasifakultasEkonomi", controller.EvaluasiFakultasEkoPrestasi)
	router.GET("/api.evaluasifakultasPertanian", controller.EvaluasiFakultasPertaPrestasi)
	router.GET("/api.evaluasifakultasHukum", controller.EvaluasiFakultasHukumPrestasi)
	router.GET("/api.evaluasifakultasTeknik", controller.EvaluasiFakultasTeknikPrestasi)
	router.GET("/api.lihatmahasiswa", controller.GetMahasiswa)
	router.GET("/api.lihatorganisasi", controller.GetOrganisasi)
	router.GET("/api.lihatpengabdian", controller.GetPengabdian)
	router.GET("/api.lihatprestasinon", controller.GetPrestasinon)
	router.GET("/api.lihatprestasi", controller.GetPrestasi)
	router.GET("/api.lihatprodi", controller.GetProdi)

	router.POST("/login", controller.Sign)
	router.POST("/signup", controller.SignUp)

	router.Run(":8085")
}
