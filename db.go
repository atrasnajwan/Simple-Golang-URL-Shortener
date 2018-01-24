package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

const DB_NAME = "simple_shortener" //database name
const DB_USER = "root"
const DB_PASSWORD = ""

func connectDb() (*sql.DB, error){
	db, err := sql.Open("mysql", DB_USER + ":" + DB_PASSWORD + "@/" + DB_NAME)
	return db, err
}

func insertDb(db *sql.DB, link ShortenLink) error{
	stmt, err := db.Prepare("INSERT shorten SET id=?, shorturl=?, longurl=?")
	defer stmt.Close()
	_, err = stmt.Exec(link.ID, link.ShortUrl, link.LongUrl)

	return err
}

func selectRow(db *sql.DB, idUrl string) (ShortenLink, error){
	stmt, err := db.Prepare("SELECT * FROM shorten WHERE id=?")
	defer stmt.Close()
	var link ShortenLink
	err = stmt.QueryRow(idUrl).Scan(&link.ID, &link.ShortUrl, &link.LongUrl)

	return link, err
}
