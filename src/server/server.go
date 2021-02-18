package server

import (
	"encoding/json"
	"io/ioutil"
	"mystery.tech/m/v2/src/handler"
	"mystery.tech/m/v2/src/structs"
	"mystery.tech/m/v2/src/structs/telegram/api/updates"
	"mystery.tech/m/v2/src/utils"
	"net/http"
	"strings"
)

func StartServer() {
	http.HandleFunc("/telegram", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "Cannot GET /telegram", http.StatusMethodNotAllowed)
			return
		}

		byteValue, _ := ioutil.ReadFile("config.json")
		var config structs.Config
		_ = json.Unmarshal(byteValue, &config)

		var update updates.Update
		body, _ := ioutil.ReadAll(request.Body)
		_ = json.Unmarshal(body, &update)

		if update.InlineQuery != nil {
			_, _ = writer.Write([]byte(http.StatusText(200)))
			handler.HandleResult(update.InlineQuery.Id, update.InlineQuery.Query)
			return
		}

		if update.CallbackQuery != nil {
			_, _ = writer.Write([]byte(http.StatusText(200)))
			if strings.Contains(update.CallbackQuery.Data, "INF") {
				trackId := strings.Split(update.CallbackQuery.Data, "INF")[0]
				trackId = strings.TrimSpace(trackId)
				handler.RetrieveInfo(trackId, update.CallbackQuery.InlineMessageId)
				return
			}
			if strings.Contains(update.CallbackQuery.Data, "STR") {
				trackId := strings.Split(update.CallbackQuery.Data, "STR")[0]
				trackId = strings.TrimSpace(trackId)
				handler.RetrievePlatforms(trackId, update.CallbackQuery.InlineMessageId)
			}
		}
	})

	utils.Info("Starting webserver at port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
