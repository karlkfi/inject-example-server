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

func (c *IndexController) RegisterHandlers(server *http.ServeMux) {
    server.HandleFunc("/", c.Handle)
}

func (c *IndexController) Handle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Nothing to see here.")
}
