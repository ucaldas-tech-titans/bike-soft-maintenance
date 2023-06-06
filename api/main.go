package main

import (
	"github.com/anything-devs/bike-soft-rest-api/configs"
	"github.com/anything-devs/bike-soft-rest-api/models"
	"github.com/anything-devs/bike-soft-rest-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	conectarDB    = configs.ConectarBD
	migrateTables = autoMigrateTables
	run           = runServer
)

func main() {
	router := gin.Default()
	DB := conectarDB()

	migrateTables(DB)
	run(router, DB)
}

func autoMigrateTables(DB *gorm.DB) {
	DB.Table("productos").AutoMigrate(models.Producto{})
	DB.Table("categorias").AutoMigrate(models.Categoria{})
}

func runServer(router *gin.Engine, DB *gorm.DB) {
	routes.Rutas(router, DB)
	// Configurar opciones CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Run(":8083")
}
