package main

import (
	"github.com/tonychinwe/libraryone/src/repository"
	"github.com/tonychinwe/libraryone/src/routes"
)

func main() {

	repository.InitDb()
	routes.InitRouter()

}
