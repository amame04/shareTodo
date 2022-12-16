package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"

	"sampleapp/DB"
	"sampleapp/user"
)

var db *sql.DB

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("session", store))
	//r.Static("/views", "./views")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
		},
		AllowHeaders: []string{
			"X-CSRF-TOKEN",
		},
		AllowCredentials: true,
	}))

	secret := func(length int) string {
		const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		b := make([]byte, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}
	r.Use(csrf.Middleware(csrf.Options{
		Secret: secret(32),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	//db
	DB.DBinit()
	db = DB.DB
	defer db.Close()

	//user init
	user.Init()

	//routing
	r.GET("/user", func(c *gin.Context) {
		session := sessions.Default(c)
		loginedUser := session.Get("id")

		c.JSON(http.StatusOK, gin.H{"user": loginedUser, "token": csrf.GetToken(c)})
	})

	r.GET("/todoList", func(c *gin.Context) {
		session := sessions.Default(c)
		loginedUser := session.Get("id")
		loginedUserStr, ok := loginedUser.(string)

		var todoList []user.Todo
		if ok && loginedUserStr != "" {
			todoList = user.GetTodoList(loginedUserStr)
		}

		c.JSON(http.StatusOK, gin.H{"user": loginedUser, "todoList": todoList})
	})

	r.GET("/userList", func(c *gin.Context) {
		userList := user.GetUserList()
		c.JSON(http.StatusOK, gin.H{"userList": userList})
	})

	r.POST("/login", func(c *gin.Context) {
		id := c.PostForm("id")
		pass := c.PostForm("password")
		success := user.LoginUser(id, pass)

		session := sessions.Default(c)
		session.Clear()
		if success {
			session.Set("id", id)
		}
		session.Save()

		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.POST("/register", func(c *gin.Context) {
		id := c.PostForm("id")
		pass := c.PostForm("password")
		success := user.RegisterUser(id, pass)
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.POST("/registerTodo", func(c *gin.Context) {
		session := sessions.Default(c)
		loginedUser := session.Get("id")
		loginedUserStr, ok := loginedUser.(string)

		todo := c.PostForm("todo")
		date := c.PostForm("date")
		share := c.PostFormArray("shareUsers")
		success := ok && user.RegisterTodo(loginedUserStr, todo, date, share)
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.PATCH("/doneTodo", func(c *gin.Context) {
		session := sessions.Default(c)
		loginedUser := session.Get("id")
		loginedUserStr, ok := loginedUser.(string)

		todoId, err := strconv.Atoi(c.PostForm("todoId"))
		success := ok && err == nil && user.DoneTodo(loginedUserStr, todoId)
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.DELETE("/deleteTodo", func(c *gin.Context) {
		session := sessions.Default(c)
		loginedUser := session.Get("id")
		loginedUserStr, ok := loginedUser.(string)

		todoId, err := strconv.Atoi(c.Query("todoId"))
		success := ok && err == nil && user.DeleteTodo(loginedUserStr, todoId)
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.Run(":8888")
}
