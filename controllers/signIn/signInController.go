package signIn

import (
	"GoRestAPI/ja"
	"GoRestAPI/utils"
	"database/sql"
	"net/http"
)



func SignInHandler(w http.ResponseWriter, r *http.Request) {
	dbconf := ""
	db, err := sql.Open("mysql", dbconf)

	if err != nil {		
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Utils.ErrorMessage(ja.ServerErrorMessage)
		w.Write([]byte(errorMessage))
		return
	}
}