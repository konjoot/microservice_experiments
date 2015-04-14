package models

import (
    "database/sql"
    "github.com/gocraft/dbr"
  _ "github.com/lib/pq"
    "../../config"
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

type PostData struct {
  Id           int64           `db:"id"          json:"id"`
  Title        dbr.NullString  `db:"title"       json:"title"`
  Body         dbr.NullString  `db:"body"        json:"body"`
  CreatedAt    dbr.NullTime    `db:"created_at"  json:"created_at"`
  UpdatedAt    dbr.NullTime    `db:"updated_at"  json:"updated_at"`
}

type Post struct {}

func (p Post) All() (posts []*PostData, err error) {
  db := dbconn.NewSession(nil)

  _, err = db.Select("*").From("posts").LoadStructs(&posts)

  return posts, err
}

func (p Post) Find(id string) (post PostData, err error) {
  db := dbconn.NewSession(nil)

  err = db.Select("*").From("posts").Where("id = ?", id).LoadStruct(&post)

  return post, err
}
