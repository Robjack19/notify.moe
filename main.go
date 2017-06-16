package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages/airing"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/awards"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/users"
	"github.com/animenotifier/notify.moe/utils"
)

var app = aero.New()

func main() {
	// HTTPS
	app.Security.Load("security/fullchain.pem", "security/privkey.pem")

	// CSS
	app.SetStyle(components.CSS())

	// Session store
	app.Sessions.Store = arn.NewAerospikeStore("Session")

	// Layout
	app.Layout = layout.Render

	// Ajax routes
	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime", search.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)
	app.Ajax("/posts/:id", posts.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/user/:nick/threads", profile.GetThreadsByUser)
	app.Ajax("/users", users.Get)
	app.Ajax("/airing", airing.Get)
	app.Ajax("/awards", awards.Get)
	// app.Ajax("/genres", genres.Get)
	// app.Ajax("/genres/:name", genre.Get)

	// Middleware
	app.Use(middleware.Log())
	app.Use(middleware.Session())

	// API
	api := api.New("/api/", arn.DB)
	api.Install(app)

	// Domain
	if utils.IsDevelopment() {
		app.Config.Domain = "beta.notify.moe"
	}

	// Authentication
	auth.Install(app)

	// Let's go
	app.Run()
}
