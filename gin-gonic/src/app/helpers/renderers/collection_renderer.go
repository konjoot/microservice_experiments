package renderer

import "github.com/gin-gonic/gin"


type CollectionRenderer interface {
  ToJSON( code int, obj interface{} )
  GetCollection() []interface{}
}

func ( c *CollectionRenderer ) Render() {
  c.ToJSON( 200, c.GetCollection() )
}

func ( c *CollectionRenderer ) Render_400( err error ) {
  c.ToJSON( 400, gin.H{ "error": err.Error() } )
}


// // Old
// func (self *postsMediator) Render() {
//   self.Context.JSON(200, self.Collection)
// }

// func (self *postsMediator) Render_400( err error ) {
//   self.Context.JSON(400, gin.H{"error": err.Error()})
// }