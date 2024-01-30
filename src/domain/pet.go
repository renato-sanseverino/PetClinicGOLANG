package domain

import (
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
