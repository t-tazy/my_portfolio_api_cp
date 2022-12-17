package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// JSONを解析し、比較検証するテスト用ヘルパー
func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	// 任意の構造のJSONをデコードするため、デコード先はany型とする
	var jsonWant, jsonGot any
	if err := json.Unmarshal(want, &jsonWant); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jsonGot); err != nil {
		t.Fatalf("cannot unmarshal got %q: %v", got, err)
	}
	if diff := cmp.Diff(jsonGot, jsonWant); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

// レスポンスと期待値を比較検証するテスト用ヘルパー
func AssertResponse(t *testing.T, got *http.Response, status int, want []byte) {
	t.Helper()

	t.Cleanup(func() { _ = got.Body.Close() })

	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}
	if len(gb) == 0 && len(want) == 0 {
		// 期待としても実体としてもレスポンスボディがないため
		// AssertJSONを呼ぶ必要がない
		return
	}
	AssertJSON(t, want, gb)
}

// ゴールデンテスト用
// 入力値・期待値をファイルから取得するテスト用ヘルパー
func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %v", path, err)
	}
	return bt
}
