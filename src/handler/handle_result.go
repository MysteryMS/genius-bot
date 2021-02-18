package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"mystery.tech/m/v2/src/structs"
	"mystery.tech/m/v2/src/structs/telegram/api"
	"mystery.tech/m/v2/src/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func HandleResult(inlineQueryId string, query string) {
	defer utils.TimeTrack(time.Now(), "Show inline")
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	utils.StartTrack("Get tracks")
	t := utils.Genius.ResolveSearch(utils.Query(url.QueryEscape(query)))
	utils.StopTrack("Get tracks")
	inlineAnswer := api.AnswerInline{
		InlineQueryId: inlineQueryId,
	}

	utils.StartTrack("Build results")
	for _, element := range *t.Response.Hits {
		article := api.InlineQueryResultArticle{
			Type:     "article",
			Id:       xid.New().String(),
			Title:    element.Result.Title,
			ThumbUrl: element.Result.Cover,
			InputMessageContent: api.InputTextMessageContent{
				MessageText: fmt.Sprintf("%s (%d views)\n[​](%s)", strings.ReplaceAll(element.Result.Title, "*", "\\*"), element.Result.Stats.Pageviews, element.Result.Cover),
				ParseMode:   "Markdown",
			},
			ReplyMarkup: api.InlineKeyboardMarkup{
				InlineKeyboard: [][]api.InlineKeyboardButton{{{
					Text:         "➕ More information",
					CallbackData: strconv.Itoa(element.Result.Id) + " INF",
				}, {
					Text: "↗ Open on Genius",
					Url:  element.Result.Url,
				}}},
			},
		}

		inlineAnswer.Results = append(inlineAnswer.Results, article)
	}

	body, _ := json.Marshal(inlineAnswer)
	utils.StopTrack("Build results")

	req, _ := http.NewRequest("POST", "https://api.telegram.org/bot"+config.Telegram+"/answerInlineQuery", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	utils.StartTrack("POST results")
	resp, e := client.Do(req)
	utils.StopTrack("POST results")

	if e != nil {
		utils.Fatal(e.Error())
		return
	}

	utils.Debug("AnswerInline: " + resp.Status)

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		utils.Debug(string(body))
	}

	resp.Body.Close()
}
