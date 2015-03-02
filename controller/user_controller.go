package controller

import (
    "net/http"
    "encoding/json"
    "regexp"

    "github.com/karlkfi/inject-example-server/model"
)

var exp = regexp.MustCompile("/user/(.+)")

type UserController struct {
    repo model.UserRepo
}

func NewUserController(repo model.UserRepo) *UserController {
    return &UserController{repo: repo}
}

func (c *UserController) RegisterHandlers(server *http.ServeMux) {
    server.HandleFunc("/user/", c.HandleUserRequest)
    server.HandleFunc("/users", c.HandleUsersRequest)
}

func (c *UserController) HandleUserRequest(w http.ResponseWriter, r *http.Request) {
    submatch := exp.FindStringSubmatch(r.RequestURI)
    if submatch == nil {
        http.Error(w, "Invalid User Request", http.StatusBadRequest)
        return
    }

    userID := submatch[1]
    user, found := c.repo.FindUser(userID)
    if !found {
        http.NotFound(w, r)
        return
    }

    bytes, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(bytes)
}

func (c *UserController) HandleUsersRequest(w http.ResponseWriter, r *http.Request) {
    userList := c.repo.All()

    bytes, err := json.Marshal(userList)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(bytes)
}
