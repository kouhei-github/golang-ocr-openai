package route

import (
	"net-http/myapp/controller"
	"net/http"
)

type Router struct {
	Mutex *http.ServeMux
}

func (router *Router) GetRouter() {
	//router.Mutex.HandleFunc("/", controller.HandlerTwo)
	// 練習
	router.Mutex.HandleFunc("/two", controller.HandlerTwo)

	// OCR
	router.Mutex.HandleFunc("/api/v1/ocr", controller.OcrHandler)
}
