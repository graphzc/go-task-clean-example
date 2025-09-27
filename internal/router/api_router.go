package router

import (
	"net/http"

	"github.com/graphzc/go-task-clean-example/internal/utils/echoutil"
)

func (r *Router) RegisterAPIRoutes() {
	// Health check
	r.echo.GET("/health", echoutil.WrapWithStatus(r.handlers.Common.HealthCheck, http.StatusOK))

	v1Public := r.echo.Group("/api/v1")

	authGroupV1 := v1Public.Group("/auth")
	{
		authGroupV1.POST("/register", echoutil.WrapWithStatus(r.handlers.Auth.Register, http.StatusCreated))
		authGroupV1.POST("/login", echoutil.WrapWithStatus(r.handlers.Auth.Login, http.StatusOK))
	}

}
