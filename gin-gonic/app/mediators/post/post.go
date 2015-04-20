package post

import (
    "../../models/post"
    "github.com/gin-gonic/gin"
    "errors"
  . "strconv"
    "fmt"
)

type postMediator struct {
  Post *post.Post
  Context *gin.Context
}

func (self *postMediator) Find() (*postMediator, error) {
  var err error

  self.Post, err = self.Post.Find()

  return self, err
}

func (self *postMediator) Create() (*postMediator, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Create() }

  return self, errors.New("can't parse Post data")
}

func (self *postMediator) Update() (*postMediator, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Update() }

  return self, errors.New("can't parse Post data")
}

func (self *postMediator) Destroy() (error) {
  return self.Post.Destroy()
}

func (self *postMediator) Render() {
  self.Context.JSON(200, self.Post)
}

func (self *postMediator) Destroyed() () {
  msg := fmt.Sprintf("post with id:%d was successfully destroyed", self.Post.Id)
  self.Context.JSON(200, gin.H{"message": msg})
}

func (self *postMediator) Render_400(err error) {
  self.Context.JSON(400, gin.H{"error": err.Error()})
}

func (self *postMediator) Render_404() {
  self.Context.JSON(404, gin.H{"error": notFound(self.Post.Id) })
}

func (self *postMediator) Render_422(err error) {
  self.Context.JSON(422, gin.H{"error": err.Error()})
}

func getIdFrom(context *gin.Context) (id int64, status bool) {
  status = true

  id, err := ParseInt(context.Params.ByName("id"), 10, 64)

  if err != nil { status = false }

  return
}

func notFound(id int64) string {
  return fmt.Sprintf("post with id:%d was not found", id)
}

func Post(context *gin.Context) (p *postMediator) {
  p = &postMediator{Post: &post.Post{}, Context: context}

  if id, ok := getIdFrom(context); ok { p.Post.Id = id }

  return
}
