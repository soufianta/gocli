/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func generatePassword(length int) string {
	// Set the seed for random number generation
	rand.Seed(time.Now().UnixNano())
	// Define the characters that can be included in the random string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Generate the random string
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = characters[rand.Intn(len(characters))]
	}
	return string(randomString)
}

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "generates a random password",
	Long:  `generates a random password with a specific length if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generatePassword(8))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
