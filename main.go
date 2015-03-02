package main

import (
    "net/http"

    "github.com/karlkfi/inject"
    "github.com/karlkfi/inject-example-server/controller"
    "github.com/karlkfi/inject-example-server/model"
)

func main() {
    var (
        users = []model.User{
            { ID: "0", Name: "John Doe" },
            { ID: "1", Name: "Jane Doe" },
        }
        userRepo model.UserRepo
        indexController controller.Controller
        userController controller.Controller
    )

    graph := inject.NewGraph()

    graph.Define(&userRepo, inject.NewProvider(model.NewUserRepo, &users))
    graph.Define(&indexController, inject.NewProvider(controller.NewIndexController))
    graph.Define(&userController, inject.NewProvider(controller.NewUserController, &userRepo))

    graph.Resolve(&indexController)
    graph.Resolve(&userController)

    server := http.NewServeMux()

    indexController.RegisterHandlers(server)
    userController.RegisterHandlers(server)

    http.ListenAndServe(":8080", server)
}