package router

import (
    "github.com/gin-gonic/gin"
    // "../../middleware"
  . "app/controllers"
)

func Router() *gin.Engine {
  router := gin.Default()

  API := router.Group("/api")

  // API.Use(middleware.DbSession())

  Posts := PostsController{}

  API.GET    ( "/posts",     Posts.Index   )
  API.GET    ( "/posts/:id", Posts.Show    )
  API.PUT    ( "/posts/:id", Posts.Update  )
  API.POST   ( "/posts/",    Posts.Create  )
  API.DELETE ( "/posts/:id", Posts.Destroy )

  return router
}
