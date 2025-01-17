package token

import (
	"context"
)

type UseCase interface {
	GenerateTokenPair(ctx context.Context, payload string) (*Pair, error)
	VerifyToken(ctx context.Context, tokenString string) (*Claims, error)
	RevokeToken(ctx context.Context, token string) (*string, error)
}

type UseCaseImpl struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &UseCaseImpl{repository: repository}
}

func (u UseCaseImpl) GenerateTokenPair(ctx context.Context, payload string) (*Pair, error) {
	return u.repository.GenerateTokenPair(ctx, payload)
}

func (u UseCaseImpl) VerifyToken(ctx context.Context, token string) (*Claims, error) {
	return u.repository.VerifyToken(ctx, token)
}

func (u UseCaseImpl) RevokeToken(ctx context.Context, token string) (*string, error) {
	claims, err := u.repository.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return u.repository.RevokeToken(ctx, token, claims.ExpirationTime)
}
