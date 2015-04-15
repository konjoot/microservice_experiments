package post

import (
    "database/sql"
    "github.com/gocraft/dbr"
  _ "github.com/lib/pq"
    "../../../config"
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

  dbConn = dbr.NewConnection(db, nil)
}

func dbSession() *dbr.Session { return dbConn.NewSession(nil) }

type Post struct {
  Id           int64           `db:"id"          json:"id"`
  Title        dbr.NullString  `db:"title"       json:"title"`
  Body         dbr.NullString  `db:"body"        json:"body"`
  CreatedAt    dbr.NullTime    `db:"created_at"  json:"created_at"`
  UpdatedAt    dbr.NullTime    `db:"updated_at"  json:"updated_at"`
}

func (p Post) Find() (post *Post, err error) {
  post = &p

  err = dbSession().Select("*").From("posts").Where("id = ?", p.Id).LoadStruct(&p)

  return
}

type Posts struct {}

func (p Posts) Find() (posts []*Post, err error) {
  _, err = dbSession().Select("*").From("posts").LoadStructs(&posts)

  return
}
