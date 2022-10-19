package handlers

import (
	dto "BE-foodways/dto/result"
	usersdto "BE-foodways/dto/users"
	"BE-foodways/models"
	"BE-foodways/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type handler struct {
	UserRepository repositories.UserRepositories
}

func HandlerUser(UserRepository repositories.UserRepositories) *handler {
	return &handler{UserRepository}
}

  func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers() // menjalankan query kedatabase
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}
	

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-type", "application/json")

	 request := new(usersdto.CreateUserRequest)
	 if err := json.NewDecoder(r.Body).Decode(&request)
	 err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	 }

	 validation := validator.New()
	 err := validation.Struct(request)
	 if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	 }

	 user := models.User{
		Name: request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Location: request.Location,
		Image: request.Image,
		Role: request.Role,
	 }

	 data, err := h.UserRepository.CreateUser(user)
	 if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	 }

	 w.WriteHeader(http.StatusOK)
	 response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	 json.NewEncoder(w).Encode(response)


}
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Phone: u.Phone,
		Location: u.Location,
	}
}
  