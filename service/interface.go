package service

import (
	"context"

	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/store"
)

// storeパッケージの直接参照を避ける

//go:generate go run github.com/matryer/moq -out moq_test.go . ExerciseAdder ExerciseLister
type ExerciseAdder interface {
	AddExercise(ctx context.Context, db store.Execer, e *entity.Exercise) error
}

type ExerciseLister interface {
	ListExercises(ctx context.Context, db store.Queryer) (entity.Exercises, error)
}
