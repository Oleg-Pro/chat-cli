package handler

import (
	"github.com/Oleg-Pro/chat-cli/internal/client/grpc/auth"
	chatServer "github.com/Oleg-Pro/chat-cli/internal/client/grpc/chat_server"
	"github.com/Oleg-Pro/chat-cli/internal/client/redis"
)

type Handler struct {
	redisClient redis.Client
	authClient  auth.Client
	chatClient  chatServer.Client
}

func NewHandler(redisClient redis.Client, authClient auth.Client, chatClient chatServer.Client) *Handler {
	return &Handler{
		redisClient: redisClient,
		authClient:  authClient,
		chatClient:  chatClient,
	}
}
