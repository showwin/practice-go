package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/css", static.LocalFile("css", true)))
	store := sessions.NewCookieStore([]byte("showwin_happy"))
	store.Options(sessions.Options{HttpOnly: true})
	r.Use(sessions.Sessions("mysssion", store))

	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("uid", 1)
		session.Save()

		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/", func(c *gin.Context) {
		addUID(sessions.Default(c))
		r.SetHTMLTemplate(template.Must(template.ParseFiles("test.tmpl")))
		c.HTML(http.StatusOK, "base", gin.H{
			"CurrentUser": "aaa",
			"Products":    "uuu",
		})
	})

	r.Run(":8080")
}

func addUID(session sessions.Session) {
	uid := session.Get("uid")
	log.Print(1 + uid.(int))
}
