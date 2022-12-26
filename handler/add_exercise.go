package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/t-tazy/my_portfolio_api/entity"
)

type AddExercise struct {
	Service   AddExerciseService
	Validator *validator.Validate
}

// リクエストボディを読み込み、バリデーション
// 結果が正常ならエクササイズを作成し、保存する
// レスポンスボディにRDBMSにより発行されたIDを書き込む
func (ae *AddExercise) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	// バリデーション
	if err := ae.Validator.Struct(body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	e, err := ae.Service.AddExercise(ctx, body.Title, body.Description)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	// レスポンスとしてRDBMSにより発行されたIDを返す
	rsp := struct {
		ID entity.ExerciseID `json:"id"`
	}{ID: e.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
