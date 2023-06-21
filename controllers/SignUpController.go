package signUp

import (
	"GoRestAPI/models"
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
	if err != nil {
		fmt.Println("mysql open error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var reqBody models.LoginParams
	body, error := io.ReadAll(r.Body)
	if error != nil {
		fmt.Println("myspl read all error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &reqBody); err != nil {
		fmt.Println("json decode error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//sqlでmysql操作
	//responseを返却する。
	_, queryErr := db.Query("SELECT * FROM User WHERE email = ? ", reqBody.Email)

	if !errors.Is(queryErr, sql.ErrNoRows) && queryErr != nil {
		fmt.Println("select user error:", queryErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodedPassword, encodeError := reqBody.EncodePassword()
	if encodeError != nil {
		fmt.Println("password encode error:", encodeError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, insertError := db.Exec(
		"INSERT User (email, password) VALUES (?, ?)",
		reqBody.Email,
		encodedPassword,
	)

	if insertError != nil {
		fmt.Println("insert user error", insertError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//201をreturn
	w.WriteHeader(http.StatusCreated)
	defer db.Close()
}
