package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestMain(t *testing.T) {
	conectarDB = func() *gorm.DB {
		var DB *gorm.DB
		return DB
	}

	migrateTables = func(DB *gorm.DB) {}

	run = func(router *gin.Engine, DB *gorm.DB) {}

	main()
}
