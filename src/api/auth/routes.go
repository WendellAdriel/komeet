package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router **gin.RouterGroup) {
	(*router).POST("login", login)
	(*router).POST("logout", logout)

	(*router).GET("me", profile)
	(*router).PUT("me", editProfile)
}
