package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Guild struct {
		OCM_GUILD_ID string `envconfig:"OCM_GUILD_ID" required:"true"`
	}
	Discord struct {
		OCM_URL_DISCORD_API_KEY string `envconfig:"OCM_URL_DISCORD_API_KEY" required:"true"`
	}
}

var guildId string

func main() {

	// Load config and env variables
	cfg := setupConfig()
	guildId = cfg.Guild.OCM_GUILD_ID

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + cfg.Discord.OCM_URL_DISCORD_API_KEY)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handleMessage)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}

func setupConfig() (cfg Config) {
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println("error decoding env variables", err)
	}
	return cfg
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	verifyIfContainsUrl(s, m, guildId, m.ChannelID, m.Content)
}

func verifyIfContainsUrl(s *discordgo.Session, m *discordgo.MessageCreate, gid, cid, message string) {
	s.ChannelMessageSend(cid, message)
}
