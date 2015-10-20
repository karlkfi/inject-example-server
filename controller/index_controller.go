package controller

import (
    "net/http"
    "fmt"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
    return &IndexController{}
}

func (c *IndexController) Name() string {
    return "IndexController"
}

func (c *IndexController) RegisterHandlers(server MuxServer) {
    server.HandleFunc("/", c.Handle)
}

func (c *IndexController) Handle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to the Inject Example Server!")
}
