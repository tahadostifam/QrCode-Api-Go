package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func MakeQRCodeAction(c *gin.Context) {
	data := c.Query("message")

	if len(strings.TrimSpace(data)) == 0 {
		c.String(400, "message not found on query")
	} else {
		c.Header("Content-Type", "image/png")

		var png, err = qrcode.Encode(data, qrcode.Medium, 500)
		if err != nil {
			c.String(http.StatusOK, "An error occurred on making the qrcode")
		}

		c.String(http.StatusOK, string(png))
	}
}
