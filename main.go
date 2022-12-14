package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// メインプロセスのrun関数を実行する
// コマンドライン引数でサーバーのポート番号を指定する
func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1] // ポート番号
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

// 動的に選択されたポート番号でリッスンし、リクエストパスをレスポンスとして返すHTTPサーバーを起動
func run(ctx context.Context, l net.Listener) error {
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
