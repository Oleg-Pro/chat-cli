package handler

import (
	myGrps "github.com/Oleg-Pro/chat-cli/internal/client/grpc"		
	"github.com/Oleg-Pro/chat-cli/internal/client/redis"
)

type Handler struct {
	redisClient redis.Client
	authClient  myGrps.AuthClient
	chatClient  myGrps.ChatServerClient
}

func NewHandler(redisClient redis.Client, authClient myGrps.AuthClient, chatClient myGrps.ChatServerClient) *Handler {
	return &Handler{
		redisClient: redisClient,
		authClient:  authClient,
		chatClient:  chatClient,
	}
}
