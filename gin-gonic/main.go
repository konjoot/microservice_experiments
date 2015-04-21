package main

import (
    "config"
  . "config/router"
    "runtime"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  router := Router()
  router.Run(config.ServerParams())
}
