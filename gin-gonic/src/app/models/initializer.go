package models

import (
    "database/sql"
    "github.com/andevery/dbr"
  _ "github.com/lib/pq"
    "config"
)

var dbConn *dbr.Connection

func init() {
  db, err := sql.Open("postgres", config.DbParams())
  if err != nil {
    panic(err)
  }
  db.Ping()
  db.SetMaxIdleConns(7)
  db.SetMaxOpenConns(75)
  dbr.Quoter = dbr.PostgresQuoter{}

  dbConn = dbr.NewConnection(db, nil)
}

func dbSession() *dbr.Session { return dbConn.NewSession(nil) }
