package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mystery.tech/m/v2/src/structs"
	"mystery.tech/m/v2/src/structs/genius/search"
	"mystery.tech/m/v2/src/structs/genius/track"
	"net/http"
)

type Genius interface {
	ResolveTrack() track.Body
	ResolveSearch() search.Body
}

type Query string

func (query Query) ResolveTrack() track.Body {
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/songs/%s?text_format=html", query), nil)

	req.Header.Add("Authorization", "Bearer "+config.Genius)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		Fatal(err.Error())
	}

	bArray, readError := ioutil.ReadAll(resp.Body)
	Debug("ResolveTrack: " + resp.Status)

	if readError != nil {
		Fatal(readError.Error())
	}

	var song track.Body
	_ = json.Unmarshal(bArray, &song)

	return song
}

func (query Query) ResolveSearch() search.Body {
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%s", query), nil)

	req.Header.Add("Authorization", "Bearer "+config.Genius)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		Fatal(err.Error())
	}

	bArray, readError := ioutil.ReadAll(resp.Body)
	Debug("ResolveSearch: " + resp.Status)
	if resp.StatusCode != 200 {
		Warn(string(bArray))
	}

	if readError != nil {
		Fatal(readError.Error())
	}

	var song search.Body
	_ = json.Unmarshal(bArray, &song)

	return song
}
