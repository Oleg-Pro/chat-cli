package auth

import (
	"context"

	authV1 "github.com/Oleg-Pro/auth/pkg/auth_v1"
	myGrps "github.com/Oleg-Pro/chat-cli/internal/client/grpc"			
	"github.com/Oleg-Pro/chat-cli/internal/model"
	"google.golang.org/grpc"
)

var _ myGrps.AuthClient = (*client)(nil)

type client struct {
	client authV1.AuthV1Client
}

func NewClient(cc *grpc.ClientConn) *client {
	return &client{
		client: authV1.NewAuthV1Client(cc),
	}
}

func (c *client) GetRefreshToken(ctx context.Context, info *model.AuthInfo) (string, error) {
	resp, err := c.client.Login(ctx, &authV1.LoginRequest{
		
		Username: info.Username,
		Password: info.Password,
	})
	if err != nil {
		return "", err
	}

	return resp.GetRefreshToken(), nil
}

func (c *client) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	resp, err := c.client.GetAccessToken(ctx, &authV1.GetAccessTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return "", err
	}

	return resp.GetAccessToken(), nil
}
