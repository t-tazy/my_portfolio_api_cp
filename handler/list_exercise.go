package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/store"
)

type ListExercise struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

// テスト用にCreated以外のフィールドを持つようにする
type exercise struct {
	ID          entity.ExerciseID `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
}

// DBから一覧を取得し、レスポンスボディに書き込む
func (le *ListExercise) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exercises, err := le.Repo.ListExercises(ctx, le.DB)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []exercise{}
	for _, e := range exercises {
		rsp = append(rsp, exercise{
			ID:          e.ID,
			Title:       e.Title,
			Description: e.Description,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
