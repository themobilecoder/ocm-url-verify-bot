package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"mvdan.cc/xurls/v2"
)

type Config struct {
	Discord struct {
		OCM_URL_DISCORD_API_KEY string `envconfig:"OCM_URL_DISCORD_API_KEY" required:"true"`
	}
}

var verifiedDomains map[string]bool

func main() {

	// Load config and env variables
	cfg := setupConfig()

	domains, err := ReadLines("verified-domains.txt")
	if err != nil {
		fmt.Println("error reading domains file,", err)
		return
	}
	verifiedDomains = domains

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

	rxStrict, err := xurls.StrictMatchingScheme("https")
	if err != nil {
		return
	}
	murls := rxStrict.FindAllString(m.Content, -1)
	if len(murls) != 0 {
		//Just verify the first valid url posted
		u, err := url.Parse(murls[0])
		if err != nil {
			return
		}
		hostName := u.Hostname()
		if verifiedDomains[hostName] {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "⬆️ URL verified ✅",
				Description: u.Hostname(),
			})
		}
	}
}
