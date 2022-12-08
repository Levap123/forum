package main

import repository "forum/internal/repository/sqlite3"

func main() {
	db, err := repository.NewDb()
	if err!=nil{
		return //TODO: IMPLEMENT
	}
}
