package auth

import (
	. "komeet/core"
)

func RegisterRoutes() {
	App.Router.POST("login", login)
	App.Router.POST("logout", logout)

	App.Router.GET("me", profile)
	App.Router.PUT("me", editProfile)
}
