package signUp

import (
	"GoRestAPI/ja"
	"GoRestAPI/models"
	"GoRestAPI/utils"

	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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

	validateError := reqBody.Validates()
	if validateError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(validateError.Error())
		w.Write([]byte(errorMessage))
		return
	}

	_, queryErr := db.Query("SELECT * FROM User WHERE email = ? ", reqBody.Email)

	if !errors.Is(queryErr, sql.ErrNoRows) && queryErr != nil {
		fmt.Println(queryErr)
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.CreateUserErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}

	encodedPassword, encodeError := reqBody.EncodePassword()
	if encodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.EncodeErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}

	_, insertError := db.Exec(
		"INSERT User (email, password) VALUES (?, ?)",
		reqBody.Email,
		encodedPassword,
	)

	if insertError != nil {
		fmt.Println(insertError)
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.CreateUserErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}
	token, err := Utils.EncodeJwtToken(reqBody.Email)
	response := models.SessionResponse{
		token,
		ja.SuccessSignUpMessage,
	}
	bytes, err := json.Marshal(response)
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
