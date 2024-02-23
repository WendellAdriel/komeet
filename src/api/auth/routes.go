package auth

import (
	v1 "komeet/api/auth/v1"
	. "komeet/core"
)

func RegisterRoutes() {
	v1Routes := App.Router.Group("v1")

	v1Routes.POST("login", v1.Login)
	v1Routes.POST("logout", v1.Logout)

	v1Routes.GET("me", v1.Profile)
	v1Routes.PUT("me", v1.EditProfile)
}
