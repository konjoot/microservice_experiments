package middleware

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/gocraft/dbr"
  _ "github.com/lib/pq"
    "../config"
)

var dbconn *dbr.Connection

func init() {
  db, err := sql.Open("postgres", config.DbParams())
  if err != nil {
    panic(err)
  }
  db.Ping()
  db.SetMaxIdleConns(7)
  db.SetMaxOpenConns(75)

  dbconn = dbr.NewConnection(db, nil)
}

func DbSession() gin.HandlerFunc {
  return func(c *gin.Context) {
    dbSession := dbconn.NewSession(nil)
    c.Set("DB", dbSession)
  }
}

func GetDb(c *gin.Context) *dbr.Session {
  db, err := c.Get("DB")
  if err != nil {
    panic(err)
  }

  return db.(*dbr.Session)
}
