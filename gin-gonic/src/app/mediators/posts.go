package mediators

import (
    "app/models"
    "github.com/gin-gonic/gin"
  . "app/helpers"
)

type postsMediator struct {
  Posts *models.Posts
  Context *gin.Context
  Collection []*models.Post
}

func (self *postsMediator) Find() (*CR, error) {
  var err error

  self.Collection, err = self.Posts.Find()

  return &CR{self}, err
}

func (self *postsMediator) ToJSON( code int, obj interface{} ) {
  self.Context.JSON( code, obj )
}

func (self *postsMediator) GetCollection() interface{} {
  return self.Collection
}

// Constructor
func Posts(context *gin.Context) *postsMediator {
  return &postsMediator{Posts: &models.Posts{}, Context: context}
}