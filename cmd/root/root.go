/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package root

import (
//	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/Oleg-Pro/chat-cli/internal/app"	
	"github.com/Oleg-Pro/chat-cli/internal/model"		
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-cli",
	Short: "Client application for chat",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

/*		model = &model.AuthInfo{
			Username: login,
			Password: password,
		}		*/

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
		usernames, err := cmd.Flags().GetString("user_names")
		if err != nil {
			log.Fatalf("failed to get login: %s\n", err.Error())
		}

		log.Printf("User login=%s \n", usernames)
	},
}

/*var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Создает нового пользователя",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		log.Printf("user %s created\n", usernamesStr)
	},
}*/

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chat-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createChatCmd)

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

}
