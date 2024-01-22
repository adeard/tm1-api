package main

import (
	"log"
	"net/http"
	"os"
	"tm1-api/docs"
	"tm1-api/modules/tm1"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Host 202.158.14.235:4222
// @title API SWAGGER FOR TM1 API SERVICE
// @version 1.0.0
// @description TM1 API SERVICE
// @termsOfService http://swagger.io/terms/

// @contact.name ICT INDOAGRI
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath

func main() {
	// db := config.Connect()

	docs.SwaggerInfo.BasePath = "/Tm1Api"

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.GET("Tm1Api/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":         "TM1 API Service",
			"documentation": "/swagger/index.html",
		})
	})

	router.GET("Tm1Api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("Tm1Api/api/v1")

	tm1.NewTm1Handler(v1, tm1.Tm1Registry())

	// router.Run(":86")

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "86"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	// Menampilkan log koneksi sukses
	log.Println("App Service run in port:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// Menampilkan log ketika koneksi gagal
		log.Fatal("Connection Fail -> port "+port+":", err)
	}
}
