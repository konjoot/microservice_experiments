package controllers

import (
  "github.com/gin-gonic/gin"
  . "app/mediators/post"
  . "app/helpers/renderers"
)


type PostsController struct {}

func (p PostsController) Index(c *gin.Context){
  posts, err := Posts(c).Find()

  if err != nil { posts.Render_400(err); return }

  posts.Render()
}

func (p PostsController) Show(c *gin.Context){
  post, err := Post(c).Find()

  if err != nil { post.Render_404(); return }

  post.Render()
}

func (p PostsController) Create(c *gin.Context){
  post, err := Post(c).Create()

  if err != nil { post.Render_422(err); return }

  post.Render()
}

func (p PostsController) Update(c *gin.Context){
  post, err := Post(c).Find()

  if err != nil { post.Render_404(); return }

  post, err = post.Update()

  if err != nil { post.Render_422(err); return }

  post.Render()
}

func (p PostsController) Destroy(c *gin.Context){
  post, err := Post(c).Find()

  if err != nil { post.Render_404(); return }

  err = post.Destroy()

  if err != nil { post.Render_400(err); return }

  post.RenderDestroyed()
}
