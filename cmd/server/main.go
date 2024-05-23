package main

import (
	"github.com/fdhhhdjd/E-commerce_Api_Golang/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}
