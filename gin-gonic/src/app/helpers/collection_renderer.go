package helpers

import "github.com/gin-gonic/gin"


type CollectionRenderer interface {
  Find() (*CR, error)
  ToJSON( code int, obj interface{} )
  GetCollection() interface{}
}

type CR struct { CollectionRenderer }

func ( r *CR ) Render() {
  r.ToJSON( 200, r.GetCollection() )
}

func ( r *CR ) Render_400( err error ) {
  r.ToJSON( 400, gin.H{ "error": err.Error() } )
}
