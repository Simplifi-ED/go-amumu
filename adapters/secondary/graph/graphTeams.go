package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-send/domain/entities"
	"net/http"
)

type GraphTeams struct{}

func (g *GraphTeams) SendAlert(data entities.WebhookData) error {
	// setup webhook url
	organization := "simplifiedfr"
	gid := "31d891a8-180c-40d2-ba10-002ba8856053"
	tenantId := "cd8312b6-130d-4078-964d-2faddd3a0aca"
	altId := "cb84ff9b05404a22a22f8adb36900ecb"
	gOwner := "7cb5c351-1a58-47ee-b6d2-cb8b652b55c5"
	webhookUrl := fmt.Sprintf("https://%s.webhook.office.com/webhookb2/%s@%s/IncomingWebhook/%s/%s", organization, gid, tenantId, altId, gOwner)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
