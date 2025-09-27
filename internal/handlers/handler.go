package handlers

import (
	"github.com/graphzc/go-task-clean-example/internal/handlers/auth"
	"github.com/graphzc/go-task-clean-example/internal/handlers/common"
)

type Handlers struct {
	Common common.Handler
	Auth   auth.Handler
}

// @WireSet("Handler")
func NewHandlers(
	commonHandler common.Handler,
	authHandler auth.Handler,
) *Handlers {
	return &Handlers{
		Common: commonHandler,
		Auth:   authHandler,
	}
}
