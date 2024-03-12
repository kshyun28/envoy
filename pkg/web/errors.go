package web

import (
	"net/http"
	"self-hosted-node/pkg"

	"github.com/gin-gonic/gin"
)

// Renders the "not found page"
func (s *Server) NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{"Version": pkg.Version()})
}

// Renders the "invalid action page"
func (s *Server) NotAllowed(c *gin.Context) {
	c.HTML(http.StatusMethodNotAllowed, "405.html", gin.H{"Version": pkg.Version()})
}