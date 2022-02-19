package main

import (
	"log"
	"net/http"
	"os"
)

//Main info
var info = map[string]string{
	"Title":   "Useless API",
	"Version": "1.0",
	"Author":  "PoorMouse",
	"URL":     "https://github.com/PoorMouse/useless-api",
}

//Logging
var (
	InfoErr *log.Logger
	ErrLog  *log.Logger
)

func main() {
	InfoErr = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	err := DBConnect()
	if err != nil {
		ErrLog.Fatal(err)
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/getUsers", getUsers)
	mux.HandleFunc("/getComments", getComments)

	port := os.Getenv("PORT")

	InfoErr.Println("Server has been started")
	ErrLog.Fatal(http.ListenAndServe(":"+port, mux))
}
