package handlers;

import (
	"net/http"
	"strconv"
	"petclinic/prisma/db"
	"petclinic/src/utils"
	"github.com/gin-gonic/gin"
)

func GetPets(c *gin.Context) {
	// var pets []db.InnerPet

	client := utils.GetPrisma(c)

	pets, err := client.Pet.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func PostPet(c *gin.Context) {
	var payload db.InnerPet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedPet, err := client.Pet.CreateOne(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet created successfully", "pet": insertedPet})
}

func PatchPet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the Pet struct
	var payload db.InnerPet
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	updatedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Update(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet patched", "pet": updatedPet})
}

func DeletePet(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully", "pet id": deletedPet.ID})
}
