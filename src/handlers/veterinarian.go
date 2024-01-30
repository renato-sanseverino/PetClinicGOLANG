package handlers

import (
	"net/http"
	"os"
	"petclinic/src/domain"
	"petclinic/src/utils"

	"github.com/gin-gonic/gin"
)

// TODO: implementar o use case de preescrição
func ReceitarTratamento(c *gin.Context) {

	pet, _ := domain.NewPet("Ollie", "Doberman", 3)
	veterinarian, _ := domain.NewVeterinarian("Doctor Who", "SP 9876543210")
	treatment, _ := domain.NewTreatment("antibiótico", pet, veterinarian)

	c.JSON(http.StatusOK, gin.H{"message": "", "Treatment": treatment})
}

func PostVeterinarian(c *gin.Context) {
	var databaseURL = os.Getenv("DATABASE_URL") // passar através do MIDDLEWARE   utils.appMiddleware()
	var pool = utils.GetPool(databaseURL)       // passar através do MIDDLEWARE   utils.appMiddleware()

	var payload domain.Veterinarian

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vet, err := pool.Exec(c, "INSERT INTO veterinarian (name, \"inscricaoCRMV\") VALUES ($1, $2)",
		// payload.ID,
		payload.Name,
		payload.InscricaoCRMV,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Veterinarian created successfully", "Veterinarian": vet})
}
