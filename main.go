package main

import (
	"log"
	"net/http"

	"github.com/b-b3rn4rd/omni-task-two/resouces/home"

	"github.com/emicklei/go-restful"
)

func main() {
	restful.Add(home.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// go run -ldflags "-X github.com/b-b3rn4rd/omni-task-two/resouces/home.Version=1.0.1" main.go
