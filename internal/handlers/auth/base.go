package auth

import (
	"context"

	"github.com/graphzc/go-task-clean-example/internal/dto"
	"github.com/graphzc/go-task-clean-example/internal/services/user"
)

type Handler interface {
	Register(ctx context.Context, req *dto.UserRegisterRequest) (*dto.MessageResponse, error)
	Login(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
}

type handler struct {
	service user.Service
}

// @WireSet("Handler")
func New(service user.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(ctx context.Context, req *dto.UserRegisterRequest) (*dto.MessageResponse, error) {
	userRegisterInput := &user.UserRegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.Register(ctx, userRegisterInput); err != nil {
		return nil, err
	}

	return &dto.MessageResponse{
		Message: "User registered successfully",
	}, nil
}

func (h *handler) Login(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	serviceInput := user.UserLoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	accessToken, err := h.service.Login(ctx, &serviceInput)
	if err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{
		AccessToken: accessToken,
	}, nil
}
