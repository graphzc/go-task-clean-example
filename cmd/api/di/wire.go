//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/graphzc/go-task-clean-example/cmd/api/server"
)

func InitializeAPI() *server.EchoServer {
	wire.Build(
		ConfigSet,
		InfrastructureSet,
		HandlerSet,
		RepositorySet,
		ServiceSet,
		server.NewEchoServer,
	)

	return &server.EchoServer{}
}
