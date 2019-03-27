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
