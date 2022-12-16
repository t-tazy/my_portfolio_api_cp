package store

import (
	"errors"

	"github.com/t-tazy/my_portfolio_api/entity"
)

var (
	Exercises = &ExerciseStore{Exercises: map[entity.ExerciseID]*entity.Exercise{}}

	ErrNotFound = errors.New("not found")
)

// Exercise.IDフィールドはRDBMSによって、割り当てられることを想定している
// LastIDフィールドをその代用
type ExerciseStore struct {
	// 動作確認用の仮実装のため、あえてexport
	LastID    entity.ExerciseID
	Exercises map[entity.ExerciseID]*entity.Exercise
}

func (es *ExerciseStore) Add(e *entity.Exercise) (entity.ExerciseID, error) {
	es.LastID++
	e.ID = es.LastID
	es.Exercises[e.ID] = e
	return e.ID, nil
}

// ソート済みのエクササイズ一覧を返す
func (es *ExerciseStore) All() entity.Exercises {
	exercises := make([]*entity.Exercise, len(es.Exercises))
	for i, e := range es.Exercises {
		exercises[i-1] = e
	}
	return exercises
}
