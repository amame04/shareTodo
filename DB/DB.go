package DB 

import (
  "database/sql"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

const app string = "mysql"

const id string = "root"
const password string = "password"
const ip string = "localhost:3306"
const database string = "todo"

const creds = id+":"+password+"@("+ip+")/"+database

var DB *sql.DB

func DBinit() {
  var err error
  DB, err = sql.Open(app, creds)
  if err != nil {
    log.Fatal(err)
  }
}