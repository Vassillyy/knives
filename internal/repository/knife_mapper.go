package repository

import "knifes/internal/models"

func knifeScanFields(k *models.Knife) []any {
	return []any{
		&k.ID, &k.Name, &k.Description, &k.Price, &k.Material,
		&k.BladeLength, &k.Handle, &k.Brand, &k.CreatedAt, &k.UpdatedAt, &k.DeletedAt,
	}
}

func knifeBaseArgs(k *models.Knife) []any {
	return []any{k.Name, k.Description, k.Price, k.Material, k.BladeLength, k.Handle, k.Brand}
}

func knifeCreateArgs(k *models.Knife) []any {
	return knifeBaseArgs(k)
}

func knifeUpdateArgs(k *models.Knife) []any {
	return append(knifeBaseArgs(k), k.ID)
}
