package html

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getTemplateByName(c *gin.Context) {
	path := c.Query("path")
	body := make(map[string]any)
	if err := c.ShouldBind(&body); err != nil {
		logrus.Info("bind error", err)
	}
	body["post"] = map[string]any{"ID": 1}
	body["user"] = map[string]any{"ID": 1, "Username": "manuelarte"}
	body["csrf"] = "csrf_value"
	body["comment"] = nil
	c.HTML(http.StatusOK, path, body)
}

func RegisterDebugHandlers(r gin.IRouter) {
	r.GET("/debug", getTemplateByName)
}
