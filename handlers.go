package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	jerr := json.NewEncoder(w).Encode(&info)
	if jerr != nil {
		serverErr(w, jerr)
		return
	}

	fmt.Fprint(w)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	j := json.NewEncoder(w)
	var jerr error

	users, err := getUsersDB(r.Context())
	if err != nil {
		jerr = j.Encode(&err)
	} else {
		jerr = j.Encode(&users)
	}

	if jerr != nil {
		serverErr(w, jerr)
		return
	}

	fmt.Fprint(w)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	j := json.NewEncoder(w)
	var jerr error

	id := r.URL.Query().Get("id")
	comments, err := getCommentsDB(r.Context(), id)
	if err != nil {
		jerr = j.Encode(&err)
	} else {
		jerr = j.Encode(&comments)
	}

	if jerr != nil {
		serverErr(w, jerr)
		return
	}

	fmt.Fprint(w)
}

func serverErr(w http.ResponseWriter, err error) {
	ErrLog.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
