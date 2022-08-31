package main

import (
	"flag"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SingleSession(r *gin.Engine) {
	// Create store
	store := cookie.NewStore()

	// Session's name -> cookie's key
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		// Get session from context
		session := sessions.Default(c)
		var count int

		// Get session's value
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		// Set session's value
		session.Set("count", count)

		// Save session's value
		session.Save()

		c.JSON(200, gin.H{"count": count})
	})

	r.Run(":9900")

}

func MutiSessions(r *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))

	sessionNames := []string{"foo", "bar"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/hello", func(c *gin.Context) {
		sessionFoo := sessions.DefaultMany(c, "foo")
		sessionBar := sessions.DefaultMany(c, "bar")

		if sessionFoo.Get("hello") != "world!" {
			sessionFoo.Set("hello", "world!")
			sessionFoo.Save()
		}

		if sessionBar.Get("hello") != "world?" {
			sessionBar.Set("hello", "world?")
			sessionBar.Save()
		}

		c.JSON(http.StatusOK, gin.H{
			"foo": sessionFoo.Get("hello"),
			"bar": sessionFoo.Get("hello"),
		})

	})
}

func main() {
	flag.Parse()

	r := gin.Default()

	SingleSession(r)
	MutiSessions(r)

	r.Run(":9900")
}
