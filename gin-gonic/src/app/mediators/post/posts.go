package post

import (
    "app/models/post"
    "github.com/gin-gonic/gin"
  . "app/helpers/renderer/collection"
)

type postsMediator struct {
  Posts *post.Posts
  Context *gin.Context
  Collection []*post.Post
}

func (self *postsMediator) Find() (*R, error) {
  var err error

  self.Collection, err = self.Posts.Find()

  return &R{self}, err
}

func (self *postsMediator) ToJSON( code int, obj interface{} ) {
  self.Context.JSON( code, obj )
}

func (self *postsMediator) GetCollection() interface{} {
  return self.Collection
}

// Constructor
func Posts(context *gin.Context) *postsMediator {
  return &postsMediator{Posts: &post.Posts{}, Context: context}
}