package app

import (
	"context"
	"log"
	"time"

	"github.com/Oleg-Pro/chat-cli/internal/client/grpc/auth"
	chatServer "github.com/Oleg-Pro/chat-cli/internal/client/grpc/chat_server"
	"github.com/Oleg-Pro/chat-cli/internal/client/redis"
	"github.com/Oleg-Pro/chat-cli/internal/closer"
	"github.com/Oleg-Pro/chat-cli/internal/handler"
	"github.com/Oleg-Pro/chat-cli/internal/interceptor"
	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceProvider struct {
	authClient       auth.Client
	chatServerClient chatServer.Client
	redisClient      redis.Client

	handlerService *handler.Handler
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) GetAuthClient(ctx context.Context) auth.Client {
	if s.authClient == nil {
		con, err := grpc.DialContext(
			ctx,
			"localhost:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to dial auth service: %s", err.Error())
		}
		closer.Add(con.Close)

		s.authClient = auth.NewClient(con)
	}

	return s.authClient
}

func (s *ServiceProvider) GetChatClient(ctx context.Context) chatServer.Client {
	if s.chatServerClient == nil {
		authInterceptor := interceptor.NewAuthInterceptor(s.GetAuthClient(ctx), s.GetRedisClient())
		authInterceptor.Run(60*time.Minute, 1*time.Minute)

		conn, err := grpc.DialContext(
			ctx,
			"localhost:50052",
			grpc.WithUnaryInterceptor(authInterceptor.Unary),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to dial chat service: %s", err.Error())
		}
		closer.Add(func() error {
			if conn != nil {
				return conn.Close()
			}
			return nil
		})

		s.chatServerClient = chatServer.NewClient(conn)
	}

	return s.chatServerClient
}

func (s *ServiceProvider) GetRedisClient() redis.Client {
	if s.redisClient == nil {
		client := redis.NewClient("localhost:6378")
		closer.Add(func() error {
			if client != nil {
				return client.Close()
			}
			return nil
		})

		err := client.Ping()
		if err != nil {
			log.Fatalf("failed to ping redis: %s", err.Error())
		}

		s.redisClient = client
	}

	return s.redisClient
}

func (s *ServiceProvider) GetHandlerService(ctx context.Context) *handler.Handler {
	if s.handlerService == nil {
		s.handlerService = handler.NewHandler(
			s.GetRedisClient(),
			s.GetAuthClient(ctx),
			s.GetChatClient(ctx),
		)
	}

	return s.handlerService
}
