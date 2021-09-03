package editor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadHandle(c *gin.Context) {
	c.String(http.StatusOK, "read file todo!")
}
