package handlers

import (
	"net/http"
	"os"
	"petclinic/src/domain"
	"petclinic/src/utils"

	"github.com/gin-gonic/gin"
)

func PostVeterinarian(c *gin.Context) {
	var databaseURL = os.Getenv("DATABASE_URL") // passar através do MIDDLEWARE   utils.appMiddleware()
	var pool = utils.GetPool(databaseURL)       // passar através do MIDDLEWARE   utils.appMiddleware()

	var payload domain.Veterinarian

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vet, err := pool.Exec(c, "INSERT INTO veterinarian (id, name, \"InscricaoCRMV\") VALUES ($1, $2, $3)",
		payload.ID,
		payload.Name,
		payload.InscricaoCRMV,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Veterinarian created successfully", "Veterinarian": vet})
}
