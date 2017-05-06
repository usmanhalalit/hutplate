package hutplate

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)

func TestNewHttp(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	got := NewHttp(w, r)

	if got.Request != r {
		t.Errorf("NewHttp(w, r): Invalid request got %v, expected %v", got.Request, r)
	}
}

func TestSession(t *testing.T) {
	var r *http.Request
	var w *httptest.ResponseRecorder
	Boot()
	w = NewRecorder()
	r, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	hut := NewHttp(w, r)
	expected := "test_value"
	err = hut.Session.Set("test_key", expected)
	if err != nil {
		t.Fatal("failed to create session store", err)
	}

	got, err := hut.Session.Get("test_key")
	if err != nil {
		t.Fatal("failed to get from the session", err)
	}

	if got != expected {
		t.Errorf("Session Test: Invalid session value got %v, expected %v", got, expected)
	}
}

func TestSessionFlash(t *testing.T) {
	var r *http.Request
	var w *httptest.ResponseRecorder
	Boot()
	w = NewRecorder()
	r, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request %v", err)
	}
	hut := NewHttp(w, r)
	expected := "test_value"
	err = hut.Session.SetFlash("test_key", expected)
	if err != nil {
		t.Fatalf("failed to set flash %v", err)
	}

	got := hut.Session.GetFlash("test_key")

	if got != expected {
		t.Errorf("Session Test: Invalid session value got %v, expected %v", got, expected)
	}

}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *httptest.ResponseRecorder {
	return &httptest.ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
	}
}