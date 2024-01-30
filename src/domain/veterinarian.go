package domain

import (
	"errors"
)

type Veterinarian struct { // Anotations para o GIN FRAMEWORK
	ID            int32  // uuid.UUID
	Name          string `json:"name" binding:"required"`
	InscricaoCRMV string `json:"inscricao_crmv" binding:"required"`
}

func NewVeterinarian(Name string, InscricaoCRMV string) (*Veterinarian, error) {
	obj := &Veterinarian{
		ID:            0, // uuid.New()
		Name:          Name,
		InscricaoCRMV: InscricaoCRMV,
	}

	if !obj.Validate() {
		return nil, errors.New("invalid data")
	}
	return obj, nil
}

func (obj *Veterinarian) Validate() bool {
	if obj.Name == "" || obj.InscricaoCRMV == "" {
		return false
	}
	return true
}
