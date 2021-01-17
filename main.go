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

	cfg.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	cfg.AllowedHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	router.Use(cors.New(cfg))
	//router.Use(gin.Logger())
	///router.Use(gin.Recovery())

	router.GET("/lihatmahasiswa", controller.GetMahasiswa)
	router.GET("/lihatorganisasi", controller.GetOrganisasi)
	router.GET("/lihatpengabdian", controller.GetPengabdian)
	router.GET("/lihatprestasinon", controller.GetPrestasiNon)
	router.GET("/lihatprestasi", controller.GetPrestasi)
	router.Run(":8085")
}
