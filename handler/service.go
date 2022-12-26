package handler

import (
	"context"

	"github.com/t-tazy/my_portfolio_api/entity"
)

// リクエストの解釈とレスポンスの構築以外を次のインターフェースに移譲

//go:generate go run github.com/matryer/moq -out moq_test.go . ListExercisesService AddExerciseService
type ListExercisesService interface {
	ListExercises(ctx context.Context) (entity.Exercises, error)
}

type AddExerciseService interface {
	AddExercise(ctx context.Context, title, description string) (*entity.Exercise, error)
}
