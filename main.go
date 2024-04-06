package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DbPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMYSQLStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewAPIServer(":3000", store)
	api.Serve()
}
