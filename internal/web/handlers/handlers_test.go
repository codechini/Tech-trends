package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
// Not written by me nned to verify.
func TestFragment(t *testing.T) {
	rec := httptest.NewRecorder()
	Fragment(rec, httptest.NewRequest("GET", "/fragment", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}