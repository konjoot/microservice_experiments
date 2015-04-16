package post

import (
    "database/sql"
    "github.com/andevery/dbr"
  _ "github.com/lib/pq"
    "../../../config"
    "fmt"
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

type Post struct {
  Id           int64           `db:"id"          json:"id"                      `
  Title        dbr.NullString          `db:"title"       json:"title"       form:"title"`
  Body         dbr.NullString          `db:"body"        json:"body"        form:"body" `
  CreatedAt    dbr.NullTime    `db:"created_at"  json:"created_at"              `
  UpdatedAt    dbr.NullTime    `db:"updated_at"  json:"updated_at"              `
}

type PostForm struct {
  Title        string  `form:"title"`
  Body         string  `form:"body"`
}

func (p Post) Find() (post *Post, err error) {
  post = &p

  err = dbSession().Select("*").From("posts").Where("id = ?", p.Id).LoadStruct(&p)

  return
}

func (p Post) Save() (status bool) {
  status = true
  fmt.Printf("%+v\n", p)

  k, _ := dbSession().InsertInto("posts").
    Columns("title", "body").Record(p).ToSql()
  fmt.Printf("%+v\n", k)

  _, err := dbSession().InsertInto("posts").
    Columns("title", "body").Record(p).Exec()

  if err != nil {
    status = false
    fmt.Printf("%+v\n", err)
  }

  return
}

type Posts struct {}

func (p Posts) Find() (posts []*Post, err error) {
  _, err = dbSession().Select("*").From("posts").LoadStructs(&posts)

  return
}
