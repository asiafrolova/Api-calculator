package application

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asiafrolova/Api-calculator/internal/application"
)

type Response struct {
	Expression string
	Result     float64
	Status     int
}
type Answer struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

func TestRequestHandlerBadRequestCase(t *testing.T) {
	testCasesSuccess := []Response{
		{Expression: "(2+2)*2",
			Result: 8,
			Status: 200},
		{Expression: "1/2",
			Result: 0.5,
			Status: 200},
		{Expression: "2+2*2",
			Result: 6,
			Status: 200},
	}
	for _, r := range testCasesSuccess {
		mcPostBody := map[string]interface{}{
			"expression": r.Expression,
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))

		w := httptest.NewRecorder()

		application.CalcHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != r.Status {
			t.Errorf("Bad status code, have: %d, want: %d", res.StatusCode, r.Status)
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error read body")
		}
		var ans Answer
		err = json.Unmarshal(data, &ans)

		i := ans.Result
		if err != nil {
			t.Errorf("Result is not number")
		}

		if i != r.Result {

			t.Errorf("Bad result have, %f, want: %f", i, r.Result)
		}
	}
	testCasesFail := []Response{
		{Expression: "1/0",
			Result: 0,
			Status: 400},
		{Expression: "(2+1)(",
			Result: 0,
			Status: 400},
		{Expression: "1++",
			Result: 0,
			Status: 400},
	}
	for _, r := range testCasesFail {
		mcPostBody := map[string]interface{}{
			"expression": r.Expression,
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))

		w := httptest.NewRecorder()

		application.CalcHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != r.Status {
			t.Errorf("Bad status code, have: %d, want: %d", res.StatusCode, r.Status)
		}

	}
}
