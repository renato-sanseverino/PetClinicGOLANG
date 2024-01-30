package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Treatment struct {
	ID           uuid.UUID     // Anotations para o GIN FRAMEWORK
	Description  string        `json:"description" binding:"required"`
	Pet          *int          `json:"pet" binding:"required"`
	Veterinarian *Veterinarian `json:"veterinarian" binding:"required"`
}

func NewTreatment(description string, pet *int, veterinarian *Veterinarian) (*Treatment, error) {
	obj := &Treatment{
		ID:           uuid.New(),
		Description:  description,
		Pet:          nil,
		Veterinarian: nil,
	}

	if !obj.Validate() {
		return nil, errors.New("invalid data")
	}
	return obj, nil
}

func (obj *Treatment) Validate() bool {
	if obj.Description == "" || obj.Veterinarian == nil {
		return false
	}
	return true
}
