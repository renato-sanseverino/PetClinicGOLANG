package main

import (
	"net/http"
	"petClinicAPI/prisma/db"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return client
}

func getPets(c *gin.Context) {
	// var pets []db.InnerPet

	client := GetPrisma(c)

	pets, err := client.Pet.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func postPet(c *gin.Context) {
	var payload db.InnerPet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := GetPrisma(c)
	insertedPet, err := client.Pet.CreateOne(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pet created successfully", "pet": insertedPet})
}

func patchPet(c *gin.Context) {
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

	client := GetPrisma(c)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pet patched", "pet": updatedPet})
}

func deletePet(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := GetPrisma(c)
	deletedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pet deleted successfully", "pet id": deletedPet.ID})
}

func main() {
	// godotenv.Load(".env")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := r.Group("/api")
	rGroup.GET("/pets", getPets)
	rGroup.POST("/pets", postPet)
	rGroup.PATCH("/pets/:id", patchPet)
	rGroup.DELETE("/pets/:id", deletePet)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
