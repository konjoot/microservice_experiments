package controllers

import (
    "github.com/gin-gonic/gin"
//     "github.com/gocraft/dbr"
//   . "../models"
//     "time"
)

type PostsController struct { }

func (p PostsController) Index(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "index"})
}

func (p PostsController) Show(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "show"})
}

func (p PostsController) Create(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "create"})
}

func (p PostsController) Update(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "update"})
}

func (p PostsController) Replace(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "replace"})
}

func (p PostsController) Destroy(c *gin.Context){
  c.JSON(200, gin.H{"thisIs": "destroy"})
}
