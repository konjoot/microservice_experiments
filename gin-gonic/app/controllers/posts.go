package controllers

import (
    "github.com/gin-gonic/gin"
  . "../models"
//     "github.com/gocraft/dbr"
//     "time"
)

type PostsController struct {}

func (p PostsController) Index(c *gin.Context){
  posts, err := Post{}.All()

  if err != nil {
    c.JSON(404, gin.H{"message": "error"})
    return
  }

  c.JSON(200, posts)
}

func (p PostsController) Show(c *gin.Context){
  post, err := Post{}.Find(c.Params.ByName("id"))

  if (err != nil) {
    c.JSON(404, gin.H{"message": "not found"})
    return
  }

  c.JSON(200, post)
}

func (p PostsController) Create(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "create"})
}

func (p PostsController) Update(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "update"})
}

func (p PostsController) Destroy(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "destroy"})
}
