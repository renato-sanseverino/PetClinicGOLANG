package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Veterinarian struct {
	ID            uuid.UUID // Anotations para o GIN FRAMEWORK
	Name          string    `json:"name" binding:"required"`
	InscricaoCRMV string    `json:"inscricao_crmv" binding:"required"`
}

func NewVeterinarian(Name string, InscricaoCRMV string) (*Veterinarian, error) {
	obj := &Veterinarian{
		ID:            uuid.New(),
		Name:          Name,
		InscricaoCRMV: InscricaoCRMV,
	}

	err := obj.Validate()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (obj *Veterinarian) Validate() error {
	if obj.Name == "" || obj.InscricaoCRMV == "" {
		return errors.New("invalid data")
	}
	return nil
}
