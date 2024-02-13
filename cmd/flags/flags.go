package flags

import (
	"flag"
	"fmt"
	"go-send/adapters/primary/cli"
	"go-send/adapters/primary/server"
	"go-send/adapters/secondary/graph"
	"go-send/infrastructure/notification"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type AmumuCmd struct {
	subject *notification.Subject
	graph   *graph.GraphHelper
}

func NewAmumuCmd() *AmumuCmd {
	log.Info("Initializing App...")
	loadEnv()

	graphHelper := graph.NewGraphHelper()
	graph.InitializeGraph(graphHelper)

	subjectmq := &notification.Subject{}
	graphEmail := &graph.GraphEmail{
		Graph: graphHelper,
	}
	subjectmq.SetGraphEmail(graphEmail)
	watcher := &notification.Watcher{}
	subjectmq.Register(watcher)

	return &AmumuCmd{
		subject: subjectmq,
		graph:   graphHelper,
	}
}

func (a *AmumuCmd) Run() {
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		server.RunServer()
	case "client":
		cli.RunClient(a.graph)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func loadEnv() {
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func SetCmdExec() {
	flag.Usage = func() {
		fmt.Println("\nUsage:")
		fmt.Printf(" %s [subcommand] [options]\n", os.Args[0])
		fmt.Println("server options:")
		fmt.Printf(" -port string\n")
		fmt.Println("client options:")
		fmt.Printf(" -to string - The email address of the recipient\n")
		fmt.Printf(" -from string - The email address of the sender\n")
		fmt.Printf(" -subject string - The subject of the email\n")
		fmt.Printf(" -message string - The message body of the email\n")
		fmt.Printf(" -channel boolean - Send to MS Teams channel (default=false)\n")
	}
}
