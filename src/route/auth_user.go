package route

import (
	"net-http/myapp/controller"
)

func (router *Router) GetAuthRouter() {
	router.Mutex.HandleFunc("/api/v1/auth", controller.HandlerTwo)
}
