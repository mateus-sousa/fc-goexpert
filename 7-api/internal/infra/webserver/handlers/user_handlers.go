package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/mateus-sousa/goexpert/7-api/internal/dto"
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/database"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}
type UserHandler struct {
	repository database.UserInterface
}

func NewUserHandler(repository database.UserInterface) *UserHandler {
	return &UserHandler{repository: repository}
}

// Create user godoc
//
//	@Summary      Create user
//	@Description  Create user
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.CreateUserInput  true  "user request"
//	@Success      201
//	@Failure      500  {object}  Error
//	@Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		httpError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(httpError)
		return
	}
	err = h.repository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		httpError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(httpError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetJWT user godoc
//
//	@Summary      Get a user JWT
//	@Description  Get a user JWT
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.GetJWTInput  true  "user credentials"
//	@Success      200 {object} dto.GetJWTOutput
//	@Failure      404 {object}  Error
//	@Failure      500  {object}  Error
//	@Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpireIn := r.Context().Value("jwtExpiresIn").(int)
	var jwtDto dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.repository.FindByEmail(jwtDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		httpError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(httpError)
		return
	}
	if !user.ValidatePassword(jwtDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		httpError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(httpError)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpireIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
