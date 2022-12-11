package main

import (
  "net/http"
  "database/sql"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-gonic/gin"

  "sampleapp/DB"
  "sampleapp/user"
)

var db *sql.DB

func main() {
  r := gin.Default()
  store := cookie.NewStore([]byte("secret"))

  r.Use(sessions.Sessions("session", store))
  r.Static("/views", "./views")

  //db
  DB.DBinit()
  db = DB.DB
  defer db.Close()

  //user init
  user.Init()

  //routing
  r.GET("/", func(c *gin.Context){
    session := sessions.Default(c)
    loginedUser := session.Get("id")
    loginedUserStr, ok := loginedUser.(string)

    var todoList []user.Todo
    if ok && loginedUserStr != ""{
      todoList = user.GetTodoList(loginedUserStr)
    }

    c.JSON(http.StatusOK, gin.H{"user" : loginedUser, "todoList" : todoList})
  })

  r.GET("/userList", func(c *gin.Context){
    userList := user.GetUserList()
    c.JSON(http.StatusOK, gin.H{"userList" : userList})
  })

  r.GET("/login", func(c *gin.Context){
    id := c.Query("id")
    pass := c.Query("password")
    success := user.LoginUser(id, pass)

    session := sessions.Default(c)
    session.Clear()
    if success {
      session.Set("id", id)
    }
    session.Save()

    c.JSON(http.StatusOK, gin.H{"success" : success})
  })

  r.GET("/register", func(c *gin.Context){
    id := c.Query("id")
    pass := c.Query("password")
    success := user.RegisterUser(id, pass)
    c.JSON(http.StatusOK, gin.H{"success" : success})
  })

  r.Run(":8080")
}

