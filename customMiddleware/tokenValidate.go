package customMiddleware

import (
	"documentService/config"
	"documentService/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var httpClient *config.HttpClientConfig

func SetHttpClient(httpConfig *config.HttpClientConfig) {
	httpClient = httpConfig
}

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authUser := c.Get("user").(*jwt.Token)
		userId := utils.GetUserId(authUser)
		url := httpClient.UserService.BaseUrl + "api/user/validate/" + userId
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		if resp.StatusCode != 200 {
			return c.JSON(401, "invalid token")
		}
		return next(c)
	}
}

func UseAuthorization(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authUser := c.Get("user").(*jwt.Token)
			if err := utils.Authorize(authUser, &roles); err != nil {
				return c.JSON(err.StatusCode, err)
			}
			return next(c)
		}
	}
}
