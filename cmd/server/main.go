package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	api "gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/api/tasks"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/auth"
	_ "gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/docs"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/middleware"
)

// @title           TODO API
// @version         1.0
// @description     A simple example API with auth
// @BasePath        /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	pollServer := api.NewPollServer()

	router := gin.Default()
	router.Use(middleware.EnsureVoterID())

	docs := router.Group("/docs", gin.BasicAuth(auth.GetBasicAuthUsers()))
	ui := router.Group("/ui")
	docs.Use()
	ui.Use()

	// docs.GET("/*filepath", func(c *gin.Context) {
	// 	http.ServeFile(c.Writer, c.Request, "./swagger-ui"+c.Param("filepath"))
	// })

	ui.GET("/*filepath", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./ui/index.html")
	})

	api.RegisterHandlers(router, pollServer)

	// And we serve HTTP until the world ends.
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
