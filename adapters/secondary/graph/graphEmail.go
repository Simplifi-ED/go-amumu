package graph

type IGraphEmail interface {
	SendMail(sender string, receiver string, subject string, body string) error
}

type GraphEmail struct {
	Graph *GraphHelper
}

func (g *GraphEmail) SendMail(sender string, receiver string, subject string, body string) error {

	err := g.Graph.Send(&subject, &body, &sender, &receiver)
	if err != nil {
		return err
	}
	return nil

}
