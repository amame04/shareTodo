package DB

import (
  "database/sql"
  "log"
  "os"

  _"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBinit() {
  app := os.Getenv("DB_APP")
  user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  host := os.Getenv("DB_HOST")
  database := os.Getenv("DB_DB")

  creds := user+":"+password+"@("+host+")/"+database

  var err error
  DB, err = sql.Open(app, creds)
  if err != nil {
    log.Fatal(err)
  }
}
