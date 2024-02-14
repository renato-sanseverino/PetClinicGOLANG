package domain

import (
	"petclinic/prisma/db"
	"petclinic/src/utils"
	"strings"

	"github.com/google/uuid"
)

type Pet struct {
	ID          uuid.UUID // Anotations para o GIN FRAMEWORK
	Name        string    `json:"name"`
	Breed       string    `json:"breed"`
	Age         int       `json:"age"`
	Owner       *int      `json:"owner"`
	FlagRemoved bool      `json:"flag_removed"`
}

func NewPet(name string, breed string, age int) (*Pet, error) {
	return &Pet{
		Name:  name,
		Breed: breed,
		Age:   age,
	}, nil
}

// Usar apenas para pequenas coleções, por questoes de performance e escalabilidade
// TODO: trazer do banco de dados os registros filtrados e paginados
func FilterByBreed(pets []db.PetModel, breed string) []db.PetModel {
	f := utils.Filter(pets, func(p db.PetModel) bool {
		return strings.Contains(strings.ToLower(p.Breed), strings.ToLower(breed))
	})
	return f
}
