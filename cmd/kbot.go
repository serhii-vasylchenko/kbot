package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "kbot is a telegram bot",
	Long:  "kbot is a telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started \n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			var (
				err error
			)

			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
				case "hello":
					err = m.Send(fmt.Sprintf("Hello, I'm Kbot %s!", appVersion))
				
			}

			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}