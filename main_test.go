package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// 別ゴルーチンでテスト対象のrun関数を実行しHTTPサーバーを起動
// エンドポイントに対してGETリクエストを送信し、レスポンスを検証する
func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)
	go func() {
		errCh <- run(ctx)
	}()
	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// HTTPサーバーの戻り値を検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数に終了通知を送信する
	cancel()
	// run関数の戻り値を検証する
	if err := <-errCh; err != nil {
		t.Fatal(err)
	}
}
