package main

import (
  "net/http"
  "database/sql"

  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  router := gin.Default()

  router.GET("/user/:name", func(c *gin.Context){
    name := c.Param("name")
    c_at, u_at := getUser(name)
    c.JSON(http.StatusOK, gin.H{
      "name": name,
      "created_at": c_at,
      "updated_at": u_at,
    })
  })

  router.GET("/bigjson", func(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
      "name1": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name2": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name3": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name4": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name5": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name6": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name7": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name8": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name9": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name10": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name11": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name12": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name13": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name14": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name15": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name16": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name17": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name18": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name19": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
      "name20": "this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.this is a name.",
    })
  })

  router.Run(":8080")
}

func getUser(name string) (string, string) {
  db, err := sql.Open("mysql", "test:test@/table_name")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  var c_at, u_at string
  err = db.QueryRow("SELECT created_at, updated_at FROM users WHERE name = ? LIMIT 1", name).Scan(&c_at, &u_at)
  if err != nil {
    panic(err.Error())
  }

  return c_at, u_at
}
