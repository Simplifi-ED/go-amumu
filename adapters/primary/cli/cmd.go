package cli

import (
	"flag"
	"go-send/adapters/primary"
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
	"go-send/domain/usecase"
	"os"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

func RunClient(u *usecase.UserCase) {
	ClientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	configFile := ClientCmd.String("config", "", "Path to the configuration file")
	to := ClientCmd.String("to", "", "The email address of the recipient")
	from := ClientCmd.String("from", "", "The email address of the sender")
	subject := ClientCmd.String("subject", "", "The subject of the email")
	message := ClientCmd.String("message", "", "The message body of the email")
	channel := ClientCmd.Bool("channel", false, "Send to MS Teams channel")
	ClientCmd.Parse(os.Args[2:])

	if *configFile != "" {
		log.Info("Using config file:", *configFile)
		primary.ValidateConfigPath(*configFile)
		clientConfig, err := NewClientConfig(*configFile)
		if err != nil {
			log.Fatal("Error loading config", "Error", err)
		}
		u.SendToGraph(clientConfig)
	} else if *to != "" && *from != "" && *subject != "" && *message != "" {
		m := &entities.Message{
			To:      *to,
			From:    *from,
			Subject: *subject,
			Body:    *message,
		}
		u.SendToGraph(m)
		if *channel {
			graphTeams := &graph.GraphTeams{}
			m.Channel.Title = *subject
			m.Channel.Text = *message
			graphTeams.SendAlert(m.Channel)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

}

func NewClientConfig(configPath string) (*entities.Message, error) {
	msg := entities.Message{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
