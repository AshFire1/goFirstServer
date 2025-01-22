package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AshFire1/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Hi",
			LastName:  "World",
			Email:     "asd",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
	t.Run("should pass if the user payload is valid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Hi",
			LastName:  "World",
			Email:     "asd@gmail.com",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}

type mockUserStore struct {
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}
func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}
