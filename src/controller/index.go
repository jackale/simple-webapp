package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Posts",
		"content": "This is a content from Server",
	})
}
