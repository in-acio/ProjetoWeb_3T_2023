package handlers

import (
	"backend/internal/database"
	"backend/internal/dto"
	"backend/internal/entity"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
	Jwt *jwtauth.JWTAuth
	JwtExpiresIn int
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func handleBadRequest(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(ErrorMessage{
        Error: message,
    })
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
		Jwt: jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	_, err = h.UserDB.FindByEmail(u.Email)
	if err == nil {
		handleBadRequest(w, "email já registrado")
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var loginData dto.LoginInput
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		handleBadRequest(w, "corpo da requisição inválido")
		return
	}

	u, err := h.UserDB.FindByEmail(loginData.Email)
	if err != nil {
		handleBadRequest(w, "usuário não encontrado")
		return
	}

	if !u.ValidatePassword(loginData.Password) {
		handleBadRequest(w, err.Error())
		return
	}

	_, tokenString, err := h.Jwt.Encode(map[string]interface{} {
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
		"userId": u.ID,
		"isAdmin": u.IsAdmin,
	})
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	accessToken := struct {
		AccessToken string `json:"token"`
		Email string `json:"email"`
		Name string `json:"name"`
		IsAdmin bool `json:"is_admin"`
	}{
		AccessToken: tokenString,
		Email: u.Email,
		Name: u.Name,
		IsAdmin: u.IsAdmin,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var userData dto.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		handleBadRequest(w, err.Error())
		return 
	}

	userId := claims["userId"].(float64)

	user, err := h.UserDB.FindById(uint(userId))
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	if !user.ValidatePassword(userData.OldPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newUser, err := entity.NewUser(userData.Name, user.Email, userData.NewPassword)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	newUser.ID = user.ID;

	err = h.UserDB.Update(newUser)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request){
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		handleBadRequest(w, err.Error())
		return 
	}
	userId := uint(claims["userId"].(float64))

	u, err := h.UserDB.FindById(userId)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}