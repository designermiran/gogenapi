package server

import (
	"github.com/designermiran/gogenapi/_example/middleware"
	"github.com/designermiran/gogenapi/_example/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
