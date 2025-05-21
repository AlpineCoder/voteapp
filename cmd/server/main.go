package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	api "gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/api/tasks"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/auth"
	_ "gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/docs"
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
	taskServer := api.NewTaskServer()

	router := gin.Default()
	// Swagger UI route using basic auth
	router.GET("/swagger/*any", auth.BasicAuthMiddleware("admin", "password"), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register the API routes with the Gin router using the generated code

	// group := router.Group("/tasks", taskServer.GetTasks, auth.ApiKeyAuthMiddleware("key") )

	// this is how we register the handlers for all the routes for the task server and
	// inject our custom middleware. in this case a hard coded api key
	api.RegisterHandlersWithOptions(router, taskServer, api.GinServerOptions{
		Middlewares: []api.MiddlewareFunc{
			api.MiddlewareFunc(auth.AlternateApiKeyAuthMiddleware),
		},
	})
	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
