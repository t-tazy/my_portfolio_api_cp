package main

import "net/http"

// マルチプレクサを返す(ルーティング情報を持つ)
func NewMux() http.Handler {
	mux := http.NewServeMux()
	// HTTPサーバーが稼働中か確認するための/healthエンドポイントを宣言
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するため明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
