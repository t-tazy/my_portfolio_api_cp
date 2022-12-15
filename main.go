package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/t-tazy/my_portfolio_api/config"
)

// メインプロセスのrun関数を実行する
func main() {
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

// 環境変数から読み込んだポート番号でリッスンし、リクエストパスをレスポンスとして返すHTTPサーバーを起動
func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	s := &http.Server{
		// 引数で受け取ったnet.Listenerを利用するため、
		// Addrフィールドは指定しない
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	errCh := make(chan error)

	// 別ゴルーチンでHTTPサーバーを起動
	go func() {
		// Serveメソッドに変更
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
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
