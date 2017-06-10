package hutplate

import (
	"net/http"
	"testing"
)

func TestResponse_Redirect(t *testing.T) {
	w := NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	hut := NewHttp(w, r)

	expectedLocation := "http://example.com/a"
	hut.Response.Redirect(expectedLocation)
	gotLocation := w.Header().Get("Location")
	if gotLocation != expectedLocation {
		t.Errorf("invalid redirect location expected %v got %v", expectedLocation, gotLocation)
	}

	expectedCode := 302
	gotCode := w.Code
	if gotCode != expectedCode {
		t.Errorf("invalid redirect code expected %v got %v", expectedCode, gotCode)
	}
}

func TestResponse_RedirectWithCode(t *testing.T) {
	w := NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	hut := NewHttp(w, r)

	hut.Response.Redirect("a", 301)

	expectedCode := 301
	gotCode := w.Code
	if gotCode != expectedCode {
		t.Errorf("invalid redirect code expected %v got %v", expectedCode, gotCode)
	}
}

func TestResponse_With(t *testing.T) {
	w := NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	hut := NewHttp(w, r)

	expected := "Test error message"
	hut.Response.With("error", expected).Redirect("")
	got := hut.Session.GetFlash("error")

	if got != expected {
		t.Errorf("invalid redirect code expected %v got %v", expected, got)
	}
}
