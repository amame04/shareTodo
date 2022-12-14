package user

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"sampleapp/DB"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id string `json:"id"`
  //Pass string `json:"pass"`
}

type ShareUser struct {
  User string `json:"user"`
  DoneFlag bool `json:"doneFlag"`
}

type Todo struct {
  TodoId int `json:"todoId"`
  TodoContent string `json:"todoContent"`
  TodoDate string `json:"todoDate"`
  CreatedUser string `json:"createdUser"`
  DoneFlag bool `json:"doneFlag"`
  ShareUsers []ShareUser `json:"shareUsers"`
}

var db *sql.DB

func Init() {
  db = DB.DB
}

func GetTodoList(user string) []Todo {
  stmt, err := db.Prepare(`
    SELECT todo.id, todo, date_format(date, '%Y/%m/%d %H:%i'), todo.user, done
    FROM todo
    INNER JOIN share
       ON todo.id = share.id
    WHERE share.user = ? AND deleteFlag != true
    ORDER BY date IS NULL ASC, date
  `)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  row, err := stmt.Query(user)
  if err != nil {
    log.Fatal(err)
  }

  stmt, err = db.Prepare(`
    SELECT user, done
    FROM share
    WHERE id = ?
  `)
  if err != nil {
    log.Fatal(err)
  }

  var todoList []Todo
  for row.Next() {
    var todo Todo
    nullDate := sql.NullString{}
    row.Scan(&todo.TodoId, 
             &todo.TodoContent,
             &nullDate,
             &todo.CreatedUser,
             &todo.DoneFlag)

    if (nullDate.Valid) {
      todo.TodoDate = nullDate.String
    }
    rowShare, err := stmt.Query(todo.TodoId)
    if err != nil {
      log.Fatal(err)
    }
    for rowShare.Next() {
      var shareUser ShareUser
      rowShare.Scan(&shareUser.User, &shareUser.DoneFlag)
      todo.ShareUsers = append(todo.ShareUsers, shareUser)
    }

    todoList = append(todoList, todo)
  }

  return todoList
}

func GetUserList() []User {
  sql, err := db.Prepare("SELECT id FROM users")
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
    err := row.Scan(&user.Id)
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

func RegisterTodo(userId, todo, date string , shareUsers []string) bool {
  // NULL check
  if userId == ""  || todo == "" {
    return false
  } 
  nullDate := sql.NullString{}
  if date != "" {
    nullDate.String, nullDate.Valid = date, true
  }
  isOwnUserExist := false
  for _, user := range shareUsers {
    if user == userId {
      isOwnUserExist = true
      break
    }
  }
  if !isOwnUserExist {
    shareUsers = append(shareUsers, userId)
  }

  // Transaction
  tx, err := db.Begin()
  if err != nil {
    return false
  }

  sql, err := tx.Prepare("INSERT INTO todo(todo, date, user) VALUES(?, ?, ?)")
  if err != nil {
    tx.Rollback()
    return false
  }

  _, err = sql.Exec(todo, nullDate, userId)
  if err != nil {
    tx.Rollback()
    fmt.Println(err)
    return false
  }
  var lastInsertId string
  err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastInsertId)
  if err != nil {
    tx.Rollback()
    return false;
  }

  for _, user := range shareUsers {
    sql, err := tx.Prepare("INSERT INTO share(id, user) VALUES(?, ?)")
    if err != nil {
      tx.Rollback()
      return false
    }

    _, err = sql.Exec(lastInsertId, user)
    if err != nil {
      tx.Rollback()
      return false
    }
  }

  if err := tx.Commit(); err != nil {
    tx.Rollback();
    return false;
  }

  return true
}

func DoneTodo(user string, todoid int) bool {
  // NULL Check
  if user == "" {
    return false
  }

  sql, err := db.Prepare("UPDATE share SET done = true WHERE user = ? AND id = ?")
  if err != nil {
    return false
  }
  defer sql.Close()

  _, err = sql.Exec(user, todoid)
  return err == nil
}

func DeleteTodo(user string, todoid int) bool {
  // NULL Check
  if user == "" {
    return false
  }

  sql, err := db.Prepare("UPDATE todo SET deleteFlag = true WHERE user = ? AND id = ?")
  if err != nil {
    return false
  }
  defer sql.Close()

  _, err = sql.Exec(user, todoid)
  return err == nil
}