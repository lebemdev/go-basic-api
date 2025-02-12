package routes

import (
	"go-basic-api/cmd/api/handlers"
)

func SetupProductRoutes(h *handlers.Handler) {
	router := h.ROUTER

	handler := handlers.ProductHandler{}

	privateRoutes := router.Group("/api/v1/products")
	privateRoutes.POST("/add", handler.CreateProduct)

}
