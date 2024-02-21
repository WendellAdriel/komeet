package cmd

import (
	"github.com/spf13/cobra"
	. "komeet/core"
	"komeet/models/auth"
	"time"
)

var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user for the application",
	Run:   createUser,
}

var userName string

var userEmail string

var userPassword string

func init() {
	createUserCmd.Flags().StringVarP(&userName, "name", "n", "", "The user name")
	createUserCmd.Flags().StringVarP(&userEmail, "email", "e", "", "The user email")
	createUserCmd.Flags().StringVarP(&userPassword, "password", "p", "", "The user password")

	createUserCmd.MarkFlagRequired("name")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(createUserCmd)
}

func createUser(cmd *cobra.Command, args []string) {
	logger := App.Logger()
	var user auth.User

	App.DB.Where("email = ?", userEmail).First(&user)
	if user.ID > 0 {
		logger.Fatal().Msgf("User with email %s already exists", userEmail)
	}

	user = auth.NewUser(userName, userEmail, userPassword)
	now := time.Now()
	user.EmailVerifiedAt = &now
	user.Active = true

	result := App.DB.Create(&user)
	if result.Error != nil {
		logger.Fatal().Err(result.Error).Msgf("Failed creating user %s (%s)", userName, userEmail)
	}

	logger.Info().Msgf("User %s (%s) created", userName, userEmail)
}
