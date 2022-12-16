package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

// http.Server型をラップした型を返す
// 動的に選択したポートをリッスンするためにnet.Listener型の値を引数で受け取る
// ルーティングの設定も引数で受け取る
func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

// HTTPサーバーを起動する
func (s *Server) Run(ctx context.Context) error {
	// シグナルをハンドリング
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error)

	// 別ゴルーチンでHTTPサーバーを起動
	go func() {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			errCh <- err
		}
		errCh <- nil
	}()

	// コンテキストを通じて処理の中断を検知したとき
	// ShutdownメソッドでHTTPサーバーの機能を終了する
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return <-errCh
}
