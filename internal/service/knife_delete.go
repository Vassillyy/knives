package service

import "context"

func (s *KnifeService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	photos, err := s.photoRepo.GetByKnifeID(ctx, id)
	if err != nil {
		return err
	}
	for _, photo := range photos {
		if err := s.deletePhoto(ctx, &photo); err != nil {
			return err
		}
	}

	return nil
}