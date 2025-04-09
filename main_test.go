package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomePage nos aseguramos que la ruta raiz devuelve el mensaje correcto y el status 200
func TestHomePage(t *testing.T) {

	router := setupRouter()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedMsg := "Bienvenidos a SysAcad!"
	if msg, ok := response["message"]; !ok || msg != expectedMsg {
		t.Errorf("Expected message '%s', got '%v'", expectedMsg, msg)
	}
}
