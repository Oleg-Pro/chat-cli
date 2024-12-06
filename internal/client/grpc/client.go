package grpc

import (
	"context"
	"github.com/Oleg-Pro/chat-cli/internal/model"	
	chatV1 "github.com/Oleg-Pro/chat-server/pkg/chat_v1"	
)

type AuthClient interface {
	GetRefreshToken(ctx context.Context, info *model.AuthInfo) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}


type ChatServerClient interface {
	CreateChat(ctx context.Context, usernames []string) (int64, error)
	ConnectChat(ctx context.Context, chatID int64, username string) (chatV1.ChatV1_ConnectClient, error)
	SendMessage(ctx context.Context, chatID int64, message *model.Message) error
}
