package controller

import (
    "net/http"
)

type Controller interface {
    RegisterHandlers(*http.ServeMux)
}
