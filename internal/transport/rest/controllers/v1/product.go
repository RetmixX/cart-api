package v1

import (
	"cart-api/internal/service"
	"cart-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	service service.Producter
	log     *log.Logger
}

func New(service service.Producter, log *log.Logger) *Product {
	return &Product{
		service: service,
		log:     log,
	}
}

func (p *Product) GetAll(c *gin.Context) {
	res, err := p.service.AllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
