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
