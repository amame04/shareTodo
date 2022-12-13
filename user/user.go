package user

import (
  "log"
  "golang.org/x/crypto/bcrypt"
  "database/sql"

  "sampleapp/DB"

  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id string `json:"id"`
  Pass string `json:"pass"`
}

type Todo struct {
  TodoId int `json:"todoId"`
  TodoContent string `json:"todoContent"`
  TodoDate string `json:"todoDate"`
  CreatedUser string `json:"createdUser"`
  DeleteFlag bool `json:"deleteFlag"`
}

var db *sql.DB

func Init() {
  db = DB.DB
}

func GetTodoList(user string) []Todo {
  sql, err := db.Prepare("SELECT * FROM todo where user = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer sql.Close()

  row, err := sql.Query(user)
  if err != nil {
    log.Fatal(err)
  }

  var todoList []Todo
  for row.Next() {
    var todo Todo
    row.Scan(&todo.TodoId, 
             &todo.TodoContent,
             &todo.TodoDate,
             &todo.CreatedUser,
             &todo.DeleteFlag)
    todoList = append(todoList, todo)
  }

  return todoList
}

func GetUserList() []User {
  sql, err := db.Prepare("SELECT * FROM users")
  if err != nil {
    log.Fatal(err)
  }
  defer sql.Close()

  row, err := sql.Query()
  if err != nil {
    log.Fatal(err)
  }

  var userList []User
  for row.Next() {
    var user User
    err := row.Scan(&user.Id, &user.Pass)
    user.Pass = "" 
    if err != nil {
      log.Fatal(err)
    }

    userList = append(userList, user)
  }

  return userList
}

func RegisterUser(id, pass string) bool {
  hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
  if err != nil {
    log.Fatal(err)
    return false
  }

  sql, err := db.Prepare("INSERT INTO users VALUES(?, ?)")
  if err != nil {
    log.Fatal(err)
    return false
  }

  _, err = sql.Exec(id, hash)
  
  return err == nil 
}

func LoginUser(id, pass string) bool {
  sql, err := db.Prepare("SELECT password FROM users where id = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer sql.Close()

  row, err := sql.Query(id)
  if err != nil {
    log.Fatal(err)
  }

  if !row.Next() {
    return false;
  }

  var dbPass string
  err = row.Scan(&dbPass)

  if err != nil {
    log.Fatal(err)
    return false;
  }

  err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass))
  return err == nil
}