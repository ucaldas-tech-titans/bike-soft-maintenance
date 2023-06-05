package controllers

import (
	"net/http"

	"github.com/anything-devs/bike-soft-rest-api/repositories"
	"github.com/gin-gonic/gin"
)

// CategoriaController es el controlador de categorias
type CategoriaControllerImpl struct {
	categoriaRepository repositories.CategoriaRepository
}

// CategoriaController es la interfaz de CategoriaController
type CategoriaController interface {
	GetCategorias(ctx *gin.Context)
}

// NewCategoriaController crea una nueva instancia de CategoriaController.
func NewCategoriaController(categoriaRepository repositories.CategoriaRepository) CategoriaController {
	return &CategoriaControllerImpl{categoriaRepository: categoriaRepository}
}

// GetCategorias obtiene las categorías y las devuelve como respuesta JSON.
func (cc *CategoriaControllerImpl) GetCategorias(ctx *gin.Context) {
	categorias, err := cc.categoriaRepository.GetCategorias()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categorias)
}
