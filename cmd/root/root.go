/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package root

import (
//	"context"
	"log"
	"os"
	"strings"
    "strconv"
	"github.com/spf13/cobra"
	"github.com/Oleg-Pro/chat-cli/internal/app"	
	"github.com/Oleg-Pro/chat-cli/internal/model"		
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-cli",
	Short: "Client application for chat",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to chat.",
	Run: func(cmd *cobra.Command, args []string) {
		login, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatalf("failed to get login: %s\n", err.Error())
		}

		log.Printf("User login=%s \n", login)

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}
		
		log.Printf("User password=%s \n", password)

		ctx := cmd.Context()		
		serviceProvider := app.NewServiceProvider()	
		handlerService := serviceProvider.GetHandlerService(ctx)

		err = handlerService.Login(ctx, &model.AuthInfo{
			Username: login,
			Password: password,
		})
		if err != nil {
			log.Fatalf("failed to login: %s\n", err.Error())
		}

		log.Println("login success")			
	},
}

var createChatCmd = &cobra.Command{
	Use:   "create_chat",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("usernames")
		if err != nil {
			log.Fatalf("failed to get login: %s\n", err.Error())
		}

		usernames := strings.Split(usernamesStr, ",")
		if len(usernames) == 0 {
			log.Fatalf("usernames must be not empty")
		}


		ctx := cmd.Context()
		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		chatID, err := handlerService.CreateChat(ctx, usernames)
		if err != nil {
			log.Fatalf("failed to create chat: %s\n", err.Error())
		}

		log.Printf("chat created with id: %d\n", chatID)		
	},
}


var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Подключается к чат-серверу",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		chatIDstr, err := cmd.Flags().GetString("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat id: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)


		chatID, err := strconv.ParseInt(chatIDstr, 10 , 64)
		if err != nil {
			log.Fatalf("failed to convert chatId to int: %s\n", err.Error())			
		}


		err = handlerService.ConnectChat(ctx, chatID)
		if err != nil {
			log.Fatalf("failed to connect: %s\n", err.Error())
		}

		log.Println("chat finished")
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createChatCmd)
	rootCmd.AddCommand(connectCmd)

	loginCmd.Flags().StringP("login", "l", "", "Имя пользователя (адрес электронной почты)")
	err := loginCmd.MarkFlagRequired("login")
	if err != nil {
		log.Fatalf("failed to mark login flag as required: %s\n", err.Error())
	}
	loginCmd.Flags().StringP("password", "p", "p", "Пароль")
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}

	createChatCmd.Flags().StringP("usernames", "u", "", "Список пользователе (строка, разделенная запятыми)")
	err = createChatCmd.MarkFlagRequired("usernames")
	if err != nil {
		log.Fatalf("failed to mark usernames flag as required: %s\n", err.Error())
	}

	connectCmd.Flags().StringP("chat-id", "c", "", "Chat id")
	err = connectCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}
}
