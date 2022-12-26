package service

import (
	"context"
	"fmt"

	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/store"
)

type ListExercise struct {
	DB   store.Queryer
	Repo ExerciseLister
}

func (l *ListExercise) ListExercises(ctx context.Context) (entity.Exercises, error) {
	exercises, err := l.Repo.ListExercises(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return exercises, nil
}
