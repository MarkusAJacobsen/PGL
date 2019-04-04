package internal

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"
)

type test struct {
	Foo  string `json:"foo"`
	Test string `json:"test"`
}

func TestGetPostData(t *testing.T) {
	t.Run("Successful getter", func(t *testing.T) {
		str := strings.NewReader(`{"foo": "bar", "test": "Hello world!"}`)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "localhost:3333", str)

		var testStructure test
		GetPostData(req.Body, &testStructure, w)

		if testStructure.Foo != "bar" {
			t.Error("Received wrong data")
		}

		if testStructure.Test != "Hello world!" {
			t.Error("Received wrong data")
		}
	})

	t.Run("Supplied io.Reader is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "localhost:3333", nil)

		var testStructure test
		GetPostData(req.Body, &testStructure, w)

		if w.Result().StatusCode != 400 {
			t.Errorf("Expected: 400. Received: %d", w.Result().StatusCode)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(w.Result().Body)

		if buf.String() != "Could not process request data\n" {
			t.Errorf("Expected: Could not process request data. Received: %s", buf.String())
		}
	})
}
