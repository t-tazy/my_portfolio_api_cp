package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/t-tazy/my_portfolio_api/handler"
	"github.com/t-tazy/my_portfolio_api/store"
)

// マルチプレクサを返す(ルーティング情報を持つ)
func NewMux() http.Handler {
	mux := chi.NewRouter()
	// HTTPサーバーが稼働中か確認するための/healthエンドポイントを宣言
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するため明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	ae := &handler.AddExercise{Store: store.Exercises, Validator: v}
	mux.Post("/exercises", ae.ServeHTTP)

	le := &handler.ListExercise{Store: store.Exercises}
	mux.Get("/exercises", le.ServeHTTP)
	return mux
}
