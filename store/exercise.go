package store

import (
	"context"

	"github.com/t-tazy/my_portfolio_api/entity"
)

// 全てのエクササイズを取得する
func (r *Repository) ListExercises(
	ctx context.Context, db Queryer,
) (entity.Exercises, error) {
	exercises := entity.Exercises{}
	sql := `SELECT
			id, title, description, created, modified
			FROM exercises;`
	if err := db.SelectContext(ctx, &exercises, sql); err != nil {
		return nil, err
	}
	return exercises, nil
}

// エクササイズを保存する
// 引数として受け取った*entity.Exercise型のIDフィールドを更新することで
// 呼び出し元にRDBMSより発行されたIDを伝える
func (r *Repository) AddExercise(
	ctx context.Context, db Execer, e *entity.Exercise,
) error {
	e.Created = r.Clocker.Now()
	e.Modified = r.Clocker.Now()
	sql := `INSERT INTO exercises
			(title, description, created, modified)
			VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, e.Title, e.Description, e.Created, e.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = entity.ExerciseID(id)
	return nil
}
