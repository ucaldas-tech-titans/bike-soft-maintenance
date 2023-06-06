package main

import (
	"github.com/anything-devs/bike-soft-rest-api/configs"
	"github.com/anything-devs/bike-soft-rest-api/models"
	"github.com/anything-devs/bike-soft-rest-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	DB := configs.ConectarBD()
	DB.Table("productos").AutoMigrate(models.Producto{})
	DB.Table("categorias").AutoMigrate(models.Categoria{})

	routes.Rutas(router, DB)

	// Configurar opciones CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Run(":8083")
}
