package handlers

import (
	authdto "BE-foodways/dto/auth"
	dto "BE-foodways/dto/result"
  usersdto "BE-foodways/dto/users"
	"BE-foodways/models"
	"BE-foodways/pkg/bcrypt"
	"BE-foodways/repositories"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepositories repositories.AuthRepositories
}

func HandlerAuth(AuthRepositories repositories.AuthRepositories) *handlerAuth {
	return &handlerAuth{AuthRepositories}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Role:     request.Role,
	}

	data, err := h.AuthRepositories.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseAuth(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseAuth(u models.User) usersdto.CreateUserRequest {
	return authdto.RegisterRequest{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Gender:   u.Gender,
		Phone:    u.Phone,
		Role:     u.Role,
	}
}
