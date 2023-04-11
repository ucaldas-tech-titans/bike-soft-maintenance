package routes

import (
	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func Rutas(router *gin.Engine) {
	rutasInicio(router)
	rutasProductos(router)
}

func rutasInicio(router *gin.Engine) {
	router.GET("/", controllers.ControladorPaginaInicio)
}

/*
* Método que contiene las rutas que se utilizan con productos
 */
func rutasProductos(router *gin.Engine) {
	router.GET("/productos", controllers.GetProductos)
	router.GET("/productos/filtradosBajasUnidades/:cantidad", controllers.FiltroBajasUnidades)
	router.GET("/producto", controllers.GetProducto)
}
