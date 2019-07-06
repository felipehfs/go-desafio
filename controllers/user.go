package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"

	"github.com/dgrijalva/jwt-go"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/felipehfs/godesafio/models"
	"github.com/felipehfs/godesafio/utils"
)

type UserHandler struct {
	DB *sql.DB
}

func (userHandler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	userDao := models.NewUserDao(userHandler.DB)
	user := models.User{}
	u, err := uuid.NewV4()
	checkError(w, err)
	user.UUID = u.String()
	err = json.NewDecoder(r.Body).Decode(&user)
	checkError(w, err)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	checkError(w, err)
	user.Password = string(hash)
	err = userDao.Register(user)

	checkError(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func (userHandler UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	userDao := models.NewUserDao(userHandler.DB)
	err := json.NewDecoder(r.Body).Decode(user)

	errorMessage := map[string]string{
		"status":  "error",
		"message": "Usuário não pode ser autenticado!",
	}

	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorMessage)
		log.Println("EncodingBodyError", err)
		return
	}

	find, err := userDao.FindOne(user.Email)

	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorMessage)
		log.Println("userDaoError", err)
		return
	}

	passwordInvalid := bcrypt.CompareHashAndPassword([]byte(find.Password), []byte(user.Password))
	if passwordInvalid != nil {
		w.WriteHeader(401)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorMessage)
		log.Println("userDaoError", err)
		return
	}

	duration := time.Now().Add(time.Hour * time.Duration(5) * 24 * 7)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["sub"] = find.Email
	claims["exp"] = duration.Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(utils.SECRET_KEY)

	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorMessage)
		log.Println("tokenError", err)
		return
	}

	body := models.SuccessMessage{
		Status:   "success",
		Message:  "Usuario encontrado  e token gerado",
		Tokenjwt: tokenString,
		Expires:  duration.Format("2006-01-02"),
		Tokenmsg: "use o token para acessar os endpoints!",
		Login: models.LoginInfo{
			ID:         find.ID,
			UUIDuser:   find.UUID,
			Avatarurl:  find.AvatarURL,
			Avatartype: "image/jpeg",
			Name:       find.Name,
			DataStart:  find.DataStart.Format("2006-01-02"),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (userHandler UserHandler) FindByUUID(w http.ResponseWriter, r *http.Request) {
	userDao := models.NewUserDao(userHandler.DB)
	vars := mux.Vars(r)
	search, err := userDao.FindByUUID(vars["uuid"])

	checkError(w, err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(search)
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Um erro ocorreu", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (userHandler UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	userDao := models.NewUserDao(userHandler.DB)
	var body struct {
		id int
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	checkError(w, err)

	userDao.RemoveUser(body.id)
	w.WriteHeader(http.StatusNoContent)
}

func (userHandler UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userDao := models.NewUserDao(userHandler.DB)
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	checkError(w, err)
	user, err := userDao.UpdateUser(*newUser)
	checkError(w, err)
	json.NewEncoder(w).Encode(user)
}
