package collection

import "github.com/gin-gonic/gin"


type CollectionRenderer interface {
  Find() (*R, error)
  ToJSON( code int, obj interface{} )
  GetCollection() interface{}
}

type R struct { CollectionRenderer }

func ( r *R ) Render() {
  r.ToJSON( 200, r.GetCollection() )
}

func ( r *R ) Render_400( err error ) {
  r.ToJSON( 400, gin.H{ "error": err.Error() } )
}


// // Old
// func (self *postsMediator) Render() {
//   self.Context.JSON(200, self.Collection)
// }

// func (self *postsMediator) Render_400( err error ) {
//   self.Context.JSON(400, gin.H{"error": err.Error()})
// }