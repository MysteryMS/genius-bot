package main

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"mystery.tech/m/v2/src/server"
	"mystery.tech/m/v2/src/structs"
	webhook2 "mystery.tech/m/v2/src/structs/telegram/api/webhook"
	"mystery.tech/m/v2/src/utils"
	"net/http"
)

func main() {
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	gwReq, gwErr := http.Get("https://api.telegram.org/bot" + config.Telegram + "/getWebhookInfo")
	if gwErr != nil {
		panic(gwReq)
	}

	defer gwReq.Body.Close()

	var webhook webhook2.Result
	webhookInfo, _ := ioutil.ReadAll(gwReq.Body)
	_ = json.Unmarshal(webhookInfo, &webhook)

	if config.Webhook != webhook.Result.Url {
		status := updateWebhook(config.Webhook, config.Telegram)
		if !status {
			utils.Warn("Setup failed")
			return
		}
		gwReq, gwErr = http.Get("https://api.telegram.org/bot" + config.Telegram + "/getWebhookInfo")
		webhookInfo, _ = ioutil.ReadAll(gwReq.Body)
		_ = json.Unmarshal(webhookInfo, &webhook)
	}

	if config.Debug {
		color.White("[INFO] Webhook status:\nCurrent URL: %s\nPending updates: %d\nLast error: %s\n",
			webhook.Result.Url,
			webhook.Result.Pending,
			webhook.Result.Error)
	}
	utils.Info("Setup complete")
	server.StartServer()
}

func updateWebhook(url string, token string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"url": url,
	})

	print("[INFO] New webhook URL detected, setting up...\n")
	swReq, err := http.Post("https://api.telegram.org/bot"+token+"/setWebhook", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		panic(err)
	}

	if swReq.StatusCode != 200 {
		print("[WARN] Request failed\n")
		return false
	}

	defer swReq.Body.Close()
	print("[INFO] Webhook POST status: " + swReq.Status + "\n")
	return true
}
