package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// メインプロセスのrun関数を実行する
func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

// :18080でリッスンし、リクエストパスをレスポンスとして返すHTTPサーバーを起動
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	errCh := make(chan error)

	// 別ゴルーチンでHTTPサーバーを起動
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			errCh <- err
		}
		errCh <- nil
	}()

	// コンテキストを通じて処理の中断を検知したとき
	// ShutdownメソッドでHTTPサーバーの機能を終了する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return <-errCh
}
