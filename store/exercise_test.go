package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/t-tazy/my_portfolio_api/clock"
	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/testutil"
)

// 実際のRDBMSを使ってテストする
func TestRepository_ListExercises(t *testing.T) {
	ctx := context.Background()
	// entity.Exerciseを作成する他のテストケースと混ざるとfail
	// そのため、トランザクションをはることでこのテストケースの中だけのテーブル状態にする
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareExercises(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListExercises(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(gots, wants); len(diff) != 0 {
		t.Errorf("differs: (-got +want)\n%s", diff)
	}
}

// exercisesテーブルの状態を整えるヘルパー
func prepareExercises(ctx context.Context, t *testing.T, con Execer) entity.Exercises {
	t.Helper()

	// 一度データをきれいにする
	if _, err := con.ExecContext(ctx, "DELETE FROM exercises;"); err != nil {
		t.Logf("failed to initialize exercises: %v", err)
	}

	c := clock.FixedClocker{}
	wants := entity.Exercises{
		{
			Title:       "exercise 1",
			Description: "want exercise 1",
			Created:     c.Now(),
			Modified:    c.Now(),
		},
		{
			Title:       "exercise 2",
			Description: "want exercise 2",
			Created:     c.Now(),
			Modified:    c.Now(),
		},
		{
			Title:       "exercise 3",
			Description: "want exercise 3",
			Created:     c.Now(),
			Modified:    c.Now(),
		},
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO exercises (title, description, created, modified)
		VALUES
			(?, ?, ?, ?),
			(?, ?, ?, ?),
			(?, ?, ?, ?);`,
		wants[0].Title, wants[0].Description, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Description, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Description, wants[2].Created, wants[2].Modified,
	)
	if err != nil {
		t.Fatal(err)
	}
	// 一つのINSERT文で複数のレコードを作成した場合のLastInsertIdメソッドの戻り値は
	// MySQLでは1つ目のレコードのIDになる
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	wants[0].ID = entity.ExerciseID(id)
	wants[1].ID = entity.ExerciseID(id + 1)
	wants[2].ID = entity.ExerciseID(id + 2)
	return wants
}

// mockを使ってテストする
func TestRepository_AddExercise(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int64 = 20
	okExercise := &entity.Exercise{
		Title:       "ok Exericse",
		Description: "test exercise",
		Created:     c.Now(),
		Modified:    c.Now(),
	}

	// mockデータベース接続と空のmockを生成
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	// mock化
	mock.ExpectExec(
		// エスケープが必要
		`INSERT INTO exercises \(title, description, created, modified\) VALUES \(\?, \?, \?, \?\)`,
	).WithArgs(okExercise.Title, okExercise.Description, okExercise.Created, okExercise.Modified).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.AddExercise(ctx, xdb, okExercise); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}
