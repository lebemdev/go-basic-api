package api

import (
	"go-basic-api/cmd/api/handlers"
	"go-basic-api/config"
	"go-basic-api/internal/domain"
	"go-basic-api/internal/routes"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupApp(config *config.Config) {
	router := gin.Default()

	db, err := gorm.Open(postgres.Open(config.DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error %v\n", err)

	}

	log.Println("dabase is connected")

	db.AutoMigrate(&domain.Product{})

	h := &handlers.Handler{
		ROUTER: router,
	}

	setupRoutes(h)

	router.Run(":8000")
}

func setupRoutes(h *handlers.Handler) {
	routes.SetupProductRoutes(h)
}
