package chat_server

import (
	"context"

	"github.com/Oleg-Pro/chat-cli/internal/model"
	chatV1 "github.com/Oleg-Pro/chat-server/pkg/chat_v1"
	myGrps "github.com/Oleg-Pro/chat-cli/internal/client/grpc"		
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ myGrps.ChatServerClient = (*client)(nil)

type client struct {
	client chatV1.ChatV1Client
}

func NewClient(cc *grpc.ClientConn) *client {
	return &client{
		client: chatV1.NewChatV1Client(cc),
	}
}

func (c *client) CreateChat(ctx context.Context, usernames []string) (int64, error) {
	res, err := c.client.Create(ctx, &chatV1.CreateRequest{
		UserNames: usernames,
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (c *client) ConnectChat(ctx context.Context, chatID int64, username string) (chatV1.ChatV1_ConnectClient, error) {
	return c.client.Connect(ctx, &chatV1.ConnectRequest{
		ChatId:   chatID,
		Username: username,
	})
}

func (c *client) SendMessage(ctx context.Context, chatID int64, message *model.Message) error {
	_, err := c.client.SendMessage(ctx, &chatV1.SendMessageRequest{
		ChatId: chatID,
		Message: &chatV1.Message{
			From:      message.From,
			Text:      message.Text,
			CreatedAt: timestamppb.New(message.CreatedAt),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
