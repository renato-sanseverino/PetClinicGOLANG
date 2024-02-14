package main

import (
	"net/http"
	"petclinic/src/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	godotenv.Load(".env")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	// TODO: verificar integração entre o GIN e o Prometheus
	// rGroup.GET("/metrics", promhttp.Handler())
	rGroup.GET("/pets", handlers.GetPets)
	rGroup.GET("/pets/findByBreed/:breed", handlers.FindByBreed)
	rGroup.POST("/pets", handlers.PostPet)
	rGroup.PATCH("/pets/:id", handlers.PatchPet)
	rGroup.DELETE("/pets/:id", handlers.DeletePet)
	rGroup.POST("/veterinarians", handlers.PostVeterinarian)
	rGroup.POST("/receitar_tratamento", handlers.ReceitarTratamento)

	router.Use(cors.Default())
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
