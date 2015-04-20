package post

import (
  "../../models/post"
  "github.com/gin-gonic/gin"
)

type postsMediator struct {
  Posts *post.Posts
  Context *gin.Context
  Collection []*post.Post
}

func (self *postsMediator) Find() (*postsMediator, error) {
  var err error

  self.Collection, err = self.Posts.Find()

  return self, err
}

func (self *postsMediator) Render() {
  self.Context.JSON(200, self.Collection)
}

func (self *postsMediator) Render_400(err error) {
  self.Context.JSON(400, gin.H{"error": err.Error()})
}

func Posts(context *gin.Context) *postsMediator {
  return &postsMediator{Posts: &post.Posts{}, Context: context}
}