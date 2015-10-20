package main

import (
    "github.com/karlkfi/inject"
    "github.com/karlkfi/inject-example-server/controller"
    "github.com/karlkfi/inject-example-server/model"

    "github.com/facebookgo/grace/gracehttp"

    "net/http"
    "log"
)

func main() {
    graph := objectGraph()

    var mux controller.MuxServer
    inject.ExtractAssignable(graph, &mux)

    var controllers []controller.Controller
    inject.FindAssignable(graph, &controllers)
    for _, c := range controllers {
        log.Println("Registering controller:", c.Name())
        c.RegisterHandlers(mux)
    }

    // serve and listen for shutdown signals
    gracehttp.Serve(
        &http.Server{Addr: "0.0.0.0:8080", Handler: mux},
    )
}

func objectGraph() inject.Graph {
    var (
        server *http.ServeMux
        userRepo model.UserRepo
        users = users()
        indexController *controller.IndexController
        userController *controller.UserController
    )

    graph := inject.NewGraph()

    graph.Define(&server, inject.NewProvider(http.NewServeMux))

    graph.Define(&userRepo, inject.NewProvider(model.NewUserRepo, &users))

    graph.Define(&indexController, inject.NewProvider(controller.NewIndexController))
    graph.Define(&userController, inject.NewProvider(controller.NewUserController, &userRepo))

    return graph
}

func users() []model.User {
    return []model.User{
        { ID: "0", Name: "Alice" },
        { ID: "1", Name: "Bob" },
        { ID: "2", Name: "Charlie" },
    }
}
