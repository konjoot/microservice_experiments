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

  if err != nil { posts.Render_400(err); return }

  posts.Render()
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
  post, err := post(c).Find()

  if err != nil { post.Render_404(); return }

  err = post.Destroy()

  if err != nil { post.Render_400(err); return }

  post.Destroyed()
}

type PostMediator struct {
  Post *Post
  Context *gin.Context
}

func (self *PostMediator) Find() (*PostMediator, error) {
  var err error

  self.Post, err = self.Post.Find()

  return self, err
}

func (self *PostMediator) Create() (*PostMediator, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Create() }

  return self, errors.New("can't parse Post data")
}

func (self *PostMediator) Update() (*PostMediator, error) {
  ok := self.Context.Bind(self.Post)

  if ok { return self, self.Post.Update() }

  return self, errors.New("can't parse Post data")
}

func (self *PostMediator) Destroy() (error) {
  return self.Post.Destroy()
}

func (self *PostMediator) Render() {
  self.Context.JSON(200, self.Post)
}

func (self *PostMediator) Destroyed() () {
  msg := fmt.Sprintf("post with id:%d was successfully destroyed", self.Post.Id)
  self.Context.JSON(200, gin.H{"message": msg})
}

func (self *PostMediator) Render_400(err error) {
  self.Context.JSON(400, gin.H{"error": err.Error()})
}

func (self *PostMediator) Render_404() {
  self.Context.JSON(404, gin.H{"error": notFound(self.Post.Id) })
}

func (self *PostMediator) Render_422(err error) {
  self.Context.JSON(422, gin.H{"error": err.Error()})
}

type PostsMediator struct {
  Posts *Posts
  Context *gin.Context
  Collection []*Post
}

func (self *PostsMediator) Find() (*PostsMediator, error) {
  var err error

  self.Collection, err = self.Posts.Find()

  return self, err
}

func (self *PostsMediator) Render() {
  self.Context.JSON(200, self.Collection)
}

func (self *PostsMediator) Render_400(err error) {
  self.Context.JSON(400, gin.H{"error": err.Error()})
}

func posts(context *gin.Context) *PostsMediator {
  return &PostsMediator{Posts: &Posts{}, Context: context}
}

func post(context *gin.Context) (p *PostMediator) {
  p = &PostMediator{Post: &Post{}, Context: context}

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
