package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"mystery.tech/m/v2/src/structs"
	"mystery.tech/m/v2/src/structs/song_link"
	"mystery.tech/m/v2/src/structs/telegram/api"
	"mystery.tech/m/v2/src/structs/telegram/api/updates"
	"mystery.tech/m/v2/src/utils"
	"net/http"
	"strings"
	"time"
)

func RetrieveInfo(trackId string, queryId string) {
	defer utils.TimeTrack(time.Now(), "Retrieve info")
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	link := md.Rule{
		Filter: []string{"a"},
		Replacement: func(content string, selec *goquery.Selection, options *md.Options) *string {
			link := strings.ReplaceAll(selec.Nodes[0].Attr[0].Val, "_", "\\_")
			return md.String("[" + content + "]" + "(" + link + ")")
		},
	}

	converter := md.NewConverter("", true, nil)
	converter.AddRules(link)

	utils.StartTrack("Fetch track")
	t := utils.Genius.ResolveTrack(utils.Query(trackId)).Response.Song
	utils.StopTrack("Fetch track")
	utils.StartTrack("Conversions")
	var album string
	convertString, err := converter.ConvertString(strings.ReplaceAll(t.Description.Html, "<hr>", ""))

	if err != nil {
		utils.Fatal(err.Error())
		return
	}

	if t.Album != nil {
		album = fmt.Sprintf("💿 From the album _%s_", t.Album.Name)
	} else if strings.Contains(t.Album.Name, "*") {
		album = "💿 Unreleased Album"
	} else {
		album = "💿 This track is not present in any album"
	}

	metaData := fmt.Sprintf(
		"\n\n🎵 %s\n📆 Released in %s\n%s\n👩‍🎤 By *%s*",
		t.Title,
		t.ReleaseDate,
		strings.ReplaceAll(album, "*", "\\*"),
		t.PrimaryArtist.Name,
	)

	/*	if len(strings.Split(bio, "\n")) > 2 {
			bio += metaData
		} else {
			bio = strings.Split(bio, "\n")[0] + metaData
		}*/

	replacer := strings.NewReplacer("-", "\\-", ".", "\\.", ">", "\\>", "!", "\\!", "+", "\\+", "=", "\\=", "(", "\\(", ")", "\\)")
	metaReplacer := strings.NewReplacer("(", "\\(", ")", "\\)", ".", "\\", "+", "\\+", "!", "\\!", "-", "\\-")
	bio := replacer.Replace(convertString)
	if bio == "?" {
		bio = "No bio available."
	}
	meta := metaReplacer.Replace(metaData)
	utils.StopTrack("Conversions")
	message := updates.Message{
		InlineMessageId: queryId,
		Text:            bio + meta,
		ParseMode:       "MarkdownV2",
		ReplyMarkup: api.InlineKeyboardMarkup{InlineKeyboard: [][]api.InlineKeyboardButton{{{
			Text: "↗ Open on Genius",
			Url:  t.Url,
		}}, {{
			Text:         "🎧 Streaming Platforms",
			CallbackData: t.AppleMusicId + " STR",
		}}}},
		DisableWebpagePreview: true,
	}

	bArray, _ := json.Marshal(message)

	utils.StartTrack("POST Telegram")
	r, e := http.Post("https://api.telegram.org/bot"+config.Telegram+"/editMessageText", "application/json", bytes.NewBuffer(bArray))
	utils.StopTrack("POST Telegram")

	if e != nil {
		utils.Fatal(e.Error())
		return
	}
	utils.Debug("Edit Text: " + r.Status)
	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(r.Body)
		utils.Debug(string(body))
	}

}

func RetrievePlatforms(trackId string, queryId string) {
	defer utils.TimeTrack(time.Now(), "Streaming Platforms")
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	utils.StartTrack("SongLink request")
	songLink, _ := http.Get(fmt.Sprintf("https://api.song.link/v1-alpha.1/links?platform=itunes&type=song&id=%s&key=%s", trackId, config.SongLink))
	utils.StopTrack("SongLink request")

	defer songLink.Body.Close()

	utils.Debug("SongLink: " + songLink.Status)
	if songLink.StatusCode != 200 {
		body, _ := ioutil.ReadAll(songLink.Body)
		utils.Debug(string(body))
	}

	bArray, _ := ioutil.ReadAll(songLink.Body)
	var song song_link.Body

	err := json.Unmarshal(bArray, &song)
	if err != nil {
		utils.Fatal(err.Error())
		return
	}

	message := updates.MarkupMessage{
		InlineMessageId: queryId,
		ReplyMarkup: api.InlineKeyboardMarkup{InlineKeyboard: [][]api.InlineKeyboardButton{{{
			Text: "Apple Music",
			Url:  song.LinksByPlatform.AppleMusic.Url,
		}, {
			Text: "Amazon Music",
			Url:  song.LinksByPlatform.AmazonMusic.Url,
		}, {
			Text: "Deezer",
			Url:  song.LinksByPlatform.Deezer.Url,
		}}, {{
			Text: "Spotify",
			Url:  song.LinksByPlatform.Spotify.Url,
		}, {
			Text: "YouTube",
			Url:  song.LinksByPlatform.Youtube.Url,
		}, {
			Text: "YouTube Music",
			Url:  song.LinksByPlatform.YoutubeMusic.Url,
		}}}},
	}

	marshaEoUrso, _ := json.Marshal(message)

	utils.StartTrack("Edit markup")
	r, e := http.Post("https://api.telegram.org/bot"+config.Telegram+"/editMessageReplyMarkup", "application/json", bytes.NewBuffer(marshaEoUrso))
	utils.StopTrack("Edit markup")
	defer r.Body.Close()

	if e != nil {
		utils.Fatal(e.Error())
		return
	}

	utils.Debug("Edit to platforms: " + r.Status)
	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(songLink.Body)
		utils.Debug(string(body))
	}

}
