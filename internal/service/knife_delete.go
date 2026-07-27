package service

import "context"

func (s *KnifeService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
