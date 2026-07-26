package service

import (
	"errors"
	"knives/internal/models"
)

func validateKnife(k *models.Knife) error {
	if k.Name == "" {
		return errors.New("name is required")
	}
	if k.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}
