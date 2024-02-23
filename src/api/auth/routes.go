package auth

import (
	v1 "komeet/api/auth/v1"
	. "komeet/core"
)

func RegisterRoutes() {
	v1Routes := App.Router.Group("v1")
	v1Routes.POST("login", v1.Login)

	protectedRoutes := v1Routes.Group("/")
	protectedRoutes.Use(AuthRequired)

	protectedRoutes.POST("logout", v1.Logout)
	protectedRoutes.GET("me", v1.Profile)
	protectedRoutes.PUT("me", v1.EditProfile)
}
