package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	_ "github.com/joho/godotenv/autoload"
)

var (
	token    = os.Getenv("ff_token")
	guildID  = snowflake.GetEnv("ff_guild")
	logLevel = os.Getenv("ff_log")

	commands = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "echo",
			Description: "/ commands testing",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "message",
					Description: "What to say",
					Required:    true,
				},
				discord.ApplicationCommandOptionBool{
					Name:        "ephemeral",
					Description: "If the response should only be visible to you",
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:        "verify-hash",
			Description: "Verify hash from a string",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "string",
					Description: "The hashed string goes here",
					Required:    true,
				},
			},
		},
	}
)

func init() {
	switch logLevel {
	case "0":
		log.SetLevel(log.LevelError)
	case "1":
		log.SetLevel(log.LevelWarn)
	case "2":
		log.SetLevel(log.LevelInfo)
	case "3":
		log.SetLevel(log.LevelDebug)
	}

	//Check if package is installed
	out, err := exec.Command("npm", "list", "-g", "hash-detector-cli", "--depth", "1", "--parseable").Output()
	if err != nil {
		log.Fatal("Failed to run the command. Check if NPM is installed.\nDetails: ", err)
	}

	if !strings.Contains(string(out), "hash-detector-cli") {
		log.Fatal("hash-detector-cli is not installed. First, install that package with npm")
	}
}

func main() {
	log.Info("Starting FastForward Bot...")
	log.Info("Library (disgo) version: ", disgo.Version)

	client, err := disgo.New(token,
		bot.WithDefaultGateway(),
		bot.WithEventListenerFunc(commandListener),
		bot.WithGatewayConfigOpts(
			gateway.WithAutoReconnect(true),
			gateway.WithDevice("FastForward"),
			gateway.WithOS("Android"),
			gateway.WithPresence(gateway.NewPresence(discord.ActivityTypeGame, "with link shortners!", "https://fastforward.team", discord.OnlineStatusOnline, false)),
		),
	)

	if err != nil {
		log.Fatal("Error while building disgo instance!\nDetails: ", err)
		return
	}

	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGuildCommands(client.ApplicationID(), guildID, commands); err != nil {
		log.Fatal("Error while registering commands!\nDetails: ", err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		log.Fatal("Error while connecting to gateway!\nDetails: ", err)
	}

	//print the bot's user name
	log.Info("Bot ID: ", client.ApplicationID())
	log.Infof("Fast Forward Bot is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()
	switch data.CommandName() {
	case "echo":
		err := event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(data.String("message")).
			SetEphemeral(data.Bool("ephemeral")).
			Build(),
		)
		if err != nil {
			event.Client().Logger().Error("error on sending response: ", err)
		}
	case "verify-hash":
		cmd := exec.Command("hash-detect", data.String("string"), "-p")
		var result bytes.Buffer
		cmd.Stdout = &result
		err := cmd.Run()
		if err != nil {
			event.CreateMessage(discord.NewMessageCreateBuilder().
				SetContent("Failed to verify the hash!").
				SetEphemeral(data.Bool("ephemeral")).
				Build())
			event.Client().Logger().Error("Error on verifying hash: ", err)
		}
		log.Info(result.String())

		err = event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(string("Input: `" + data.String("string") + "`\nPossible hashes: " + result.String())).
			SetEphemeral(true).
			Build(),
		)
		if err != nil {
			event.Client().Logger().Error("error on sending response: ", err)
		}
	}
}
