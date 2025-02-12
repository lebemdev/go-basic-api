package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func (h ProductHandler) CreateProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello %v", ctx.Param("name"))

}

func (h ProductHandler) GetProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Products")
}
