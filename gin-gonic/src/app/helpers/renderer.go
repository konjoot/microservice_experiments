package helpers

import "github.com/gin-gonic/gin"


type Renderer interface {
  Find() (*R, error)
  Create() (*R, error)
  Update() (*R, error)
  Destroy() (error)
  Self() interface{}
  ToJSON( code int, obj interface{} )
  NotFound() string
  Destroyed() string
}

type R struct{ Renderer }

func (r *R) Render() {
  r.ToJSON(200, r.Self())
}

func (r *R) RenderDestroyed() () {
  r.ToJSON(200, gin.H{ "message": r.Destroyed() })
}

func (r *R) Render_400(err error) {
  r.ToJSON(400, gin.H{ "error": err.Error() })
}

func (r *R) Render_404() {
  r.ToJSON(404, gin.H{ "error": r.NotFound() })
}

func (r *R) Render_422(err error) {
  r.ToJSON(422, gin.H{ "error": err.Error() })
}
