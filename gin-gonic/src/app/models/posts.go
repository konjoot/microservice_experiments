package models

type Posts struct {}

func (p *Posts) Find() (posts []*Post, err error) {
  _, err = dbSession().Select("*").From("posts").LoadStructs(&posts)

  return
}
