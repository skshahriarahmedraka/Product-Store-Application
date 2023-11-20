package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func (H *DatabaseCollections) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.SetCookie("Auth1", "", -1, "/", os.Getenv("DOMAIN_ADDR"), false, true)
		c.SetCookie("Refresh", "", -1, "/", os.Getenv("DOMAIN_ADDR"), false, true)
		logger.Info().Msg("📢 Info message :" + "user logged out")
		c.Redirect(http.StatusTemporaryRedirect, "/public")

	}
}
