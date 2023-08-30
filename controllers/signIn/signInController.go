package signIn

import (
	"GoRestAPI/ja"
	"GoRestAPI/models"
	Utils "GoRestAPI/utils"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	dbconf := ""
	db, err := sql.Open("mysql", dbconf)
	defer db.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.ServerErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}

	var reqBody models.LoginParams
	body, error := io.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.ServerErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}

	if err := json.Unmarshal(body, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.EncodeErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}

	result := db.QueryRow("SELECT password FROM User WHERE email = ? ", reqBody.Email)

	var passwordDigest string
	if scanError := result.Scan(&passwordDigest); scanError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.PasswordMissingMessage)
		w.Write([]byte(errorMessage))
		return
	}

	missingPasswordError := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(reqBody.Password))
	if missingPasswordError != nil {
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := Utils.ErrorMessage(ja.PasswordMissingMessage)
		w.Write([]byte(errorMessage))
		return
	}

	token, err := Utils.EncodeJwtToken(reqBody.Email)
	response := models.SessionResponse{
		token,
		ja.SuccessSignInMessage,
	}
	bytes, err := json.Marshal(response)
	w.Write(bytes)
}
