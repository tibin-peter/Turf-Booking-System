package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
