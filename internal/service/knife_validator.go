package service

import (
	"errors"
	"knives/internal/models"
)

func validateKnifeCreate(k *models.Knife) error {
	if k.Name == nil || *k.Name == "" {
		return errors.New("name is required")
	}
	if k.Price == nil || *k.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}

func validateKnifeUpdate(k *models.Knife) error {
	if k.Name != nil && *k.Name == "" {
		return errors.New("name cannot be empty")
	}
	if k.Price != nil && *k.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}
