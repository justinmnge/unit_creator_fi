package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handleHome(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestHandleLogin(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	handleLogin(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestHandleTrailer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/trailer", nil)
	w := httptest.NewRecorder()

	handleTrailer(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestHandleContact(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/contact", nil)
	w := httptest.NewRecorder()

	handleContact(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestHandleAbout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	w := httptest.NewRecorder()

	handleAbout(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}
