package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/store"
	"github.com/t-tazy/my_portfolio_api/testutil"
)

func TestListExercise(t *testing.T) {
	t.Parallel()

	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		exercises map[entity.ExerciseID]*entity.Exercise
		want      want
	}{
		"ok": {
			exercises: map[entity.ExerciseID]*entity.Exercise{
				1: {
					ID:          1,
					Title:       "test1",
					Description: "",
				},
				2: {
					ID:          2,
					Title:       "test2",
					Description: "テスト用",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_exercise/ok_rsp.json.golden",
			},
		},
		"empty": {
			exercises: map[entity.ExerciseID]*entity.Exercise{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_exercise/empty_rsp.json.golden",
			},
		},
	}
	for key, test := range tests {
		test := test
		t.Run(key, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/exercises", nil)

			sut := ListExercise{
				Store: &store.ExerciseStore{Exercises: test.exercises},
			}
			sut.ServeHTTP(w, r)
			rsp := w.Result()
			testutil.AssertResponse(
				t, rsp, test.want.status, testutil.LoadFile(t, test.want.rspFile),
			)
		})
	}
}
