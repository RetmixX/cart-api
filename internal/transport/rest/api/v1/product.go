package v1

import (
	v1 "cart-api/internal/transport/rest/controllers/v1"
	"github.com/gin-gonic/gin"
)

const productURL = "/v1/products"

func RegisterRoutes(engine *gin.Engine, productController *v1.Product) {
	groupProduct := engine.Group(productURL)

	groupProduct.GET("all", productController.GetAll)
}
