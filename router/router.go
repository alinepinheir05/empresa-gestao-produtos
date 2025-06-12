package router

import (
	"github.com/alinepinheir05/empresa-gestao-produtos/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	ProductHandler := handler.NewProductHandler(db)

	products := r.Group("/products")
	{
		products.POST("", ProductHandler.CreateProduct)
		products.GET("", ProductHandler.GetAllProducts)
		products.GET("/:id", ProductHandler.GetProductByID)
		products.DELETE("/:id", ProductHandler.DeleteProduct)
		products.PUT("/:id", ProductHandler.UpdateProduct)
	}

	return r
}
