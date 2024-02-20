package entities

type Message struct {
	To      string `yaml:to`
	From    string `yaml:from`
	Subject string `yaml:suject`
	Body    string `yaml:body`
	Channel bool   `yaml:channel`
}
