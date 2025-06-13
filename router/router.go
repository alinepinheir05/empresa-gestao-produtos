package router

import (
	"github.com/alinepinheir05/empresa-gestao-produtos/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func StartRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ph := handler.NewProductHandler(db)

	products := r.Group("/products")
	{
		products.POST("", ph.CreateProduct)
		products.GET("", ph.GetAllProducts)
		products.GET("/:id", ph.GetProductByID)
		products.PUT("/:id", ph.UpdateProduct)
		products.DELETE("/:id", ph.DeleteProduct)
	}

	return r
}
