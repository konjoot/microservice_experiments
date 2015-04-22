package models

import (
  "time"
  "github.com/andevery/dbr"
)

type Post struct {
  Id           int64           `db:"id"          json:"id"                      `
  Title        dbr.NullString  `db:"title"       json:"title"       form:"title"`
  Body         string          `db:"body"        json:"body"        form:"body" `
  CreatedAt    time.Time       `db:"created_at"  json:"created_at"              `
  UpdatedAt    time.Time       `db:"updated_at"  json:"updated_at"              `
}

func (p *Post) Find() (post *Post, err error) {
  post = p

  err = dbSession().Select("*").From("posts").Where("id = ?", p.Id).LoadStruct(p)

  return
}

func (p *Post) Create() (err error) {
  p.CreatedAt = time.Now()
  p.UpdatedAt = p.CreatedAt

  _, err = dbSession().InsertInto("posts").
    Columns("title", "body", "created_at", "updated_at").Record(p).Exec()

  // load Id of newely created post
  if err == nil {
    dbSession().Select("id").From("posts").
      Where("title = ? AND body = ? AND created_at = ? AND updated_at = ?", p.Title, p.Body, p.CreatedAt, p.UpdatedAt).
        LoadStruct(p)
  }

  return
}

func (p *Post) Update() (err error) {
  p.UpdatedAt = time.Now()

  _, err = dbSession().Update("posts").
    Set("title", p.Title).
    Set("body", p.Body).
    Set("updated_at", p.UpdatedAt).
    Where("id = ?", p.Id).Exec()

  return
}

func (p *Post) Destroy() (err error) {
  _, err = dbSession().DeleteFrom("posts").
    Where("id = ?", p.Id).Exec()

  return
}
