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

var pets []db.InnerPet

type Pet struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Owner *Owner `json:"owner"`
}

type Owner struct {
	name  string
	email string
}

func getPets(c *gin.Context) {
	client := GetPrisma(c)

	pets, err := client.Pet.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func postPet(c *gin.Context) {
	var payload Pet

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := GetPrisma(c)
	insertedPet, err := client.Pet.CreateOne(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.Age.Set(payload.Age),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming successful creation, return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Pet created successfully", "pet": insertedPet})
}

func patchPet(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the Pet struct for updates
	var payload Pet
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		client := GetPrisma(c)
		updatedPet, err := client.Pet.UpsertOne(
			payload,
		).Exec(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	*/

	c.JSON(http.StatusOK, gin.H{"message": "Pet " + strconv.Itoa(id) + " updated successfully"})
}

func deletePet(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	/*
		client := GetPrisma(c)
		client.Pet.Delete(id).Exec(c)
	*/

	c.JSON(http.StatusOK, gin.H{"message": "Pet " + strconv.Itoa(id) + " deleted successfully"})
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
