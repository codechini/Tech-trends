package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
// Not written by me nned to verify.
func TestPage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	Page(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// func TestFragment(t *testing.T) {
// 	rec := httptest.NewRecorder()
// 	Fragment(rec, httptest.NewRequest("GET", "/fragment", nil))
// 	if rec.Code != http.StatusOK {
// 		t.Fatalf("status = %d, want 200", rec.Code)
// 	}
// }