package post

import (
    "app/models/post"
    "github.com/gin-gonic/gin"
  . "app/helpers/renderer"
    "errors"
  . "strconv"
    "fmt"
)

type postMediator struct {
  Post *post.Post
  Context *gin.Context
}

func (self *postMediator) Find() (*R, error) {
  var err error

  self.Post, err = self.Post.Find()

  return &R{self}, err
}

func (self *postMediator) Create() (*R, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return &R{self}, self.Post.Create() }

  return &R{self}, errors.New("can't parse Post data")
}

func (self *postMediator) Update() (*R, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return &R{self}, self.Post.Update() }

  return &R{self}, errors.New("can't parse Post data")
}

func (self *postMediator) Destroy() (error) {
  return self.Post.Destroy()
}

func (self *postMediator) Self() interface{} {
  return self.Post
}

func (self *postMediator) ToJSON( code int, obj interface{} ) {
  self.Context.JSON( code, obj )
}

func (self *postMediator) NotFound() string {
  return fmt.Sprintf("post with id:%d was not found", self.Post.Id)
}

func (self *postMediator) Destroyed() string {
  return fmt.Sprintf("post with id:%d was successfully destroyed", self.Post.Id)
}

/*
// helper to get 'id' from context params
*/
func getIdFrom(context *gin.Context) (id int64, status bool) {
  status = true

  id, err := ParseInt(context.Params.ByName("id"), 10, 64)

  if err != nil { status = false }

  return
}

/*
// Constructor
*/
func Post(context *gin.Context) (p *postMediator) {
  p = &postMediator{Post: &post.Post{}, Context: context}

  if id, ok := getIdFrom(context); ok { p.Post.Id = id }

  return
}
