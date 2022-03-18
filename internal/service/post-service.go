package service

import (
	"context"
	"someApp/domain"
)

type service struct {
	repo domain.UserRepo
}

type PostService interface {
	Create(ctx context.Context, user *domain.User) error
	Find(ctx context.Context, id string) (*[]domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

func NewPostService(repository domain.UserRepo) PostService {
	return &service{repo: repository}
}

func (s *service) Create(ctx context.Context, user *domain.User) error {
	return s.repo.InsertData(ctx, user)
}

func (s *service) Find(ctx context.Context, id string) (*[]domain.User, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Update(ctx context.Context, user *domain.User) error {
	return s.repo.Update(ctx, user)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
