package renderer

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











// old


// // move to NotFound
// func notFound(id int64) string {
//   return fmt.Sprintf("post with id:%d was not found", id)
// }

// // move to NotFound
// func Destroyed(id int64) string {
//   return fmt.Sprintf("post with id:%d was successfully destroyed", id)
// }

// func (self *postMediator) Render() {
//   self.Context.JSON(200, self.Post)
// }

// func (self *postMediator) RenderDestroyed() () {
//   msg := fmt.Sprintf("post with id:%d was successfully destroyed", self.Post.Id)
//   self.Context.JSON(200, gin.H{"message": msg})
// }

// func (self *postMediator) Render_400(err error) {
//   self.Context.JSON(400, gin.H{"error": err.Error()})
// }

// func (self *postMediator) Render_404() {
//   self.Context.JSON(404, gin.H{"error": notFound(self.Post.Id) })
// }

// func (self *postMediator) Render_422(err error) {
//   self.Context.JSON(422, gin.H{"error": err.Error()})
// }

// func notFound(id int64) string {
//   return fmt.Sprintf("post with id:%d was not found", id)
// }