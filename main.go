package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"math/rand"
	"io/ioutil"
	"database/sql"
)

const LENGTH_CODE = 8
const PORT  = ":8000"

type ShortenLink struct {
	ID string `json:"id"`
	ShortUrl string `json:"shorturl"`
	LongUrl string `json:"longurl"`
}

var pattern = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var db *sql.DB

func CreateShorten(w http.ResponseWriter, r *http.Request) {
	var l ShortenLink

	body, err := ioutil.ReadAll(r.Body)
	//parse json to struct
	err = json.Unmarshal(body, &l);
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	l.ID = getRandCode(LENGTH_CODE)
	l.ShortUrl = r.Host + "/" + l.ID

	//insert to db
	err = insertDb(db, l)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(l)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//get link from db
	link, err := selectRow(db, params["code"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		http.Redirect(w, r, link.LongUrl, http.StatusSeeOther)
	}
}

func getRandCode(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = pattern[rand.Intn(len(pattern))]
	}
	return string(b)
}

func main() {
	db1, err := connectDb()
	db = db1
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateShorten).Methods("POST")
	router.HandleFunc("/{code}", Redirect).Methods("GET")
	fmt.Printf("Server is starting")
	log.Fatal(http.ListenAndServe(PORT, router))
}