//package awesomeProject
package main

import (
	"database/sql"
	"fmt"
)

type Book struct{
	Name string
	Year string
	Length string
}
const (
	DB_USER = "postgres"
	DB_PASSWORD = "I586qwer"
	DB_NAME = "lab"
)
func dbConnect() error {
	var err error
	//var connStr = "host=http://127.0.0.1/ port=14259 user=postgres password=I586qwer dbname=lab sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS confectionery (conf_type text,conf_price text,conf_date text)"); err != nil {
		return err
	}
	return nil
}
func dbAddBook(conf_type, conf_price, conf_date string) error {
	sqlstmt := "INSERT INTO confectionery VALUES ($1, $2, $3)"
	_, err := db.Exec(sqlstmt, conf_type, conf_price, conf_date)
	if err != nil {
		return err
	}
	return nil
}
func dbGetBooks() ([]Book, error) {
	var books []Book
	stmt, err := db.Prepare("SELECT conf_type, conf_price, conf_date FROM books")
	if err != nil {
		return books, err
	}
	res, err := stmt.Query()
	if err != nil {
		return books, err
	}
	var tempBook Book
	for res.Next() {
		err = res.Scan(&tempBook.Name, &tempBook.Year, &tempBook.Length)
		if err != nil {
			return books, err
		}
		books = append(books, tempBook)
	}
	return books, err
}
