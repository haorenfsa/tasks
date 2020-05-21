package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// HTTPServer can run to serve http requests
type HTTPServer struct {
	engine     *gin.Engine
	rootRouter gin.IRouter
}

// NewHTTPServer builds a New HTTPServer
func NewHTTPServer(middlewares ...gin.HandlerFunc) *HTTPServer {
	engine := gin.Default()
	engine.Use(middlewares...)

	return &HTTPServer{
		engine:     engine,
		rootRouter: engine.Group("/api/v1"),
	}
}

// UseControllers registers given controllers to rootRouter
func (h *HTTPServer) UseControllers(ctrls []Controller) {
	for _, ctrl := range ctrls {
		ctrl.Register(h.rootRouter)
	}
}

// ServeStaticPath serves given static path
func (h *HTTPServer) ServeStaticPath(path string) {
	if path == "" {
		return
	}
	h.engine.Use(static.Serve("/", static.LocalFile(path, true)))
}

// Controller can register request handler
type Controller interface {
	Register(gin.IRouter)
}

// Run runs the server
func (h *HTTPServer) Run(addr string) error {
	return h.engine.Run(addr)
}
