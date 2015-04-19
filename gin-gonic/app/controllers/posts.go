package controllers

import (
  "github.com/gin-gonic/gin"
  . "../models/post"
  . "strconv"
  "errors"
  "fmt"
)


type PostsController struct {}

func (p PostsController) Index(c *gin.Context){
  posts, err := posts(c).Find()

  if err != nil { render_400(c, err); return }

  c.JSON(200, posts)
}

func (p PostsController) Show(c *gin.Context){
  post, err := post(c).Find()

  if err != nil { post.Render_404(); return }

  post.Render()
}

func (p PostsController) Create(c *gin.Context){
  post, err := post(c).Create()

  if err != nil { post.Render_422(err); return }

  post.Render()
}

func (p PostsController) Update(c *gin.Context){
  post, err := post(c).Find()

  if err != nil { post.Render_404(); return }

  post, err = post.Update()

  if err != nil { post.Render_422(err); return }

  post.Render()
}

func (p PostsController) Destroy(c *gin.Context){
  post, err := post(c).Find();

  if err != nil { post.Render_404(); return }

  err = post.Destroy()

  if err != nil { post.Render_400(err); return }

  post.Destroyed()
}

type PostProc struct {
  Post *Post
  Context *gin.Context
}

func (self *PostProc) Find() (*PostProc, error) {
  var err error

  self.Post, err = self.Post.Find()

  return self, err
}

func (self *PostProc) Create() (*PostProc, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Create() }

  return self, errors.New("can't parse Post data")
}

func (self *PostProc) Update() (*PostProc, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Update() }

  return self, errors.New("can't parse Post data")
}

func (self *PostProc) Destroy() (error) {
  return self.Post.Destroy()
}

func (self *PostProc) Render() {
  self.Context.JSON(200, self.Post)
}

func (self *PostProc) Destroyed() () {
  msg := fmt.Sprintf("post with id:%d was successfully destroyed", self.Post.Id)
  self.Context.JSON(200, gin.H{"message": msg})
}

func (self *PostProc) Render_400(err error) {
  self.Context.JSON(400, gin.H{"error": err.Error()})
}

func (self *PostProc) Render_404() {
  self.Context.JSON(404, gin.H{"error": notFound(self.Post.Id) })
}

func (self *PostProc) Render_422(err error) {
  self.Context.JSON(422, gin.H{"error": err.Error()})
}

func posts(context *gin.Context) *Posts { return &Posts{} }

func post(context *gin.Context) (p *PostProc) {
  p = &PostProc{Post: &Post{}, Context: context}

  if id, ok := getIdFrom(context); ok { p.Post.Id = id }

  return
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

func render_400(c *gin.Context, err error) {
  c.JSON(400, gin.H{"error": err.Error()})
}
