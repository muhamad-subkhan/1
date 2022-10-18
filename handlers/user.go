package handlers

import (
	"BE-foodways/repositories"
	"net/http"
)

type handler struct {
	UserRepositories repositories.UserRepositories
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {

  }
  