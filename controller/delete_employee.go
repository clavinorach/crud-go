package controller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"net/http"

)

func NewDeleteEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

			_, err := db.Exec("DELETE FROM employee WHERE id = ? ", id) 
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
}