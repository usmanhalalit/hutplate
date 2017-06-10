package hutplate

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHttp(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	got := NewHttp(w, r)

	if got.Request != r {
		t.Errorf("NewHttp(w, r): Invalid request got %v, expected %v", got.Request, r)
	}
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *httptest.ResponseRecorder {
	return &httptest.ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
	}
}
