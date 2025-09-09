package handler

import (
    "encoding/json"
    "net/http"
    "github.com/zinan-c/howGO/internal/service"
)

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // http request
}