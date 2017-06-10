package hutplate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	hut := GetARequest(t)
	expected := "test_value"
	err := hut.Session.Set("test_key", expected)
	if err != nil {
		t.Fatal("failed to create session store", err)
	}

	got, err := hut.Session.Get("test_key")
	if err != nil {
		t.Fatal("failed to get from the session", err)
	}

	if got != expected {
		t.Errorf("invalid session value got %v, expected %v", got, expected)
	}
}
func GetARequest(t *testing.T) Http {
	var r *http.Request
	var w *httptest.ResponseRecorder
	w = NewRecorder()
	r, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	hut := NewHttp(w, r)
	return hut
}

func TestSessionFlash(t *testing.T) {
	var r *http.Request
	var w *httptest.ResponseRecorder
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
		t.Errorf("invalid session value got %v, expected %v", got, expected)
	}

}
